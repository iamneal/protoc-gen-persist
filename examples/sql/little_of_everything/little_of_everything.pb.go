// Code generated by protoc-gen-go.
// source: examples/sql/little_of_everything/little_of_everything.proto
// DO NOT EDIT!

/*
Package little_of_everything is a generated protocol buffer package.

It is generated from these files:
	examples/sql/little_of_everything/little_of_everything.proto

It has these top-level messages:
	ExampleTable1
	ExternalTypeTestMessage
	CountRows
*/
package little_of_everything

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/tcncloud/protoc-gen-persist/persist"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/protoc-gen-go/descriptor"
import examples_test "github.com/tcncloud/protoc-gen-persist/examples/test"

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

type TestEnum int32

const (
	TestEnum_OPTION_0 TestEnum = 0
	TestEnum_OPTION_1 TestEnum = 1
)

var TestEnum_name = map[int32]string{
	0: "OPTION_0",
	1: "OPTION_1",
}
var TestEnum_value = map[string]int32{
	"OPTION_0": 0,
	"OPTION_1": 1,
}

func (x TestEnum) String() string {
	return proto.EnumName(TestEnum_name, int32(x))
}
func (TestEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ExampleTable1_InnerEnum int32

const (
	ExampleTable1_OP_0 ExampleTable1_InnerEnum = 0
	ExampleTable1_OP_1 ExampleTable1_InnerEnum = 1
)

var ExampleTable1_InnerEnum_name = map[int32]string{
	0: "OP_0",
	1: "OP_1",
}
var ExampleTable1_InnerEnum_value = map[string]int32{
	"OP_0": 0,
	"OP_1": 1,
}

func (x ExampleTable1_InnerEnum) String() string {
	return proto.EnumName(ExampleTable1_InnerEnum_name, int32(x))
}
func (ExampleTable1_InnerEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ExampleTable1 struct {
	TableId      int32                       `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	Key          string                      `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value        string                      `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	InnerMessage *ExampleTable1_InnerMessage `protobuf:"bytes,4,opt,name=inner_message,json=innerMessage" json:"inner_message,omitempty"`
	InnerEnum    ExampleTable1_InnerEnum     `protobuf:"varint,5,opt,name=inner_enum,json=innerEnum,enum=examples.ExampleTable1_InnerEnum" json:"inner_enum,omitempty"`
	StringArray  []string                    `protobuf:"bytes,6,rep,name=string_array,json=stringArray" json:"string_array,omitempty"`
	BytesField   []byte                      `protobuf:"bytes,7,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	StartTime    *google_protobuf1.Timestamp `protobuf:"bytes,10,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	TestField    *examples_test.Test         `protobuf:"bytes,100,opt,name=test_field,json=testField" json:"test_field,omitempty"`
	// Types that are valid to be assigned to SkipOneOf:
	//	*ExampleTable1_TestId
	//	*ExampleTable1_TestValue
	//	*ExampleTable1_TestOutMessage
	SkipOneOf isExampleTable1_SkipOneOf `protobuf_oneof:"skip_one_of"`
}

func (m *ExampleTable1) Reset()                    { *m = ExampleTable1{} }
func (m *ExampleTable1) String() string            { return proto.CompactTextString(m) }
func (*ExampleTable1) ProtoMessage()               {}
func (*ExampleTable1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isExampleTable1_SkipOneOf interface {
	isExampleTable1_SkipOneOf()
}

type ExampleTable1_TestId struct {
	TestId int32 `protobuf:"varint,1001,opt,name=test_id,json=testId,oneof"`
}
type ExampleTable1_TestValue struct {
	TestValue string `protobuf:"bytes,1002,opt,name=test_value,json=testValue,oneof"`
}
type ExampleTable1_TestOutMessage struct {
	TestOutMessage *examples_test.Test `protobuf:"bytes,1003,opt,name=test_out_message,json=testOutMessage,oneof"`
}

func (*ExampleTable1_TestId) isExampleTable1_SkipOneOf()         {}
func (*ExampleTable1_TestValue) isExampleTable1_SkipOneOf()      {}
func (*ExampleTable1_TestOutMessage) isExampleTable1_SkipOneOf() {}

func (m *ExampleTable1) GetSkipOneOf() isExampleTable1_SkipOneOf {
	if m != nil {
		return m.SkipOneOf
	}
	return nil
}

func (m *ExampleTable1) GetTableId() int32 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *ExampleTable1) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ExampleTable1) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ExampleTable1) GetInnerMessage() *ExampleTable1_InnerMessage {
	if m != nil {
		return m.InnerMessage
	}
	return nil
}

