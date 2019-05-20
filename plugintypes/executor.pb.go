// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package plugintypes

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

type ExecuteRequest struct {
	JobName              string            `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	Config               map[string]string `protobuf:"bytes,2,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExecuteRequest) Reset()         { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()    {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_8f3a756833c7ef16, []int{0}
}
func (m *ExecuteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteRequest.Unmarshal(m, b)
}
func (m *ExecuteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteRequest.Marshal(b, m, deterministic)
}
func (dst *ExecuteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteRequest.Merge(dst, src)
}
func (m *ExecuteRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteRequest.Size(m)
}
func (m *ExecuteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteRequest proto.InternalMessageInfo

func (m *ExecuteRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *ExecuteRequest) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

type ExecuteResponse struct {
	Output               []byte   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteResponse) Reset()         { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()    {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_8f3a756833c7ef16, []int{1}
}
func (m *ExecuteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteResponse.Unmarshal(m, b)
}
func (m *ExecuteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteResponse.Marshal(b, m, deterministic)
}
func (dst *ExecuteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteResponse.Merge(dst, src)
}
func (m *ExecuteResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteResponse.Size(m)
}
func (m *ExecuteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteResponse proto.InternalMessageInfo

func (m *ExecuteResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *ExecuteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecuteRequest)(nil), "plugintypes.ExecuteRequest")
	proto.RegisterMapType((map[string]string)(nil), "plugintypes.ExecuteRequest.ConfigEntry")
	proto.RegisterType((*ExecuteResponse)(nil), "plugintypes.ExecuteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/plugintypes.Executor/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugintypes.Executor/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugintypes.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Executor_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executor.proto",
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_executor_8f3a756833c7ef16) }

var fileDescriptor_executor_8f3a756833c7ef16 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xad, 0x48, 0x4d,
	0x2e, 0x2d, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x5a, 0xce, 0xc8, 0xc5, 0xe7, 0x0a, 0x96, 0x4f,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe4, 0xe2, 0xc8, 0xca, 0x4f, 0x8a, 0xcf,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xcf, 0xca, 0x4f, 0xf2, 0x4b,
	0xcc, 0x4d, 0x15, 0xb2, 0xe7, 0x62, 0x4b, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x97, 0x60, 0x52, 0x60,
	0xd6, 0xe0, 0x36, 0x52, 0xd7, 0x43, 0x32, 0x4b, 0x0f, 0xd5, 0x1c, 0x3d, 0x67, 0xb0, 0x4a, 0xd7,
	0xbc, 0x92, 0xa2, 0xca, 0x20, 0xa8, 0x36, 0x29, 0x4b, 0x2e, 0x6e, 0x24, 0x61, 0x21, 0x01, 0x2e,
	0xe6, 0xec, 0xd4, 0x4a, 0xa8, 0x2d, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69,
	0xaa, 0x04, 0x13, 0x58, 0x0c, 0xc2, 0xb1, 0x62, 0xb2, 0x60, 0x54, 0xb2, 0xe7, 0xe2, 0x87, 0x5b,
	0x50, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x96, 0x5f, 0x5a, 0x52, 0x50, 0x5a,
	0x02, 0x36, 0x81, 0x27, 0x08, 0xca, 0x03, 0x19, 0x92, 0x5a, 0x54, 0x94, 0x5f, 0x04, 0x33, 0x04,
	0xcc, 0x31, 0x0a, 0xe0, 0xe2, 0x70, 0x85, 0x86, 0x84, 0x90, 0x0b, 0x17, 0x3b, 0xd4, 0x30, 0x21,
	0x69, 0x3c, 0x7e, 0x90, 0x92, 0xc1, 0x2e, 0x09, 0xb1, 0x3f, 0x89, 0x0d, 0x1c, 0xa0, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xe2, 0x5b, 0xbf, 0x62, 0x01, 0x00, 0x00,
}
