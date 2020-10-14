// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: hobbyfarm/environment_service.proto

package protobuf

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type EnvironmentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environments []*Environment `protobuf:"bytes,1,rep,name=Environments,proto3" json:"Environments,omitempty"`
}

func (x *EnvironmentList) Reset() {
	*x = EnvironmentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hobbyfarm_environment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentList) ProtoMessage() {}

func (x *EnvironmentList) ProtoReflect() protoreflect.Message {
	mi := &file_hobbyfarm_environment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentList.ProtoReflect.Descriptor instead.
func (*EnvironmentList) Descriptor() ([]byte, []int) {
	return file_hobbyfarm_environment_service_proto_rawDescGZIP(), []int{0}
}

func (x *EnvironmentList) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

var File_hobbyfarm_environment_service_proto protoreflect.FileDescriptor

var file_hobbyfarm_environment_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x68,
	0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43,
	0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xe0, 0x02, 0x0a, 0x12, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x0c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x0c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4c,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x1a, 0x1b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x7b, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x38, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67,
	0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hobbyfarm_environment_service_proto_rawDescOnce sync.Once
	file_hobbyfarm_environment_service_proto_rawDescData = file_hobbyfarm_environment_service_proto_rawDesc
)

func file_hobbyfarm_environment_service_proto_rawDescGZIP() []byte {
	file_hobbyfarm_environment_service_proto_rawDescOnce.Do(func() {
		file_hobbyfarm_environment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_hobbyfarm_environment_service_proto_rawDescData)
	})
	return file_hobbyfarm_environment_service_proto_rawDescData
}

var file_hobbyfarm_environment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hobbyfarm_environment_service_proto_goTypes = []interface{}{
	(*EnvironmentList)(nil), // 0: EnvironmentList
	(*Environment)(nil),     // 1: Environment
	(*Empty)(nil),           // 2: Empty
	(*ID)(nil),              // 3: ID
}
var file_hobbyfarm_environment_service_proto_depIdxs = []int32{
	1, // 0: EnvironmentList.Environments:type_name -> Environment
	2, // 1: EnvironmentService.List:input_type -> Empty
	3, // 2: EnvironmentService.Get:input_type -> ID
	1, // 3: EnvironmentService.Create:input_type -> Environment
	1, // 4: EnvironmentService.Update:input_type -> Environment
	3, // 5: EnvironmentService.Delete:input_type -> ID
	0, // 6: EnvironmentService.List:output_type -> EnvironmentList
	1, // 7: EnvironmentService.Get:output_type -> Environment
	1, // 8: EnvironmentService.Create:output_type -> Environment
	1, // 9: EnvironmentService.Update:output_type -> Environment
	2, // 10: EnvironmentService.Delete:output_type -> Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hobbyfarm_environment_service_proto_init() }
func file_hobbyfarm_environment_service_proto_init() {
	if File_hobbyfarm_environment_service_proto != nil {
		return
	}
	file_hobbyfarm_hobbyfarm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hobbyfarm_environment_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentList); i {
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
			RawDescriptor: file_hobbyfarm_environment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hobbyfarm_environment_service_proto_goTypes,
		DependencyIndexes: file_hobbyfarm_environment_service_proto_depIdxs,
		MessageInfos:      file_hobbyfarm_environment_service_proto_msgTypes,
	}.Build()
	File_hobbyfarm_environment_service_proto = out.File
	file_hobbyfarm_environment_service_proto_rawDesc = nil
	file_hobbyfarm_environment_service_proto_goTypes = nil
	file_hobbyfarm_environment_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EnvironmentServiceClient is the client API for EnvironmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnvironmentServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EnvironmentList, error)
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Environment, error)
	Create(ctx context.Context, in *Environment, opts ...grpc.CallOption) (*Environment, error)
	Update(ctx context.Context, in *Environment, opts ...grpc.CallOption) (*Environment, error)
	Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type environmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentServiceClient(cc grpc.ClientConnInterface) EnvironmentServiceClient {
	return &environmentServiceClient{cc}
}

func (c *environmentServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EnvironmentList, error) {
	out := new(EnvironmentList)
	err := c.cc.Invoke(ctx, "/EnvironmentService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := c.cc.Invoke(ctx, "/EnvironmentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) Create(ctx context.Context, in *Environment, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := c.cc.Invoke(ctx, "/EnvironmentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) Update(ctx context.Context, in *Environment, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := c.cc.Invoke(ctx, "/EnvironmentService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServiceClient) Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/EnvironmentService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentServiceServer is the server API for EnvironmentService service.
type EnvironmentServiceServer interface {
	List(context.Context, *Empty) (*EnvironmentList, error)
	Get(context.Context, *ID) (*Environment, error)
	Create(context.Context, *Environment) (*Environment, error)
	Update(context.Context, *Environment) (*Environment, error)
	Delete(context.Context, *ID) (*Empty, error)
}

// UnimplementedEnvironmentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEnvironmentServiceServer struct {
}

func (*UnimplementedEnvironmentServiceServer) List(context.Context, *Empty) (*EnvironmentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEnvironmentServiceServer) Get(context.Context, *ID) (*Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEnvironmentServiceServer) Create(context.Context, *Environment) (*Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEnvironmentServiceServer) Update(context.Context, *Environment) (*Environment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedEnvironmentServiceServer) Delete(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEnvironmentServiceServer(s *grpc.Server, srv EnvironmentServiceServer) {
	s.RegisterService(&_EnvironmentService_serviceDesc, srv)
}

func _EnvironmentService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EnvironmentService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EnvironmentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Environment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EnvironmentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).Create(ctx, req.(*Environment))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Environment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EnvironmentService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).Update(ctx, req.(*Environment))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EnvironmentService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentServiceServer).Delete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnvironmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EnvironmentService",
	HandlerType: (*EnvironmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _EnvironmentService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EnvironmentService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EnvironmentService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EnvironmentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EnvironmentService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hobbyfarm/environment_service.proto",
}
