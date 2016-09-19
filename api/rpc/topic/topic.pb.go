// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/topic/topic.proto
// DO NOT EDIT!

/*
Package topic is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/topic/topic.proto

It has these top-level messages:
	TopicEntry
	CreateRequest
	CreateReply
	ListRequest
	ListReply
	DeleteRequest
	DeleteReply
*/
package topic

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type TopicEntry struct {
	Id                string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Partitions        uint64 `protobuf:"varint,3,opt,name=partitions" json:"partitions,omitempty"`
	ReplicationFactor uint64 `protobuf:"varint,4,opt,name=replication_factor,json=replicationFactor" json:"replication_factor,omitempty"`
}

func (m *TopicEntry) Reset()                    { *m = TopicEntry{} }
func (m *TopicEntry) String() string            { return proto.CompactTextString(m) }
func (*TopicEntry) ProtoMessage()               {}
func (*TopicEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateRequest struct {
	Topic *TopicEntry `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetTopic() *TopicEntry {
	if m != nil {
		return m.Topic
	}
	return nil
}

type CreateReply struct {
	Topic *TopicEntry `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *CreateReply) Reset()                    { *m = CreateReply{} }
func (m *CreateReply) String() string            { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()               {}
func (*CreateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateReply) GetTopic() *TopicEntry {
	if m != nil {
		return m.Topic
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListReply struct {
	Topics []*TopicEntry `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListReply) GetTopics() []*TopicEntry {
	if m != nil {
		return m.Topics
	}
	return nil
}

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteReply struct {
}

func (m *DeleteReply) Reset()                    { *m = DeleteReply{} }
func (m *DeleteReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteReply) ProtoMessage()               {}
func (*DeleteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*TopicEntry)(nil), "topic.TopicEntry")
	proto.RegisterType((*CreateRequest)(nil), "topic.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "topic.CreateReply")
	proto.RegisterType((*ListRequest)(nil), "topic.ListRequest")
	proto.RegisterType((*ListReply)(nil), "topic.ListReply")
	proto.RegisterType((*DeleteRequest)(nil), "topic.DeleteRequest")
	proto.RegisterType((*DeleteReply)(nil), "topic.DeleteReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Topic service

type TopicClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error)
}

type topicClient struct {
	cc *grpc.ClientConn
}

func NewTopicClient(cc *grpc.ClientConn) TopicClient {
	return &topicClient{cc}
}

func (c *topicClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := grpc.Invoke(ctx, "/topic.Topic/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/topic.Topic/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := grpc.Invoke(ctx, "/topic.Topic/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Topic service

type TopicServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
}

func RegisterTopicServer(s *grpc.Server, srv TopicServer) {
	s.RegisterService(&_Topic_serviceDesc, srv)
}

func _Topic_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topic.Topic/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topic_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topic.Topic/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topic_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topic.Topic/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Topic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topic.Topic",
	HandlerType: (*TopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Topic_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Topic_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Topic_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/topic/topic.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x43, 0x4b, 0x49, 0x3a, 0x4d, 0x8d, 0x9d, 0x78, 0x20, 0x3d, 0x28, 0xe1, 0x62, 0x4d,
	0x14, 0x12, 0x8c, 0x8d, 0x77, 0xff, 0x9c, 0x3c, 0x11, 0xef, 0x66, 0x4b, 0x57, 0xdd, 0x04, 0xd8,
	0x75, 0x99, 0x1e, 0x38, 0xf9, 0x49, 0xfc, 0xae, 0x66, 0x17, 0x68, 0xa1, 0xf1, 0xe0, 0x85, 0xec,
	0xcc, 0xbc, 0x7d, 0xef, 0xc7, 0x64, 0xe1, 0xee, 0x43, 0xd0, 0xe7, 0x6e, 0x13, 0x65, 0xb2, 0x88,
	0x99, 0x52, 0x19, 0xcf, 0xb9, 0x66, 0x24, 0x75, 0xcc, 0x0a, 0x15, 0x33, 0x25, 0x62, 0xad, 0xb2,
	0x98, 0xa4, 0x12, 0xed, 0x37, 0x52, 0x5a, 0x92, 0xc4, 0x89, 0x2d, 0xc2, 0x6f, 0x80, 0x57, 0x73,
	0x78, 0x2a, 0x49, 0xd7, 0x78, 0x02, 0x23, 0xb1, 0xf5, 0x9d, 0xc0, 0x59, 0x4d, 0xd3, 0x91, 0xd8,
	0x22, 0x82, 0x5b, 0xb2, 0x82, 0xfb, 0x23, 0xdb, 0xb1, 0x67, 0x3c, 0x07, 0x50, 0x4c, 0x93, 0x20,
	0x21, 0xcb, 0xca, 0x1f, 0x07, 0xce, 0xca, 0x4d, 0x7b, 0x1d, 0xbc, 0x01, 0xd4, 0x5c, 0xe5, 0x22,
	0x63, 0xa6, 0x7e, 0x7b, 0x67, 0x19, 0x49, 0xed, 0xbb, 0x56, 0xb7, 0xe8, 0x4d, 0x9e, 0xed, 0x20,
	0xbc, 0x87, 0xf9, 0x83, 0xe6, 0x8c, 0x78, 0xca, 0xbf, 0x76, 0xbc, 0x22, 0xbc, 0x84, 0x06, 0xcd,
	0x62, 0xcc, 0x92, 0x45, 0xd4, 0x50, 0x1f, 0x28, 0xd3, 0x16, 0x7d, 0x0d, 0xb3, 0xee, 0xa6, 0xca,
	0xeb, 0xff, 0xdf, 0x9b, 0xc3, 0xec, 0x45, 0x54, 0xd4, 0xe6, 0x85, 0x6b, 0x98, 0x36, 0xa5, 0x31,
	0xb9, 0x02, 0xcf, 0x8a, 0x2a, 0xdf, 0x09, 0xc6, 0x7f, 0xbb, 0xb4, 0x82, 0xf0, 0x02, 0xe6, 0x8f,
	0x3c, 0xe7, 0x07, 0xf0, 0xa3, 0xe5, 0x99, 0x9c, 0x4e, 0xa0, 0xf2, 0x3a, 0xf9, 0x71, 0x60, 0x62,
	0x6d, 0x30, 0x01, 0xaf, 0x01, 0xc7, 0xb3, 0xd6, 0x7e, 0xb0, 0x81, 0x25, 0x1e, 0x75, 0x0d, 0xd8,
	0x35, 0xb8, 0x86, 0x12, 0xbb, 0x59, 0xef, 0x0f, 0x96, 0xa7, 0x83, 0x9e, 0x51, 0x27, 0xe0, 0x35,
	0xd1, 0xfb, 0x84, 0x01, 0xea, 0x3e, 0xa1, 0xc7, 0xb7, 0xf1, 0xec, 0xbb, 0xb8, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xe9, 0x3a, 0xf4, 0xf5, 0x50, 0x02, 0x00, 0x00,
}