func (m *ExampleTable1) GetInnerEnum() ExampleTable1_InnerEnum {
	if m != nil {
		return m.InnerEnum
	}
	return ExampleTable1_OP_0
}

func (m *ExampleTable1) GetStringArray() []string {
	if m != nil {
		return m.StringArray
	}
	return nil
}

func (m *ExampleTable1) GetBytesField() []byte {
	if m != nil {
		return m.BytesField
	}
	return nil
}

func (m *ExampleTable1) GetStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ExampleTable1) GetTestField() *examples_test.Test {
	if m != nil {
		return m.TestField
	}
	return nil
}

func (m *ExampleTable1) GetTestId() int32 {
	if x, ok := m.GetSkipOneOf().(*ExampleTable1_TestId); ok {
		return x.TestId
	}
	return 0
}

func (m *ExampleTable1) GetTestValue() string {
	if x, ok := m.GetSkipOneOf().(*ExampleTable1_TestValue); ok {
		return x.TestValue
	}
	return ""
}

func (m *ExampleTable1) GetTestOutMessage() *examples_test.Test {
	if x, ok := m.GetSkipOneOf().(*ExampleTable1_TestOutMessage); ok {
		return x.TestOutMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExampleTable1) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExampleTable1_OneofMarshaler, _ExampleTable1_OneofUnmarshaler, _ExampleTable1_OneofSizer, []interface{}{
		(*ExampleTable1_TestId)(nil),
		(*ExampleTable1_TestValue)(nil),
		(*ExampleTable1_TestOutMessage)(nil),
	}
}

func _ExampleTable1_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExampleTable1)
	// skip_one_of
	switch x := m.SkipOneOf.(type) {
	case *ExampleTable1_TestId:
		b.EncodeVarint(1001<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.TestId))
	case *ExampleTable1_TestValue:
		b.EncodeVarint(1002<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TestValue)
	case *ExampleTable1_TestOutMessage:
		b.EncodeVarint(1003<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TestOutMessage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ExampleTable1.SkipOneOf has unexpected type %T", x)
	}
	return nil
}

func _ExampleTable1_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExampleTable1)
	switch tag {
	case 1001: // skip_one_of.test_id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.SkipOneOf = &ExampleTable1_TestId{int32(x)}
		return true, err
	case 1002: // skip_one_of.test_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.SkipOneOf = &ExampleTable1_TestValue{x}
		return true, err
	case 1003: // skip_one_of.test_out_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(examples_test.Test)
		err := b.DecodeMessage(msg)
		m.SkipOneOf = &ExampleTable1_TestOutMessage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ExampleTable1_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExampleTable1)
	// skip_one_of
	switch x := m.SkipOneOf.(type) {
	case *ExampleTable1_TestId:
		n += proto.SizeVarint(1001<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.TestId))
	case *ExampleTable1_TestValue:
		n += proto.SizeVarint(1002<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TestValue)))
		n += len(x.TestValue)
	case *ExampleTable1_TestOutMessage:
		s := proto.Size(x.TestOutMessage)
		n += proto.SizeVarint(1003<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ExampleTable1_InnerMessage struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ExampleTable1_InnerMessage) Reset()                    { *m = ExampleTable1_InnerMessage{} }
