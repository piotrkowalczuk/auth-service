// Code generated by mockery v1.0.0. DO NOT EDIT.

package charondmock

import (
	context "context"

	charond "github.com/piotrkowalczuk/charon/pb/rpc/charond/v1"

	empty "github.com/golang/protobuf/ptypes/empty"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// AuthClient is an autogenerated mock type for the AuthClient type
type AuthClient struct {
	mock.Mock
}

// Actor provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) Actor(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*charond.ActorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *charond.ActorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *wrappers.StringValue, ...grpc.CallOption) *charond.ActorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*charond.ActorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *wrappers.StringValue, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BelongsTo provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) BelongsTo(ctx context.Context, in *charond.BelongsToRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wrappers.BoolValue
	if rf, ok := ret.Get(0).(func(context.Context, *charond.BelongsToRequest, ...grpc.CallOption) *wrappers.BoolValue); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrappers.BoolValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.BelongsToRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthenticated provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) IsAuthenticated(ctx context.Context, in *charond.IsAuthenticatedRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wrappers.BoolValue
	if rf, ok := ret.Get(0).(func(context.Context, *charond.IsAuthenticatedRequest, ...grpc.CallOption) *wrappers.BoolValue); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrappers.BoolValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.IsAuthenticatedRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsGranted provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) IsGranted(ctx context.Context, in *charond.IsGrantedRequest, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wrappers.BoolValue
	if rf, ok := ret.Get(0).(func(context.Context, *charond.IsGrantedRequest, ...grpc.CallOption) *wrappers.BoolValue); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrappers.BoolValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.IsGrantedRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) Login(ctx context.Context, in *charond.LoginRequest, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wrappers.StringValue
	if rf, ok := ret.Get(0).(func(context.Context, *charond.LoginRequest, ...grpc.CallOption) *wrappers.StringValue); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wrappers.StringValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) Logout(ctx context.Context, in *charond.LogoutRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *charond.LogoutRequest, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *charond.LogoutRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
