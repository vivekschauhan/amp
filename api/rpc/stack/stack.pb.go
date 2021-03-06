// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/stack/stack.proto
// DO NOT EDIT!

/*
Package stack is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/stack/stack.proto

It has these top-level messages:
	UpRequest
	UpReply
	StackRequest
	RemoveRequest
	StackReply
	StackID
	Stack
	ListRequest
	ServiceIdList
	ListReply
	StackInfo
*/
package stack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import service "github.com/appcelerator/amp/api/rpc/service"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StackState int32

const (
	StackState_Stopped     StackState = 0
	StackState_Starting    StackState = 1
	StackState_Running     StackState = 2
	StackState_Redeploying StackState = 3
)

var StackState_name = map[int32]string{
	0: "Stopped",
	1: "Starting",
	2: "Running",
	3: "Redeploying",
}
var StackState_value = map[string]int32{
	"Stopped":     0,
	"Starting":    1,
	"Running":     2,
	"Redeploying": 3,
}

func (x StackState) String() string {
	return proto.EnumName(StackState_name, int32(x))
}
func (StackState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UpRequest struct {
	StackName string `protobuf:"bytes,1,opt,name=stack_name,json=stackName" json:"stack_name,omitempty"`
	Stackfile string `protobuf:"bytes,2,opt,name=stackfile" json:"stackfile,omitempty"`
}

func (m *UpRequest) Reset()                    { *m = UpRequest{} }
func (m *UpRequest) String() string            { return proto.CompactTextString(m) }
func (*UpRequest) ProtoMessage()               {}
func (*UpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UpReply struct {
	StackId string `protobuf:"bytes,1,opt,name=stack_id,json=stackId" json:"stack_id,omitempty"`
}

func (m *UpReply) Reset()                    { *m = UpReply{} }
func (m *UpReply) String() string            { return proto.CompactTextString(m) }
func (*UpReply) ProtoMessage()               {}
func (*UpReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StackRequest struct {
	StackIdent string `protobuf:"bytes,1,opt,name=stack_ident,json=stackIdent" json:"stack_ident,omitempty"`
}

func (m *StackRequest) Reset()                    { *m = StackRequest{} }
func (m *StackRequest) String() string            { return proto.CompactTextString(m) }
func (*StackRequest) ProtoMessage()               {}
func (*StackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RemoveRequest struct {
	StackIdent string `protobuf:"bytes,1,opt,name=stack_ident,json=stackIdent" json:"stack_ident,omitempty"`
	Force      bool   `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StackReply struct {
	StackId string `protobuf:"bytes,1,opt,name=stack_id,json=stackId" json:"stack_id,omitempty"`
}

func (m *StackReply) Reset()                    { *m = StackReply{} }
func (m *StackReply) String() string            { return proto.CompactTextString(m) }
func (*StackReply) ProtoMessage()               {}
func (*StackReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type StackID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StackID) Reset()                    { *m = StackID{} }
func (m *StackID) String() string            { return proto.CompactTextString(m) }
func (*StackID) ProtoMessage()               {}
func (*StackID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Stack struct {
	Name     string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id       string                 `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Services []*service.ServiceSpec `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
}

func (m *Stack) Reset()                    { *m = Stack{} }
func (m *Stack) String() string            { return proto.CompactTextString(m) }
func (*Stack) ProtoMessage()               {}
func (*Stack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Stack) GetServices() []*service.ServiceSpec {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ServiceIdList struct {
	List []string `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ServiceIdList) Reset()                    { *m = ServiceIdList{} }
func (m *ServiceIdList) String() string            { return proto.CompactTextString(m) }
func (*ServiceIdList) ProtoMessage()               {}
func (*ServiceIdList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListReply struct {
	List []*StackInfo `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListReply) GetList() []*StackInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type StackInfo struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	State string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *StackInfo) Reset()                    { *m = StackInfo{} }
func (m *StackInfo) String() string            { return proto.CompactTextString(m) }
func (*StackInfo) ProtoMessage()               {}
func (*StackInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*UpRequest)(nil), "stack.UpRequest")
	proto.RegisterType((*UpReply)(nil), "stack.UpReply")
	proto.RegisterType((*StackRequest)(nil), "stack.StackRequest")
	proto.RegisterType((*RemoveRequest)(nil), "stack.removeRequest")
	proto.RegisterType((*StackReply)(nil), "stack.StackReply")
	proto.RegisterType((*StackID)(nil), "stack.StackID")
	proto.RegisterType((*Stack)(nil), "stack.Stack")
	proto.RegisterType((*ListRequest)(nil), "stack.ListRequest")
	proto.RegisterType((*ServiceIdList)(nil), "stack.ServiceIdList")
	proto.RegisterType((*ListReply)(nil), "stack.ListReply")
	proto.RegisterType((*StackInfo)(nil), "stack.StackInfo")
	proto.RegisterEnum("stack.StackState", StackState_name, StackState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for StackService service

type StackServiceClient interface {
	Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpReply, error)
	Start(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error)
	Stop(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*StackReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type stackServiceClient struct {
	cc *grpc.ClientConn
}

func NewStackServiceClient(cc *grpc.ClientConn) StackServiceClient {
	return &stackServiceClient{cc}
}

func (c *stackServiceClient) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpReply, error) {
	out := new(UpReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Up", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Start(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Stop(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/stack.StackService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StackService service

type StackServiceServer interface {
	Up(context.Context, *UpRequest) (*UpReply, error)
	Start(context.Context, *StackRequest) (*StackReply, error)
	Stop(context.Context, *StackRequest) (*StackReply, error)
	Remove(context.Context, *RemoveRequest) (*StackReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
}

func RegisterStackServiceServer(s *grpc.Server, srv StackServiceServer) {
	s.RegisterService(&_StackService_serviceDesc, srv)
}

func _StackService_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Up(ctx, req.(*UpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Start(ctx, req.(*StackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Stop(ctx, req.(*StackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stack.StackService",
	HandlerType: (*StackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Up",
			Handler:    _StackService_Up_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _StackService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _StackService_Stop_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _StackService_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StackService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/stack/stack.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xa6, 0x49, 0xbb, 0x36, 0x97, 0x75, 0x14, 0xd3, 0x87, 0xae, 0x02, 0x31, 0x85, 0x09, 0x26,
	0x84, 0x12, 0x56, 0xc4, 0x03, 0xef, 0x0c, 0x51, 0x09, 0xf1, 0x90, 0x68, 0x8f, 0x08, 0x65, 0xc9,
	0x75, 0x58, 0x24, 0xb1, 0x71, 0xdc, 0x49, 0xfd, 0xcf, 0x79, 0x44, 0x67, 0x3b, 0x55, 0x2a, 0x21,
	0xd1, 0x97, 0xd8, 0xf7, 0xdd, 0x77, 0x3f, 0xbe, 0xf3, 0x05, 0x3e, 0xdc, 0x73, 0xfd, 0x73, 0x7b,
	0x17, 0x17, 0xa2, 0x4e, 0x72, 0x29, 0x0b, 0xac, 0x50, 0xe5, 0x5a, 0xa8, 0x24, 0xaf, 0x65, 0x92,
	0x4b, 0x9e, 0x28, 0x59, 0x24, 0xad, 0xce, 0x8b, 0x5f, 0xf6, 0x1b, 0x4b, 0x25, 0xb4, 0x60, 0x23,
	0x63, 0x2c, 0x3f, 0x1e, 0x15, 0x8d, 0xea, 0x81, 0x17, 0xd8, 0x9d, 0x36, 0x43, 0xf4, 0x05, 0x82,
	0x5b, 0x99, 0xe2, 0xef, 0x2d, 0xb6, 0x9a, 0x3d, 0x07, 0x30, 0x09, 0x7f, 0x34, 0x79, 0x8d, 0x8b,
	0xc1, 0xc5, 0xe0, 0x2a, 0x48, 0x03, 0x83, 0x7c, 0xcb, 0x6b, 0x64, 0xcf, 0xc0, 0x1a, 0x1b, 0x5e,
	0xe1, 0xc2, 0xeb, 0x79, 0x09, 0x88, 0x2e, 0x61, 0x4c, 0x99, 0x64, 0xb5, 0x63, 0xe7, 0x30, 0xb1,
	0x79, 0x78, 0xe9, 0xb2, 0x8c, 0x8d, 0xbd, 0x2e, 0xa3, 0x04, 0x4e, 0x33, 0xba, 0x76, 0x25, 0x5f,
	0x40, 0xd8, 0x51, 0xb1, 0xd1, 0x8e, 0x0d, 0x8e, 0x8d, 0x8d, 0x8e, 0x3e, 0xc3, 0x54, 0x61, 0x2d,
	0x1e, 0xf0, 0xd8, 0x08, 0x36, 0x87, 0xd1, 0x46, 0xa8, 0xc2, 0xb6, 0x38, 0x49, 0xad, 0x11, 0xbd,
	0x06, 0x70, 0x85, 0xff, 0xd3, 0xe1, 0x39, 0x8c, 0x0d, 0x71, 0xfd, 0x89, 0x9d, 0x81, 0xb7, 0xf7,
	0x7b, 0xbc, 0x8c, 0xbe, 0xc3, 0xc8, 0xb8, 0x18, 0x83, 0x61, 0x6f, 0x44, 0xe6, 0xee, 0xc8, 0x5e,
	0x47, 0x66, 0xef, 0x60, 0xe2, 0x46, 0xdd, 0x2e, 0xfc, 0x0b, 0xff, 0x2a, 0x5c, 0xcd, 0xe3, 0x6e,
	0xf6, 0x99, 0x3d, 0x33, 0x89, 0x45, 0xba, 0x67, 0x45, 0x53, 0x08, 0xbf, 0xf2, 0x56, 0x3b, 0xa1,
	0xd1, 0x4b, 0x98, 0x3a, 0xde, 0xba, 0x24, 0x9c, 0xaa, 0x56, 0xbc, 0x25, 0xc9, 0x3e, 0x55, 0xa5,
	0x7b, 0x74, 0x0d, 0x81, 0x8d, 0x21, 0x55, 0x97, 0x3d, 0x42, 0xb8, 0x9a, 0xc5, 0x76, 0x55, 0xac,
	0x9a, 0x66, 0x23, 0x5c, 0xc8, 0x0d, 0x04, 0x7b, 0xe8, 0x28, 0x25, 0x73, 0xa0, 0x3d, 0xd3, 0xb8,
	0xf0, 0x0d, 0x64, 0x8d, 0x37, 0x37, 0x6e, 0xa0, 0x19, 0x59, 0x2c, 0xa4, 0xa9, 0x09, 0x29, 0xb1,
	0x9c, 0x3d, 0x62, 0xa7, 0x30, 0xc9, 0x74, 0xae, 0x34, 0x6f, 0xee, 0x67, 0x03, 0x72, 0xa5, 0xdb,
	0xa6, 0x21, 0xc3, 0x63, 0x8f, 0x21, 0x4c, 0xb1, 0x44, 0x59, 0x89, 0x1d, 0x01, 0xfe, 0xea, 0xcf,
	0xc0, 0x6d, 0x84, 0xd3, 0xca, 0x5e, 0x81, 0x77, 0x2b, 0x59, 0xd7, 0xfc, 0x7e, 0x39, 0x97, 0x67,
	0x3d, 0x84, 0xc4, 0x26, 0xe6, 0x31, 0x94, 0x66, 0x4f, 0xfb, 0x3a, 0x3b, 0xf6, 0x93, 0x43, 0x90,
	0x02, 0x62, 0x18, 0x52, 0x8b, 0x47, 0xf3, 0xaf, 0xe1, 0x24, 0x35, 0x9b, 0xc7, 0xe6, 0xce, 0x79,
	0xb0, 0x88, 0xff, 0x0a, 0x79, 0x0b, 0x43, 0xfb, 0x52, 0xce, 0xd5, 0x7b, 0xce, 0xe5, 0xec, 0x00,
	0x93, 0xd5, 0xee, 0xee, 0xc4, 0xfc, 0x82, 0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x1a,
	0x24, 0xb1, 0xfd, 0x03, 0x00, 0x00,
}
