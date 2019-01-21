// Code generated by protoc-gen-go. DO NOT EDIT.
// source: more_rpc.proto

package moreRpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type Data struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_196f4518da60af7b, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Data) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CommitIDMsg struct {
	CommitID             string   `protobuf:"bytes,3,opt,name=CommitID,proto3" json:"CommitID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitIDMsg) Reset()         { *m = CommitIDMsg{} }
func (m *CommitIDMsg) String() string { return proto.CompactTextString(m) }
func (*CommitIDMsg) ProtoMessage()    {}
func (*CommitIDMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_196f4518da60af7b, []int{1}
}

func (m *CommitIDMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitIDMsg.Unmarshal(m, b)
}
func (m *CommitIDMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitIDMsg.Marshal(b, m, deterministic)
}
func (m *CommitIDMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitIDMsg.Merge(m, src)
}
func (m *CommitIDMsg) XXX_Size() int {
	return xxx_messageInfo_CommitIDMsg.Size(m)
}
func (m *CommitIDMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitIDMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CommitIDMsg proto.InternalMessageInfo

func (m *CommitIDMsg) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "moreRpc.Data")
	proto.RegisterType((*CommitIDMsg)(nil), "moreRpc.CommitIDMsg")
}

func init() { proto.RegisterFile("more_rpc.proto", fileDescriptor_196f4518da60af7b) }

var fileDescriptor_196f4518da60af7b = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcd, 0x2f, 0x4a,
	0x8d, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0xf1, 0x83, 0x0a,
	0x92, 0x95, 0xf4, 0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c,
	0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18, 0x84, 0xa3, 0xa4, 0xc9, 0xc5, 0xed, 0x9c, 0x9f, 0x9b, 0x9b,
	0x59, 0xe2, 0xe9, 0xe2, 0x5b, 0x9c, 0x2e, 0x24, 0xc5, 0xc5, 0x01, 0xe3, 0x4a, 0x30, 0x83, 0xd5,
	0xc1, 0xf9, 0x46, 0xd7, 0x18, 0xb9, 0x78, 0x7c, 0x21, 0xd6, 0x04, 0x80, 0x2d, 0xd5, 0xe2, 0xe2,
	0xf0, 0xcc, 0x73, 0x4f, 0x2d, 0xf1, 0x4e, 0xad, 0x14, 0xe2, 0xd5, 0x83, 0xba, 0x40, 0x0f, 0x64,
	0xbd, 0x14, 0x2a, 0x57, 0x89, 0x41, 0x48, 0x87, 0x8b, 0xcb, 0x33, 0x2f, 0x38, 0xb5, 0x24, 0x0c,
	0x64, 0x2b, 0x41, 0xd5, 0x66, 0x5c, 0x6c, 0x10, 0x6b, 0x85, 0x44, 0xe0, 0x52, 0x48, 0xce, 0x94,
	0xc2, 0x2a, 0xaa, 0xc4, 0x20, 0x64, 0xc2, 0xc5, 0xe2, 0x52, 0x94, 0x5f, 0x40, 0x9a, 0x2e, 0x27,
	0x0d, 0x2e, 0xb1, 0xcc, 0x7c, 0xbd, 0x74, 0x50, 0x68, 0xa6, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4,
	0x16, 0xeb, 0x55, 0x25, 0x26, 0x67, 0xa4, 0x3a, 0xa1, 0xf8, 0x37, 0x80, 0x31, 0x89, 0x0d, 0x1c,
	0xda, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x4e, 0xf9, 0x5f, 0x7f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoreRpcProtoClient is the client API for MoreRpcProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoreRpcProtoClient interface {
	// 获取单个value
	InGetKey(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	// 设置key/value
	InSetValue(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
	// 提交事务
	Commit(ctx context.Context, in *CommitIDMsg, opts ...grpc.CallOption) (*CommitIDMsg, error)
	// 丢弃事务
	Drop(ctx context.Context, in *CommitIDMsg, opts ...grpc.CallOption) (*CommitIDMsg, error)
}

type moreRpcProtoClient struct {
	cc *grpc.ClientConn
}

func NewMoreRpcProtoClient(cc *grpc.ClientConn) MoreRpcProtoClient {
	return &moreRpcProtoClient{cc}
}

func (c *moreRpcProtoClient) InGetKey(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/moreRpc.MoreRpcProto/InGetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moreRpcProtoClient) InSetValue(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/moreRpc.MoreRpcProto/InSetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moreRpcProtoClient) Commit(ctx context.Context, in *CommitIDMsg, opts ...grpc.CallOption) (*CommitIDMsg, error) {
	out := new(CommitIDMsg)
	err := c.cc.Invoke(ctx, "/moreRpc.MoreRpcProto/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moreRpcProtoClient) Drop(ctx context.Context, in *CommitIDMsg, opts ...grpc.CallOption) (*CommitIDMsg, error) {
	out := new(CommitIDMsg)
	err := c.cc.Invoke(ctx, "/moreRpc.MoreRpcProto/Drop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoreRpcProtoServer is the server API for MoreRpcProto service.
type MoreRpcProtoServer interface {
	// 获取单个value
	InGetKey(context.Context, *Data) (*Data, error)
	// 设置key/value
	InSetValue(context.Context, *Data) (*Data, error)
	// 提交事务
	Commit(context.Context, *CommitIDMsg) (*CommitIDMsg, error)
	// 丢弃事务
	Drop(context.Context, *CommitIDMsg) (*CommitIDMsg, error)
}

func RegisterMoreRpcProtoServer(s *grpc.Server, srv MoreRpcProtoServer) {
	s.RegisterService(&_MoreRpcProto_serviceDesc, srv)
}

func _MoreRpcProto_InGetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoreRpcProtoServer).InGetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moreRpc.MoreRpcProto/InGetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoreRpcProtoServer).InGetKey(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoreRpcProto_InSetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoreRpcProtoServer).InSetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moreRpc.MoreRpcProto/InSetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoreRpcProtoServer).InSetValue(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoreRpcProto_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitIDMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoreRpcProtoServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moreRpc.MoreRpcProto/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoreRpcProtoServer).Commit(ctx, req.(*CommitIDMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoreRpcProto_Drop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitIDMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoreRpcProtoServer).Drop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moreRpc.MoreRpcProto/Drop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoreRpcProtoServer).Drop(ctx, req.(*CommitIDMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _MoreRpcProto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moreRpc.MoreRpcProto",
	HandlerType: (*MoreRpcProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InGetKey",
			Handler:    _MoreRpcProto_InGetKey_Handler,
		},
		{
			MethodName: "InSetValue",
			Handler:    _MoreRpcProto_InSetValue_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _MoreRpcProto_Commit_Handler,
		},
		{
			MethodName: "Drop",
			Handler:    _MoreRpcProto_Drop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "more_rpc.proto",
}