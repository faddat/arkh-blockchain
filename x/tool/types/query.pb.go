// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tool/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("tool/query.proto", fileDescriptor_de03359b5e5d5ab2) }

var fileDescriptor_de03359b5e5d5ab2 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0x31, 0x0e, 0x82, 0x30,
	0x18, 0x05, 0x60, 0x18, 0xd4, 0x84, 0xc9, 0xe8, 0x46, 0x4c, 0x47, 0x07, 0x87, 0xfe, 0x41, 0x2f,
	0x60, 0xbc, 0x81, 0xab, 0xdb, 0x5f, 0x6c, 0x4a, 0x23, 0xf4, 0xaf, 0xb4, 0x18, 0xb9, 0x85, 0xc7,
	0x72, 0x64, 0x74, 0x34, 0x70, 0x11, 0x43, 0x49, 0xdc, 0xde, 0xf0, 0xbd, 0x97, 0x97, 0x2c, 0x3d,
	0x51, 0x09, 0xf7, 0x46, 0xd6, 0x2d, 0xb7, 0x35, 0x79, 0x5a, 0xad, 0xb1, 0xbe, 0x15, 0x78, 0xd5,
	0x68, 0xf8, 0x98, 0xf8, 0x08, 0xd2, 0x8d, 0x22, 0x52, 0xa5, 0x04, 0xb4, 0x1a, 0xd0, 0x18, 0xf2,
	0xe8, 0x35, 0x19, 0x37, 0x55, 0xd2, 0x5d, 0x4e, 0xae, 0x22, 0x07, 0x02, 0x9d, 0x9c, 0xb6, 0xe0,
	0x91, 0x09, 0xe9, 0x31, 0x03, 0x8b, 0x4a, 0x9b, 0x80, 0x27, 0xbb, 0x5f, 0x24, 0xb3, 0xf3, 0x28,
	0x4e, 0xc7, 0x77, 0xcf, 0xe2, 0xae, 0x67, 0xf1, 0xb7, 0x67, 0xf1, 0x6b, 0x60, 0x51, 0x37, 0xb0,
	0xe8, 0x33, 0xb0, 0xe8, 0xb2, 0x55, 0xda, 0x17, 0x8d, 0xe0, 0x39, 0x55, 0xf0, 0x3f, 0x13, 0x12,
	0x3c, 0x21, 0xfc, 0xf5, 0xad, 0x95, 0x4e, 0xcc, 0xc3, 0xe2, 0xe1, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x98, 0xcc, 0x23, 0xc4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arkhadian.arkh.tool.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "tool/query.proto",
}