func (m *ExampleTable1_InnerMessage) String() string            { return proto.CompactTextString(m) }
func (*ExampleTable1_InnerMessage) ProtoMessage()               {}
func (*ExampleTable1_InnerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ExampleTable1_InnerMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ExternalTypeTestMessage struct {
	Time *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
}

func (m *ExternalTypeTestMessage) Reset()                    { *m = ExternalTypeTestMessage{} }
func (m *ExternalTypeTestMessage) String() string            { return proto.CompactTextString(m) }
func (*ExternalTypeTestMessage) ProtoMessage()               {}
func (*ExternalTypeTestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExternalTypeTestMessage) GetTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type CountRows struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *CountRows) Reset()                    { *m = CountRows{} }
func (m *CountRows) String() string            { return proto.CompactTextString(m) }
func (*CountRows) ProtoMessage()               {}
func (*CountRows) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CountRows) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*ExampleTable1)(nil), "examples.ExampleTable1")
	proto.RegisterType((*ExampleTable1_InnerMessage)(nil), "examples.ExampleTable1.InnerMessage")
	proto.RegisterType((*ExternalTypeTestMessage)(nil), "examples.ExternalTypeTestMessage")
	proto.RegisterType((*CountRows)(nil), "examples.CountRows")
	proto.RegisterEnum("examples.TestEnum", TestEnum_name, TestEnum_value)
	proto.RegisterEnum("examples.ExampleTable1_InnerEnum", ExampleTable1_InnerEnum_name, ExampleTable1_InnerEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ExampleService1 service

type ExampleService1Client interface {
	UnaryExample1(ctx context.Context, in *ExampleTable1, opts ...grpc.CallOption) (*ExampleTable1, error)
	UnaryExample2(ctx context.Context, in *examples_test.Test, opts ...grpc.CallOption) (*ExampleTable1, error)
	ServerStreamSelect(ctx context.Context, in *ExampleTable1, opts ...grpc.CallOption) (ExampleService1_ServerStreamSelectClient, error)
	ClientStreamingExample(ctx context.Context, opts ...grpc.CallOption) (ExampleService1_ClientStreamingExampleClient, error)
}

type exampleService1Client struct {
	cc *grpc.ClientConn
}

func NewExampleService1Client(cc *grpc.ClientConn) ExampleService1Client {
	return &exampleService1Client{cc}
}

func (c *exampleService1Client) UnaryExample1(ctx context.Context, in *ExampleTable1, opts ...grpc.CallOption) (*ExampleTable1, error) {
	out := new(ExampleTable1)
	err := grpc.Invoke(ctx, "/examples.ExampleService1/UnaryExample1", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService1Client) UnaryExample2(ctx context.Context, in *examples_test.Test, opts ...grpc.CallOption) (*ExampleTable1, error) {
	out := new(ExampleTable1)
	err := grpc.Invoke(ctx, "/examples.ExampleService1/UnaryExample2", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService1Client) ServerStreamSelect(ctx context.Context, in *ExampleTable1, opts ...grpc.CallOption) (ExampleService1_ServerStreamSelectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ExampleService1_serviceDesc.Streams[0], c.cc, "/examples.ExampleService1/ServerStreamSelect", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleService1ServerStreamSelectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService1_ServerStreamSelectClient interface {
	Recv() (*ExampleTable1, error)
	grpc.ClientStream
}

type exampleService1ServerStreamSelectClient struct {
	grpc.ClientStream
}

func (x *exampleService1ServerStreamSelectClient) Recv() (*ExampleTable1, error) {
	m := new(ExampleTable1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleService1Client) ClientStreamingExample(ctx context.Context, opts ...grpc.CallOption) (ExampleService1_ClientStreamingExampleClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ExampleService1_serviceDesc.Streams[1], c.cc, "/examples.ExampleService1/ClientStreamingExample", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleService1ClientStreamingExampleClient{stream}
	return x, nil
}

type ExampleService1_ClientStreamingExampleClient interface {
	Send(*ExampleTable1) error
	CloseAndRecv() (*CountRows, error)
	grpc.ClientStream
}

type exampleService1ClientStreamingExampleClient struct {
	grpc.ClientStream
}

func (x *exampleService1ClientStreamingExampleClient) Send(m *ExampleTable1) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleService1ClientStreamingExampleClient) CloseAndRecv() (*CountRows, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CountRows)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ExampleService1 service

type ExampleService1Server interface {
	UnaryExample1(context.Context, *ExampleTable1) (*ExampleTable1, error)
	UnaryExample2(context.Context, *examples_test.Test) (*ExampleTable1, error)
	ServerStreamSelect(*ExampleTable1, ExampleService1_ServerStreamSelectServer) error
	ClientStreamingExample(ExampleService1_ClientStreamingExampleServer) error
}

func RegisterExampleService1Server(s *grpc.Server, srv ExampleService1Server) {
	s.RegisterService(&_ExampleService1_serviceDesc, srv)
}

func _ExampleService1_UnaryExample1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleTable1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleService1Server).UnaryExample1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.ExampleService1/UnaryExample1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleService1Server).UnaryExample1(ctx, req.(*ExampleTable1))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService1_UnaryExample2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(examples_test.Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleService1Server).UnaryExample2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.ExampleService1/UnaryExample2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleService1Server).UnaryExample2(ctx, req.(*examples_test.Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService1_ServerStreamSelect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExampleTable1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleService1Server).ServerStreamSelect(m, &exampleService1ServerStreamSelectServer{stream})
}

type ExampleService1_ServerStreamSelectServer interface {
	Send(*ExampleTable1) error
	grpc.ServerStream
}

type exampleService1ServerStreamSelectServer struct {
	grpc.ServerStream
}

func (x *exampleService1ServerStreamSelectServer) Send(m *ExampleTable1) error {
	return x.ServerStream.SendMsg(m)
}

func _ExampleService1_ClientStreamingExample_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleService1Server).ClientStreamingExample(&exampleService1ClientStreamingExampleServer{stream})
}

