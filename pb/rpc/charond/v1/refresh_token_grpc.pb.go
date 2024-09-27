// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: github.com/piotrkowalczuk/charon/pb/rpc/charond/v1/refresh_token.proto

package charond

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RefreshTokenManager_Create_FullMethodName = "/charon.rpc.charond.v1.RefreshTokenManager/Create"
	RefreshTokenManager_Revoke_FullMethodName = "/charon.rpc.charond.v1.RefreshTokenManager/Revoke"
	RefreshTokenManager_List_FullMethodName   = "/charon.rpc.charond.v1.RefreshTokenManager/List"
)

// RefreshTokenManagerClient is the client API for RefreshTokenManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RefreshTokenManagerClient interface {
	Create(ctx context.Context, in *CreateRefreshTokenRequest, opts ...grpc.CallOption) (*CreateRefreshTokenResponse, error)
	Revoke(ctx context.Context, in *RevokeRefreshTokenRequest, opts ...grpc.CallOption) (*RevokeRefreshTokenResponse, error)
	List(ctx context.Context, in *ListRefreshTokensRequest, opts ...grpc.CallOption) (*ListRefreshTokensResponse, error)
}

type refreshTokenManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewRefreshTokenManagerClient(cc grpc.ClientConnInterface) RefreshTokenManagerClient {
	return &refreshTokenManagerClient{cc}
}

func (c *refreshTokenManagerClient) Create(ctx context.Context, in *CreateRefreshTokenRequest, opts ...grpc.CallOption) (*CreateRefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRefreshTokenResponse)
	err := c.cc.Invoke(ctx, RefreshTokenManager_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refreshTokenManagerClient) Revoke(ctx context.Context, in *RevokeRefreshTokenRequest, opts ...grpc.CallOption) (*RevokeRefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeRefreshTokenResponse)
	err := c.cc.Invoke(ctx, RefreshTokenManager_Revoke_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refreshTokenManagerClient) List(ctx context.Context, in *ListRefreshTokensRequest, opts ...grpc.CallOption) (*ListRefreshTokensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRefreshTokensResponse)
	err := c.cc.Invoke(ctx, RefreshTokenManager_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RefreshTokenManagerServer is the server API for RefreshTokenManager service.
// All implementations must embed UnimplementedRefreshTokenManagerServer
// for forward compatibility.
type RefreshTokenManagerServer interface {
	Create(context.Context, *CreateRefreshTokenRequest) (*CreateRefreshTokenResponse, error)
	Revoke(context.Context, *RevokeRefreshTokenRequest) (*RevokeRefreshTokenResponse, error)
	List(context.Context, *ListRefreshTokensRequest) (*ListRefreshTokensResponse, error)
	mustEmbedUnimplementedRefreshTokenManagerServer()
}

// UnimplementedRefreshTokenManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRefreshTokenManagerServer struct{}

func (UnimplementedRefreshTokenManagerServer) Create(context.Context, *CreateRefreshTokenRequest) (*CreateRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRefreshTokenManagerServer) Revoke(context.Context, *RevokeRefreshTokenRequest) (*RevokeRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}
func (UnimplementedRefreshTokenManagerServer) List(context.Context, *ListRefreshTokensRequest) (*ListRefreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRefreshTokenManagerServer) mustEmbedUnimplementedRefreshTokenManagerServer() {}
func (UnimplementedRefreshTokenManagerServer) testEmbeddedByValue()                             {}

// UnsafeRefreshTokenManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RefreshTokenManagerServer will
// result in compilation errors.
type UnsafeRefreshTokenManagerServer interface {
	mustEmbedUnimplementedRefreshTokenManagerServer()
}

func RegisterRefreshTokenManagerServer(s grpc.ServiceRegistrar, srv RefreshTokenManagerServer) {
	// If the following call pancis, it indicates UnimplementedRefreshTokenManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RefreshTokenManager_ServiceDesc, srv)
}

func _RefreshTokenManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefreshTokenManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RefreshTokenManager_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefreshTokenManagerServer).Create(ctx, req.(*CreateRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RefreshTokenManager_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefreshTokenManagerServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RefreshTokenManager_Revoke_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefreshTokenManagerServer).Revoke(ctx, req.(*RevokeRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RefreshTokenManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRefreshTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefreshTokenManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RefreshTokenManager_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefreshTokenManagerServer).List(ctx, req.(*ListRefreshTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RefreshTokenManager_ServiceDesc is the grpc.ServiceDesc for RefreshTokenManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RefreshTokenManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "charon.rpc.charond.v1.RefreshTokenManager",
	HandlerType: (*RefreshTokenManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RefreshTokenManager_Create_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _RefreshTokenManager_Revoke_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RefreshTokenManager_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/piotrkowalczuk/charon/pb/rpc/charond/v1/refresh_token.proto",
}
