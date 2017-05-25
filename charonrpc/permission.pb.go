// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission.proto

package charonrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import qtypes "github.com/piotrkowalczuk/qtypes"
import ntypes "github.com/piotrkowalczuk/ntypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegisterPermissionsRequest struct {
	Permissions []string `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *RegisterPermissionsRequest) Reset()                    { *m = RegisterPermissionsRequest{} }
func (m *RegisterPermissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterPermissionsRequest) ProtoMessage()               {}
func (*RegisterPermissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RegisterPermissionsRequest) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type RegisterPermissionsResponse struct {
	Created   int64 `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Removed   int64 `protobuf:"varint,2,opt,name=removed" json:"removed,omitempty"`
	Untouched int64 `protobuf:"varint,3,opt,name=untouched" json:"untouched,omitempty"`
}

func (m *RegisterPermissionsResponse) Reset()                    { *m = RegisterPermissionsResponse{} }
func (m *RegisterPermissionsResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterPermissionsResponse) ProtoMessage()               {}
func (*RegisterPermissionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *RegisterPermissionsResponse) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RegisterPermissionsResponse) GetRemoved() int64 {
	if m != nil {
		return m.Removed
	}
	return 0
}

func (m *RegisterPermissionsResponse) GetUntouched() int64 {
	if m != nil {
		return m.Untouched
	}
	return 0
}

type ListPermissionsRequest struct {
	Subsystem *qtypes.String    `protobuf:"bytes,1,opt,name=subsystem" json:"subsystem,omitempty"`
	Module    *qtypes.String    `protobuf:"bytes,2,opt,name=module" json:"module,omitempty"`
	Action    *qtypes.String    `protobuf:"bytes,3,opt,name=action" json:"action,omitempty"`
	CreatedAt *qtypes.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	CreatedBy *qtypes.Int64     `protobuf:"bytes,5,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	Offset    *ntypes.Int64     `protobuf:"bytes,100,opt,name=offset" json:"offset,omitempty"`
	Limit     *ntypes.Int64     `protobuf:"bytes,101,opt,name=limit" json:"limit,omitempty"`
	Sort      map[string]bool   `protobuf:"bytes,102,rep,name=sort" json:"sort,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ListPermissionsRequest) Reset()                    { *m = ListPermissionsRequest{} }
func (m *ListPermissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPermissionsRequest) ProtoMessage()               {}
func (*ListPermissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ListPermissionsRequest) GetSubsystem() *qtypes.String {
	if m != nil {
		return m.Subsystem
	}
	return nil
}

func (m *ListPermissionsRequest) GetModule() *qtypes.String {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *ListPermissionsRequest) GetAction() *qtypes.String {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ListPermissionsRequest) GetCreatedAt() *qtypes.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ListPermissionsRequest) GetCreatedBy() *qtypes.Int64 {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *ListPermissionsRequest) GetOffset() *ntypes.Int64 {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *ListPermissionsRequest) GetLimit() *ntypes.Int64 {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *ListPermissionsRequest) GetSort() map[string]bool {
	if m != nil {
		return m.Sort
	}
	return nil
}