type ExampleService1_ClientStreamingExampleServer interface {
	SendAndClose(*CountRows) error
	Recv() (*ExampleTable1, error)
	grpc.ServerStream
}

type exampleService1ClientStreamingExampleServer struct {
	grpc.ServerStream
}

func (x *exampleService1ClientStreamingExampleServer) SendAndClose(m *CountRows) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleService1ClientStreamingExampleServer) Recv() (*ExampleTable1, error) {
	m := new(ExampleTable1)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ExampleService1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "examples.ExampleService1",
	HandlerType: (*ExampleService1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryExample1",
			Handler:    _ExampleService1_UnaryExample1_Handler,
		},
		{
			MethodName: "UnaryExample2",
			Handler:    _ExampleService1_UnaryExample2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamSelect",
			Handler:       _ExampleService1_ServerStreamSelect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingExample",
			Handler:       _ExampleService1_ClientStreamingExample_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "examples/sql/little_of_everything/little_of_everything.proto",
}

// Client API for NotEnabledService service

type NotEnabledServiceClient interface {
}

type notEnabledServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotEnabledServiceClient(cc *grpc.ClientConn) NotEnabledServiceClient {
	return &notEnabledServiceClient{cc}
}

// Server API for NotEnabledService service

type NotEnabledServiceServer interface {
}

func RegisterNotEnabledServiceServer(s *grpc.Server, srv NotEnabledServiceServer) {
	s.RegisterService(&_NotEnabledService_serviceDesc, srv)
}

var _NotEnabledService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "examples.NotEnabledService",
	HandlerType: (*NotEnabledServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "examples/sql/little_of_everything/little_of_everything.proto",
}

func init() {
	proto.RegisterFile("examples/sql/little_of_everything/little_of_everything.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 873 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xf6, 0xda, 0xf9, 0xb1, 0x8f, 0x9d, 0x60, 0xa6, 0x40, 0x17, 0x83, 0xc8, 0x64, 0x55, 0x21,
	0xab, 0x4a, 0xd7, 0xb5, 0x91, 0x90, 0x10, 0x20, 0x35, 0x89, 0x5c, 0xd9, 0x12, 0xa9, 0xab, 0xb5,
	0x01, 0xa9, 0x12, 0x2c, 0x6b, 0xef, 0xf1, 0x66, 0xc8, 0xfe, 0x98, 0x9d, 0xd9, 0x90, 0xbd, 0xe1,
	0x02, 0x89, 0x1b, 0xb8, 0x21, 0x0f, 0xc0, 0x73, 0x90, 0x47, 0xe0, 0x02, 0x09, 0x89, 0x47, 0xe0,
	0x0e, 0xb8, 0xe5, 0x01, 0xd0, 0xcc, 0xae, 0x9d, 0xb4, 0xb2, 0x5b, 0x81, 0xaa, 0xe6, 0x66, 0x35,
	0xe7, 0xff, 0x7c, 0xdf, 0x39, 0xb3, 0x03, 0x1f, 0xe0, 0x99, 0x13, 0xcc, 0x7c, 0xe4, 0x2d, 0xfe,
	0x95, 0xdf, 0xf2, 0x99, 0x10, 0x3e, 0xda, 0xd1, 0xd4, 0xc6, 0x53, 0x8c, 0x53, 0x71, 0xcc, 0x42,
	0x6f, 0xa9, 0xd2, 0x9c, 0xc5, 0x91, 0x88, 0x48, 0x79, 0x1e, 0xdd, 0x78, 0x75, 0x86, 0x31, 0x67,
	0x5c, 0xb4, 0xa2, 0x99, 0x60, 0x51, 0xc8, 0x33, 0x87, 0xc6, 0x8e, 0x17, 0x45, 0x9e, 0x8f, 0x2d,
	0x25, 0x8d, 0x93, 0x69, 0x4b, 0xb0, 0x00, 0xb9, 0x70, 0x82, 0x59, 0xee, 0x40, 0x9f, 0x74, 0x70,
	0x91, 0x4f, 0x62, 0x36, 0x13, 0x51, 0x9c, 0x7b, 0xe8, 0x8b, 0x0e, 0x05, 0x72, 0xa1, 0x3e, 0x99,
	0xc5, 0xf8, 0x63, 0x0d, 0xb6, 0xba, 0x99, 0x71, 0xe4, 0x8c, 0x7d, 0x6c, 0x93, 0xd7, 0xa1, 0x2c,
	0xe4, 0xc9, 0x66, 0xae, 0xae, 0x51, 0xad, 0xb9, 0x6e, 0x6d, 0x2a, 0xb9, 0xef, 0x92, 0x3a, 0x94,
	0x4e, 0x30, 0xd5, 0x8b, 0x54, 0x6b, 0x56, 0x2c, 0x79, 0x24, 0xaf, 0xc0, 0xfa, 0xa9, 0xe3, 0x27,
	0xa8, 0x97, 0x94, 0x2e, 0x13, 0x48, 0x1f, 0xb6, 0x58, 0x18, 0x62, 0x6c, 0x07, 0xc8, 0xb9, 0xe3,
	0xa1, 0xbe, 0x46, 0xb5, 0x66, 0xb5, 0x73, 0xcb, 0x9c, 0xb7, 0x61, 0x3e, 0x56, 0xd2, 0xec, 0x4b,
	0xe7, 0xa3, 0xcc, 0xd7, 0xaa, 0xb1, 0x2b, 0x12, 0xb9, 0x07, 0x90, 0xa5, 0xc2, 0x30, 0x09, 0xf4,
	0x75, 0xaa, 0x35, 0xb7, 0x3b, 0xbb, 0x4f, 0xcd, 0xd3, 0x0d, 0x93, 0xc0, 0xaa, 0xb0, 0xf9, 0x91,
	0xec, 0x42, 0x8d, 0x8b, 0x98, 0x85, 0x9e, 0xed, 0xc4, 0xb1, 0x93, 0xea, 0x1b, 0xb4, 0xd4, 0xac,
	0x58, 0xd5, 0x4c, 0xb7, 0x2f, 0x55, 0x64, 0x07, 0xaa, 0xe3, 0x54, 0x20, 0xb7, 0xa7, 0x0c, 0x7d,
	0x57, 0xdf, 0xa4, 0x5a, 0xb3, 0x66, 0x81, 0x52, 0xdd, 0x97, 0x1a, 0xf2, 0x1e, 0x00, 0x17, 0x4e,
	0x2c, 0x6c, 0x49, 0xbd, 0x0e, 0x0a, 0x4d, 0xc3, 0xcc, 0x68, 0x37, 0xe7, 0xb4, 0x9b, 0xa3, 0xf9,
	0x5c, 0xac, 0x8a, 0xf2, 0x96, 0x32, 0xe9, 0x00, 0x48, 0xba, 0xf3, 0xd4, 0xae, 0x0a, 0xbd, 0x71,
	0x09, 0x40, 0x8d, 0x62, 0x84, 0x5c, 0x58, 0x15, 0x79, 0xcc, 0xca, 0x35, 0x60, 0x53, 0xc5, 0x30,
	0x57, 0xff, 0x53, 0x36, 0xb3, 0xde, 0x2b, 0x58, 0x1b, 0x52, 0xd3, 0x77, 0x09, 0xcd, 0xf3, 0x65,
	0xb4, 0xff, 0x25, 0xcd, 0x95, 0x5e, 0x21, 0x8b, 0xfe, 0x44, 0xb1, 0x7f, 0x0f, 0xea, 0xca, 0x23,
	0x4a, 0xc4, 0x62, 0x00, 0x7f, 0x6f, 0xae, 0x2c, 0xdc, 0x2b, 0x58, 0xdb, 0x52, 0x18, 0x24, 0x22,
	0x27, 0xbd, 0xf1, 0x16, 0xd4, 0xae, 0x8e, 0x84, 0x6c, 0x43, 0x71, 0xb1, 0x0c, 0x45, 0xe6, 0x1a,
	0x3b, 0x50, 0x59, 0x50, 0x4d, 0xca, 0xb0, 0x36, 0x78, 0x68, 0xdf, 0xad, 0x17, 0xf2, 0x53, 0xbb,
	0xae, 0x1d, 0x6c, 0x41, 0x95, 0x9f, 0xb0, 0x99, 0x1d, 0x85, 0x72, 0xe7, 0x8d, 0x3e, 0xdc, 0xec,
	0x9e, 0x09, 0x8c, 0x43, 0xc7, 0x1f, 0xa5, 0x33, 0x94, 0x55, 0xe7, 0xa9, 0x4d, 0x58, 0x53, 0x9c,
	0x6a, 0xcf, 0xe4, 0x54, 0xf9, 0x19, 0xbb, 0x50, 0x39, 0x8c, 0x92, 0x50, 0x58, 0xd1, 0xd7, 0x5c,
	0x6e, 0xdf, 0x44, 0x0a, 0x2a, 0xba, 0x64, 0x65, 0xc2, 0xed, 0xb7, 0xa1, 0x2c, 0x2b, 0xa8, 0xe6,
	0x6a, 0x50, 0x1e, 0x3c, 0x1c, 0xf5, 0x07, 0x0f, 0x54, 0x83, 0x97, 0x52, 0xbb, 0xae, 0x75, 0xfe,
	0x29, 0xc3, 0x4b, 0xf9, 0xfe, 0x0c, 0x31, 0x3e, 0x65, 0x13, 0x6c, 0x93, 0xdf, 0x34, 0xd8, 0xfa,
	0x38, 0x74, 0xe2, 0x34, 0x37, 0xb4, 0xc9, 0xcd, 0x15, 0xcb, 0xd6, 0x58, 0x65, 0x30, 0x7e, 0xd0,
	0xbe, 0xbd, 0x38, 0x2f, 0x7e, 0xa7, 0x41, 0x30, 0xec, 0x7e, 0xd4, 0x3d, 0x1c, 0x51, 0xe6, 0xd2,
	0xfd, 0x21, 0x35, 0xe6, 0xb7, 0xca, 0xd8, 0xa3, 0x27, 0x98, 0xee, 0x51, 0x35, 0xc1, 0x3d, 0x1a,
	0x70, 0x8f, 0x3a, 0x9c, 0x3e, 0x76, 0x61, 0xf6, 0x28, 0x17, 0x8e, 0x48, 0xf8, 0xa5, 0x41, 0xae,
	0x3f, 0xbd, 0x6f, 0x0d, 0x8e, 0xa8, 0x9a, 0xad, 0xca, 0x45, 0x3f, 0xed, 0x75, 0xad, 0xae, 0x4c,
	0xff, 0x21, 0xbd, 0xd5, 0x26, 0x8b, 0x6b, 0x4b, 0xae, 0x2c, 0x2b, 0xf9, 0xf9, 0x09, 0x44, 0x1d,
	0xb2, 0x6c, 0x09, 0x56, 0xa3, 0x39, 0x93, 0x60, 0xf8, 0x8b, 0xc6, 0x52, 0x64, 0x2e, 0xf9, 0x45,
	0x03, 0x22, 0x07, 0x83, 0xf1, 0x50, 0xc4, 0xe8, 0x04, 0x43, 0xf4, 0x71, 0x22, 0xfe, 0xc7, 0x40,
	0xbe, 0x91, 0x10, 0xd2, 0x6b, 0x1b, 0xc7, 0x5d, 0x8d, 0xfc, 0x54, 0x82, 0xd7, 0x0e, 0x7d, 0x86,
	0xa1, 0xc8, 0xa0, 0xb0, 0xd0, 0xcb, 0x5b, 0x5c, 0x0d, 0xe7, 0xca, 0x98, 0x16, 0x1b, 0x6f, 0xfc,
	0x5e, 0x94, 0x58, 0x7e, 0x2d, 0x5e, 0x1b, 0x18, 0xeb, 0x18, 0xde, 0x5c, 0x8e, 0xe4, 0x00, 0xa7,
	0x51, 0x8c, 0xa4, 0xe7, 0x31, 0x71, 0x9c, 0x8c, 0xcd, 0x49, 0x14, 0xb4, 0xc4, 0x24, 0x9c, 0xf8,
	0x51, 0xe2, 0x66, 0x2f, 0xd3, 0xe4, 0x8e, 0x87, 0xe1, 0x9d, 0xf9, 0x1b, 0xf7, 0xcc, 0x37, 0xf3,
	0x91, 0x07, 0x6f, 0x2c, 0xaf, 0xb4, 0x3f, 0x15, 0x18, 0x3f, 0xbf, 0x42, 0x4d, 0xad, 0xf1, 0xe5,
	0xf7, 0x17, 0xe7, 0xc5, 0x2f, 0xe0, 0x73, 0x78, 0xca, 0x0f, 0xa8, 0x5e, 0x35, 0x36, 0x8e, 0x52,
	0x29, 0xde, 0x7e, 0xf7, 0xbf, 0x96, 0x0e, 0x52, 0x79, 0x17, 0x7f, 0xbc, 0x38, 0x2f, 0x16, 0x3a,
	0x37, 0xe0, 0xe5, 0x07, 0x91, 0xe8, 0x86, 0x92, 0x50, 0x37, 0xff, 0xf1, 0x1c, 0xd8, 0x8f, 0x3e,
	0x7b, 0x5e, 0x70, 0xde, 0x5f, 0xa6, 0x1c, 0x6f, 0xa8, 0x54, 0xef, 0xfc, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0xdf, 0xab, 0xd5, 0xac, 0x08, 0x00, 0x00,
}
