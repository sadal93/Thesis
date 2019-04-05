// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Input struct {
	First                int32    `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second               int32    `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{0}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *Input) GetSecond() int32 {
	if m != nil {
		return m.Second
	}
	return 0
}

type Range struct {
	Begin                int32    `protobuf:"varint,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End                  int32    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{1}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetBegin() int32 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *Range) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

type Output struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{2}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Input)(nil), "api.Input")
	proto.RegisterType((*Range)(nil), "api.Range")
	proto.RegisterType((*Output)(nil), "api.Output")
}

func init() { proto.RegisterFile("sum.proto", fileDescriptor_62743f9cdc99b9fd) }

var fileDescriptor_62743f9cdc99b9fd = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0xab, 0xc2, 0x30,
	0x14, 0x85, 0x5f, 0x5f, 0x49, 0x78, 0xef, 0x76, 0x91, 0x20, 0xa5, 0x38, 0x95, 0x4c, 0xe2, 0x50,
	0x45, 0xf1, 0x07, 0x74, 0x74, 0x10, 0xa1, 0x2e, 0xae, 0xad, 0x89, 0xe5, 0x42, 0x9b, 0x84, 0x26,
	0xf9, 0xff, 0xd2, 0xb4, 0x75, 0x70, 0x3b, 0x5f, 0x38, 0x27, 0xe7, 0x5c, 0xf8, 0xb7, 0xbe, 0x2f,
	0xcc, 0xa0, 0x9d, 0x66, 0x71, 0x6d, 0x90, 0x9f, 0x81, 0x5c, 0x94, 0xf1, 0x8e, 0xad, 0x81, 0xbc,
	0x70, 0xb0, 0x2e, 0x8b, 0xf2, 0x68, 0x4b, 0xaa, 0x09, 0x58, 0x0a, 0xd4, 0xca, 0xa7, 0x56, 0x22,
	0xfb, 0x0d, 0xcf, 0x33, 0xf1, 0x3d, 0x90, 0xaa, 0x56, 0xad, 0x1c, 0x63, 0x8d, 0x6c, 0x51, 0x2d,
	0xb1, 0x00, 0x6c, 0x05, 0xb1, 0xfc, 0x64, 0x46, 0xc9, 0x73, 0xa0, 0x37, 0xef, 0xc6, 0xa2, 0x14,
	0xe8, 0x20, 0xad, 0xef, 0x96, 0xa6, 0x99, 0x8e, 0x0f, 0xf8, 0x2b, 0x85, 0x40, 0x87, 0x5a, 0xb1,
	0x1c, 0xe2, 0x52, 0x08, 0x06, 0x45, 0x6d, 0xb0, 0x08, 0xfb, 0x36, 0x49, 0xd0, 0xd3, 0x1f, 0xfc,
	0x87, 0xed, 0x20, 0xb9, 0xfa, 0xce, 0xa1, 0xe9, 0xe4, 0xdd, 0xf7, 0xb3, 0x33, 0x4c, 0xfa, 0x72,
	0x1e, 0xa2, 0x86, 0x86, 0x7b, 0x4f, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xbe, 0x87, 0x74,
	0xfc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdditionClient is the client API for Addition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdditionClient interface {
	Add(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
	MultipleSum(ctx context.Context, in *Range, opts ...grpc.CallOption) (Addition_MultipleSumClient, error)
}

type additionClient struct {
	cc *grpc.ClientConn
}

func NewAdditionClient(cc *grpc.ClientConn) AdditionClient {
	return &additionClient{cc}
}

func (c *additionClient) Add(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/api.Addition/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *additionClient) MultipleSum(ctx context.Context, in *Range, opts ...grpc.CallOption) (Addition_MultipleSumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Addition_serviceDesc.Streams[0], "/api.Addition/MultipleSum", opts...)
	if err != nil {
		return nil, err
	}
	x := &additionMultipleSumClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Addition_MultipleSumClient interface {
	Recv() (*Output, error)
	grpc.ClientStream
}

type additionMultipleSumClient struct {
	grpc.ClientStream
}

func (x *additionMultipleSumClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdditionServer is the server API for Addition service.
type AdditionServer interface {
	Add(context.Context, *Input) (*Output, error)
	MultipleSum(*Range, Addition_MultipleSumServer) error
}

// UnimplementedAdditionServer can be embedded to have forward compatible implementations.
type UnimplementedAdditionServer struct {
}

func (*UnimplementedAdditionServer) Add(ctx context.Context, req *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedAdditionServer) MultipleSum(req *Range, srv Addition_MultipleSumServer) error {
	return status.Errorf(codes.Unimplemented, "method MultipleSum not implemented")
}

func RegisterAdditionServer(s *grpc.Server, srv AdditionServer) {
	s.RegisterService(&_Addition_serviceDesc, srv)
}

func _Addition_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Addition/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServer).Add(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addition_MultipleSum_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Range)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdditionServer).MultipleSum(m, &additionMultipleSumServer{stream})
}

type Addition_MultipleSumServer interface {
	Send(*Output) error
	grpc.ServerStream
}

type additionMultipleSumServer struct {
	grpc.ServerStream
}

func (x *additionMultipleSumServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

var _Addition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Addition",
	HandlerType: (*AdditionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Addition_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MultipleSum",
			Handler:       _Addition_MultipleSum_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sum.proto",
}