type ListPermissionsResponse struct {
	Permissions []string `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *ListPermissionsResponse) Reset()                    { *m = ListPermissionsResponse{} }
func (m *ListPermissionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPermissionsResponse) ProtoMessage()               {}
func (*ListPermissionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ListPermissionsResponse) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type GetPermissionRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetPermissionRequest) Reset()                    { *m = GetPermissionRequest{} }
func (m *GetPermissionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPermissionRequest) ProtoMessage()               {}
func (*GetPermissionRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *GetPermissionRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPermissionResponse struct {
	Permission string `protobuf:"bytes,1,opt,name=permission" json:"permission,omitempty"`
}

func (m *GetPermissionResponse) Reset()                    { *m = GetPermissionResponse{} }
func (m *GetPermissionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPermissionResponse) ProtoMessage()               {}
func (*GetPermissionResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *GetPermissionResponse) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterPermissionsRequest)(nil), "charonrpc.RegisterPermissionsRequest")
	proto.RegisterType((*RegisterPermissionsResponse)(nil), "charonrpc.RegisterPermissionsResponse")
	proto.RegisterType((*ListPermissionsRequest)(nil), "charonrpc.ListPermissionsRequest")
	proto.RegisterType((*ListPermissionsResponse)(nil), "charonrpc.ListPermissionsResponse")
	proto.RegisterType((*GetPermissionRequest)(nil), "charonrpc.GetPermissionRequest")
	proto.RegisterType((*GetPermissionResponse)(nil), "charonrpc.GetPermissionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PermissionManager service

type PermissionManagerClient interface {
	Register(ctx context.Context, in *RegisterPermissionsRequest, opts ...grpc.CallOption) (*RegisterPermissionsResponse, error)
	List(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error)
	Get(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*GetPermissionResponse, error)
}

type permissionManagerClient struct {
	cc *grpc.ClientConn
}

func NewPermissionManagerClient(cc *grpc.ClientConn) PermissionManagerClient {
	return &permissionManagerClient{cc}
}

func (c *permissionManagerClient) Register(ctx context.Context, in *RegisterPermissionsRequest, opts ...grpc.CallOption) (*RegisterPermissionsResponse, error) {
	out := new(RegisterPermissionsResponse)
	err := grpc.Invoke(ctx, "/charonrpc.PermissionManager/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionManagerClient) List(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error) {
	out := new(ListPermissionsResponse)
	err := grpc.Invoke(ctx, "/charonrpc.PermissionManager/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionManagerClient) Get(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*GetPermissionResponse, error) {
	out := new(GetPermissionResponse)
	err := grpc.Invoke(ctx, "/charonrpc.PermissionManager/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PermissionManager service

type PermissionManagerServer interface {
	Register(context.Context, *RegisterPermissionsRequest) (*RegisterPermissionsResponse, error)
	List(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error)
	Get(context.Context, *GetPermissionRequest) (*GetPermissionResponse, error)
}

func RegisterPermissionManagerServer(s *grpc.Server, srv PermissionManagerServer) {
	s.RegisterService(&_PermissionManager_serviceDesc, srv)
}

func _PermissionManager_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionManagerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.PermissionManager/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionManagerServer).Register(ctx, req.(*RegisterPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.PermissionManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionManagerServer).List(ctx, req.(*ListPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charonrpc.PermissionManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionManagerServer).Get(ctx, req.(*GetPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PermissionManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "charonrpc.PermissionManager",
	HandlerType: (*PermissionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PermissionManager_Register_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PermissionManager_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PermissionManager_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission.proto",
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xed, 0x8a, 0xd3, 0x4c,
	0x14, 0xc7, 0x9f, 0x36, 0x6d, 0x9f, 0xe6, 0x14, 0x97, 0xee, 0xec, 0xaa, 0x43, 0x14, 0xad, 0x91,
	0x2d, 0x05, 0x25, 0x4a, 0x15, 0x57, 0x14, 0x14, 0x05, 0x59, 0x5c, 0x14, 0x25, 0xeb, 0x37, 0x3f,
	0x48, 0x9a, 0x9c, 0x76, 0x07, 0x9b, 0x99, 0xec, 0xcc, 0x64, 0x21, 0x17, 0xe1, 0x4d, 0x79, 0x65,
	0x92, 0x64, 0xf2, 0x82, 0x9b, 0xb5, 0x9f, 0xda, 0x39, 0xff, 0xdf, 0x79, 0xc9, 0x7f, 0x4e, 0x02,
	0xd3, 0x04, 0x65, 0xcc, 0x94, 0x62, 0x82, 0x7b, 0x89, 0x14, 0x5a, 0x10, 0x3b, 0x3c, 0x0f, 0xa4,
	0xe0, 0x32, 0x09, 0x9d, 0x83, 0x0b, 0x9d, 0x25, 0xa8, 0x9e, 0x94, 0x3f, 0xa5, 0xee, 0x1c, 0xf0,
	0x32, 0xc8, 0x5b, 0x41, 0xf7, 0x0d, 0x38, 0x3e, 0x6e, 0x98, 0xd2, 0x28, 0xbf, 0xd6, 0x05, 0x95,
	0x8f, 0x17, 0x29, 0x2a, 0x4d, 0x66, 0x30, 0x69, 0xda, 0x28, 0xda, 0x9b, 0x59, 0x0b, 0xdb, 0x6f,
	0x87, 0x5c, 0x01, 0x77, 0x3a, 0xf3, 0x55, 0x22, 0xb8, 0x42, 0x42, 0xe1, 0xff, 0x50, 0x62, 0xa0,
	0x31, 0xa2, 0xbd, 0x59, 0x6f, 0x61, 0xf9, 0xd5, 0x31, 0x57, 0x24, 0xc6, 0xe2, 0x12, 0x23, 0xda,
	0x2f, 0x15, 0x73, 0x24, 0x77, 0xc1, 0x4e, 0xb9, 0x16, 0x69, 0x78, 0x8e, 0x11, 0xb5, 0x0a, 0xad,
	0x09, 0xb8, 0xbf, 0x2d, 0xb8, 0xf5, 0x89, 0x29, 0xdd, 0x31, 0xed, 0x63, 0xb0, 0x55, 0xba, 0x52,
	0x99, 0xd2, 0x18, 0x17, 0xed, 0x26, 0xcb, 0x3d, 0xcf, 0x58, 0x70, 0xa6, 0x25, 0xe3, 0x1b, 0xbf,
	0x01, 0xc8, 0x1c, 0x46, 0xb1, 0x88, 0xd2, 0x2d, 0x16, 0xfd, 0xaf, 0xa2, 0x46, 0xcd, 0xb9, 0x20,
	0xd4, 0x4c, 0xf0, 0x62, 0x96, 0x0e, 0xae, 0x54, 0xc9, 0x53, 0x00, 0xf3, 0x6c, 0x3f, 0x02, 0x4d,
	0x07, 0x05, 0xbb, 0x5f, 0xb1, 0xdf, 0x58, 0x8c, 0x4a, 0x07, 0x71, 0xe2, 0xdb, 0x06, 0x7a, 0x97,
	0xcf, 0x5b, 0x67, 0xac, 0x32, 0x3a, 0x2c, 0x32, 0x6e, 0x54, 0x19, 0x1f, 0xb9, 0x7e, 0xf1, 0xbc,
	0xa6, 0xdf, 0x67, 0xe4, 0x08, 0x46, 0x62, 0xbd, 0x56, 0xa8, 0x69, 0x64, 0x48, 0xde, 0x26, 0x8d,
	0x48, 0x1e, 0xc2, 0x70, 0xcb, 0x62, 0xa6, 0x29, 0x76, 0x51, 0xa5, 0x46, 0xde, 0xc2, 0x40, 0x09,
	0xa9, 0xe9, 0x7a, 0x66, 0x2d, 0x26, 0xcb, 0x47, 0x5e, 0xbd, 0x39, 0x5e, 0xb7, 0xb5, 0xde, 0x99,
	0x90, 0xfa, 0x03, 0xd7, 0x32, 0xf3, 0x8b, 0x44, 0xe7, 0x18, 0xec, 0x3a, 0x44, 0xa6, 0x60, 0xfd,
	0xc4, 0xac, 0x70, 0xdc, 0xf6, 0xf3, 0xbf, 0xe4, 0x10, 0x86, 0x97, 0xc1, 0x36, 0x2d, 0xad, 0x1d,
	0xfb, 0xe5, 0xe1, 0x55, 0xff, 0x65, 0xef, 0x74, 0x30, 0x1e, 0x4d, 0x23, 0xf7, 0x35, 0xdc, 0xbe,
	0xd2, 0xc8, 0x6c, 0xcc, 0xee, 0x95, 0x9b, 0xc3, 0xe1, 0x09, 0xb6, 0x72, 0xab, 0xeb, 0xdf, 0x83,
	0x3e, 0xab, 0xd6, 0xac, 0xcf, 0x22, 0xf7, 0x18, 0x6e, 0xfe, 0xc5, 0x99, 0x16, 0xf7, 0x00, 0x9a,
	0x7a, 0x66, 0xec, 0x56, 0x64, 0xf9, 0xab, 0x0f, 0xfb, 0x4d, 0xda, 0xe7, 0x80, 0x07, 0x1b, 0x94,
	0xe4, 0x3b, 0x8c, 0xab, 0x4d, 0x27, 0x47, 0x2d, 0xc7, 0xae, 0x7f, 0x7d, 0x9c, 0xf9, 0x2e, 0xac,
	0x1c, 0xc8, 0xfd, 0x8f, 0x7c, 0x81, 0x41, 0x6e, 0x08, 0x79, 0xb0, 0xf3, 0x2a, 0x1c, 0xf7, 0x5f,
	0x48, 0x5d, 0xf0, 0x14, 0xac, 0x13, 0xd4, 0xe4, 0x7e, 0x0b, 0xee, 0x32, 0xcd, 0x99, 0x5d, 0x0f,
	0x54, 0xb5, 0x56, 0xa3, 0xe2, 0x53, 0xf1, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xc9,
	0x59, 0xad, 0x73, 0x04, 0x00, 0x00,
}
