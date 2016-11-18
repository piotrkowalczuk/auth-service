package charond

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-ldap/ldap"
	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/charon/internal/model"
	"github.com/piotrkowalczuk/charon/internal/password"
	"github.com/piotrkowalczuk/mnemosyne/mnemosynerpc"
	"github.com/piotrkowalczuk/promgrpc"
	"github.com/piotrkowalczuk/sklog"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

func initPostgres(address string, test bool, logger log.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", address)
	if err != nil {
		return nil, fmt.Errorf("connection failure: %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		cancel := time.NewTimer(10 * time.Second)
		attempts := 1
	PingLoop:
		for {
			select {
			case <-time.After(1 * time.Second):
				if err := db.Ping(); err != nil {
					attempts++
					continue PingLoop
				}
				break PingLoop
			case <-cancel.C:
				return nil, fmt.Errorf("postgres connection failed after %d attempts", attempts)
			}
		}
	}

	if test {
		if err = teardownDatabase(db); err != nil {
			return nil, err
		}
		sklog.Info(logger, "database has been cleared upfront")
	}
	err = setupDatabase(db)
	if err != nil {
		return nil, err
	}

	sklog.Info(logger, "postgres connection has been established", "address", address)

	return db, nil
}

func initMnemosyne(address string, interceptor *promgrpc.Interceptor, logger log.Logger) (mnemosynerpc.SessionManagerClient, *grpc.ClientConn) {
	if address == "" {
		sklog.Fatal(logger, errors.New("missing mnemosyne address"))

	}
	conn, err := grpc.Dial(address,
		grpc.WithUserAgent("charond"),
		grpc.WithInsecure(),
		grpc.WithTimeout(5*time.Second),
		grpc.WithDialer(interceptor.Dialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("tcp", addr, timeout)
		})),
		grpc.WithUnaryInterceptor(interceptor.UnaryClient()),
		grpc.WithStreamInterceptor(interceptor.StreamClient()),
	)
	if err != nil {
		sklog.Fatal(logger, err, "address", address)
	}

	sklog.Info(logger, "rpc connection to mnemosyne has been established", "address", address)

	return mnemosynerpc.NewSessionManagerClient(conn), conn
}

func initHasher(cost int, logger log.Logger) password.Hasher {
	bh, err := password.NewBCryptHasher(cost)
	if err != nil {
		sklog.Fatal(logger, err)
	}

	return bh
}

func initPermissionRegistry(r model.PermissionProvider, permissions charon.Permissions, logger log.Logger) (pr model.PermissionRegistry) {
	pr = model.NewPermissionRegistry(r)
	created, untouched, removed, err := pr.Register(permissions)
	if err != nil {
		sklog.Fatal(logger, err)
	}

	sklog.Info(logger, "charon permissions has been registered", "created", created, "untouched", untouched, "removed", removed)

	return
}

func initPrometheus(namespace string, enabled bool, constLabels prometheus.Labels) *monitoring {
	rpcRequests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "rpc",
			Name:        "requests_total",
			Help:        "Total number of RPC requests made.",
			ConstLabels: constLabels,
		},
		monitoringRPCLabels,
	)
	rpcErrors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "rpc",
			Name:        "errors_total",
			Help:        "Total number of errors that happen during RPC calles.",
			ConstLabels: constLabels,
		},
		monitoringRPCLabels,
	)
	rpcDuration := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:   namespace,
			Subsystem:   "rpc",
			Name:        "request_duration_microseconds",
			Help:        "The RPC request latencies in microseconds.",
			ConstLabels: constLabels,
		},
		[]string{"handler", "code"},
	)

	postgresQueries := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "dal",
			Name:        "postgres_queries_total",
			Help:        "Total number of SQL queries made.",
			ConstLabels: constLabels,
		},
		monitoringPostgresLabels,
	)
	postgresErrors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "dal",
			Name:        "postgres_errors_total",
			Help:        "Total number of errors that happen during SQL queries.",
			ConstLabels: constLabels,
		},
		monitoringPostgresLabels,
	)

	if enabled {
		prometheus.MustRegisterOrGet(rpcRequests)
		prometheus.MustRegisterOrGet(rpcDuration)
		prometheus.MustRegisterOrGet(rpcErrors)
		prometheus.MustRegisterOrGet(postgresQueries)
		prometheus.MustRegisterOrGet(postgresErrors)
	}

	return &monitoring{
		enabled: true,
		rpc: monitoringRPC{
			enabled:  true,
			requests: rpcRequests,
			errors:   rpcErrors,
			duration: rpcDuration,
		},
		postgres: monitoringPostgres{
			enabled: true,
			queries: postgresQueries,
			errors:  postgresErrors,
		},
	}
}

func initLDAP(address, baseDN, password string, logger log.Logger) (*ldap.Conn, error) {
	init := func(addr string) (*ldap.Conn, error) {
		conn, err := ldap.Dial("tcp", addr)
		if err != nil {
			return nil, fmt.Errorf("ldap connection failure: %s", err.Error())
		}
		err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return nil, fmt.Errorf("ldap start tls reconnection failure: %s", err.Error())
		}
		if err = conn.Bind(baseDN, password); err != nil {
			return nil, fmt.Errorf("ldap bind failure: %s", err.Error())
		}

		sklog.Info(logger, "ldap connection has been established", "address", address, "dn", baseDN)
		return conn, nil
	}

	_, addresses, err := net.LookupSRV("ldap", "tcp", address)
	if err != nil || len(addresses) == 0 {
		return init(address)
	}

	for _, addr := range addresses {
		addressFound := fmt.Sprintf("%s:%d", addr.Target, addr.Port)
		c, err := init(addressFound)
		if err != nil {
			sklog.Error(logger, err, "target", addr.Target, "port", addr.Port, "ldap_base_dn", baseDN)
			continue
		}
		return c, nil
	}

	return nil, errors.New("ldap connection failed to a single server")
}
