//
//Copyright (c) 2020 Docker Inc.
//
//Permission is hereby granted, free of charge, to any person
//obtaining a copy of this software and associated documentation
//files (the "Software"), to deal in the Software without
//restriction, including without limitation the rights to use, copy,
//modify, merge, publish, distribute, sublicense, and/or sell copies
//of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be
//included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//EXPRESS OR IMPLIED,
//INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
//IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//HOLDERS BE LIABLE FOR ANY CLAIM,
//DAMAGES OR OTHER LIABILITY,
//WHETHER IN AN ACTION OF CONTRACT,
//TORT OR OTHERWISE,
//ARISING FROM, OUT OF OR IN CONNECTION WITH
//THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.6.1
// source: cli/v1/cli.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ContextType string `protobuf:"bytes,2,opt,name=contextType,proto3" json:"contextType,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_v1_cli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_cli_v1_cli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_cli_v1_cli_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Context) GetContextType() string {
	if x != nil {
		return x.ContextType
	}
	return ""
}

type ContextsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContextsRequest) Reset() {
	*x = ContextsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_v1_cli_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextsRequest) ProtoMessage() {}

func (x *ContextsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cli_v1_cli_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextsRequest.ProtoReflect.Descriptor instead.
func (*ContextsRequest) Descriptor() ([]byte, []int) {
	return file_cli_v1_cli_proto_rawDescGZIP(), []int{1}
}

type ContextsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contexts []*Context `protobuf:"bytes,1,rep,name=contexts,proto3" json:"contexts,omitempty"`
}

func (x *ContextsResponse) Reset() {
	*x = ContextsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cli_v1_cli_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextsResponse) ProtoMessage() {}

func (x *ContextsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cli_v1_cli_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextsResponse.ProtoReflect.Descriptor instead.
func (*ContextsResponse) Descriptor() ([]byte, []int) {
	return file_cli_v1_cli_proto_rawDescGZIP(), []int{2}
}

func (x *ContextsResponse) GetContexts() []*Context {
	if x != nil {
		return x.Contexts
	}
	return nil
}

var File_cli_v1_cli_proto protoreflect.FileDescriptor

var file_cli_v1_cli_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x3f, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x32, 0x62, 0x0a,
	0x03, 0x43, 0x6c, 0x69, 0x12, 0x5b, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73,
	0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cli_v1_cli_proto_rawDescOnce sync.Once
	file_cli_v1_cli_proto_rawDescData = file_cli_v1_cli_proto_rawDesc
)

func file_cli_v1_cli_proto_rawDescGZIP() []byte {
	file_cli_v1_cli_proto_rawDescOnce.Do(func() {
		file_cli_v1_cli_proto_rawDescData = protoimpl.X.CompressGZIP(file_cli_v1_cli_proto_rawDescData)
	})
	return file_cli_v1_cli_proto_rawDescData
}

var file_cli_v1_cli_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cli_v1_cli_proto_goTypes = []interface{}{
	(*Context)(nil),          // 0: com.docker.api.cli.v1.Context
	(*ContextsRequest)(nil),  // 1: com.docker.api.cli.v1.ContextsRequest
	(*ContextsResponse)(nil), // 2: com.docker.api.cli.v1.ContextsResponse
}
var file_cli_v1_cli_proto_depIdxs = []int32{
	0, // 0: com.docker.api.cli.v1.ContextsResponse.contexts:type_name -> com.docker.api.cli.v1.Context
	1, // 1: com.docker.api.cli.v1.Cli.Contexts:input_type -> com.docker.api.cli.v1.ContextsRequest
	2, // 2: com.docker.api.cli.v1.Cli.Contexts:output_type -> com.docker.api.cli.v1.ContextsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cli_v1_cli_proto_init() }
func file_cli_v1_cli_proto_init() {
	if File_cli_v1_cli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cli_v1_cli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cli_v1_cli_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cli_v1_cli_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cli_v1_cli_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cli_v1_cli_proto_goTypes,
		DependencyIndexes: file_cli_v1_cli_proto_depIdxs,
		MessageInfos:      file_cli_v1_cli_proto_msgTypes,
	}.Build()
	File_cli_v1_cli_proto = out.File
	file_cli_v1_cli_proto_rawDesc = nil
	file_cli_v1_cli_proto_goTypes = nil
	file_cli_v1_cli_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CliClient is the client API for Cli service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CliClient interface {
	// Returns the list of existing contexts
	Contexts(ctx context.Context, in *ContextsRequest, opts ...grpc.CallOption) (*ContextsResponse, error)
}

type cliClient struct {
	cc grpc.ClientConnInterface
}

func NewCliClient(cc grpc.ClientConnInterface) CliClient {
	return &cliClient{cc}
}

func (c *cliClient) Contexts(ctx context.Context, in *ContextsRequest, opts ...grpc.CallOption) (*ContextsResponse, error) {
	out := new(ContextsResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.cli.v1.Cli/Contexts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CliServer is the server API for Cli service.
type CliServer interface {
	// Returns the list of existing contexts
	Contexts(context.Context, *ContextsRequest) (*ContextsResponse, error)
}

// UnimplementedCliServer can be embedded to have forward compatible implementations.
type UnimplementedCliServer struct {
}

func (*UnimplementedCliServer) Contexts(context.Context, *ContextsRequest) (*ContextsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Contexts not implemented")
}

func RegisterCliServer(s *grpc.Server, srv CliServer) {
	s.RegisterService(&_Cli_serviceDesc, srv)
}

func _Cli_Contexts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).Contexts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.cli.v1.Cli/Contexts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).Contexts(ctx, req.(*ContextsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cli_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.docker.api.cli.v1.Cli",
	HandlerType: (*CliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Contexts",
			Handler:    _Cli_Contexts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli/v1/cli.proto",
}
