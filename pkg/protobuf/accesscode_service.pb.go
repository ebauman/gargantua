// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: hobbyfarm/accesscode_service.proto

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

type AccessCodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessCodes []*AccessCode `protobuf:"bytes,1,rep,name=AccessCodes,proto3" json:"AccessCodes,omitempty"`
}

func (x *AccessCodeList) Reset() {
	*x = AccessCodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hobbyfarm_accesscode_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessCodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessCodeList) ProtoMessage() {}

func (x *AccessCodeList) ProtoReflect() protoreflect.Message {
	mi := &file_hobbyfarm_accesscode_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessCodeList.ProtoReflect.Descriptor instead.
func (*AccessCodeList) Descriptor() ([]byte, []int) {
	return file_hobbyfarm_accesscode_service_proto_rawDescGZIP(), []int{0}
}

func (x *AccessCodeList) GetAccessCodes() []*AccessCode {
	if x != nil {
		return x.AccessCodes
	}
	return nil
}

var File_hobbyfarm_accesscode_service_proto protoreflect.FileDescriptor

var file_hobbyfarm_accesscode_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x68, 0x6f,
	0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a,
	0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x32, 0xd8,
	0x02, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0xc0, 0xb8, 0x02, 0x01, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x03, 0x2e,
	0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x7b, 0x49, 0x44,
	0x7d, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x0b, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x49, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x0b, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x1a, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x7b, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x01, 0x2a,
	0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a,
	0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72,
	0x6d, 0x2f, 0x67, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hobbyfarm_accesscode_service_proto_rawDescOnce sync.Once
	file_hobbyfarm_accesscode_service_proto_rawDescData = file_hobbyfarm_accesscode_service_proto_rawDesc
)

func file_hobbyfarm_accesscode_service_proto_rawDescGZIP() []byte {
	file_hobbyfarm_accesscode_service_proto_rawDescOnce.Do(func() {
		file_hobbyfarm_accesscode_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_hobbyfarm_accesscode_service_proto_rawDescData)
	})
	return file_hobbyfarm_accesscode_service_proto_rawDescData
}

var file_hobbyfarm_accesscode_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hobbyfarm_accesscode_service_proto_goTypes = []interface{}{
	(*AccessCodeList)(nil), // 0: AccessCodeList
	(*AccessCode)(nil),     // 1: AccessCode
	(*Empty)(nil),          // 2: Empty
	(*ID)(nil),             // 3: ID
}
var file_hobbyfarm_accesscode_service_proto_depIdxs = []int32{
	1, // 0: AccessCodeList.AccessCodes:type_name -> AccessCode
	2, // 1: AccessCodeService.List:input_type -> Empty
	3, // 2: AccessCodeService.Get:input_type -> ID
	1, // 3: AccessCodeService.Create:input_type -> AccessCode
	1, // 4: AccessCodeService.Update:input_type -> AccessCode
	3, // 5: AccessCodeService.Delete:input_type -> ID
	0, // 6: AccessCodeService.List:output_type -> AccessCodeList
	1, // 7: AccessCodeService.Get:output_type -> AccessCode
	1, // 8: AccessCodeService.Create:output_type -> AccessCode
	1, // 9: AccessCodeService.Update:output_type -> AccessCode
	2, // 10: AccessCodeService.Delete:output_type -> Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hobbyfarm_accesscode_service_proto_init() }
func file_hobbyfarm_accesscode_service_proto_init() {
	if File_hobbyfarm_accesscode_service_proto != nil {
		return
	}
	file_hobbyfarm_hobbyfarm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hobbyfarm_accesscode_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessCodeList); i {
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
			RawDescriptor: file_hobbyfarm_accesscode_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hobbyfarm_accesscode_service_proto_goTypes,
		DependencyIndexes: file_hobbyfarm_accesscode_service_proto_depIdxs,
		MessageInfos:      file_hobbyfarm_accesscode_service_proto_msgTypes,
	}.Build()
	File_hobbyfarm_accesscode_service_proto = out.File
	file_hobbyfarm_accesscode_service_proto_rawDesc = nil
	file_hobbyfarm_accesscode_service_proto_goTypes = nil
	file_hobbyfarm_accesscode_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccessCodeServiceClient is the client API for AccessCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessCodeServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccessCodeList, error)
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AccessCode, error)
	Create(ctx context.Context, in *AccessCode, opts ...grpc.CallOption) (*AccessCode, error)
	Update(ctx context.Context, in *AccessCode, opts ...grpc.CallOption) (*AccessCode, error)
	Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type accessCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessCodeServiceClient(cc grpc.ClientConnInterface) AccessCodeServiceClient {
	return &accessCodeServiceClient{cc}
}

func (c *accessCodeServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccessCodeList, error) {
	out := new(AccessCodeList)
	err := c.cc.Invoke(ctx, "/AccessCodeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeServiceClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AccessCode, error) {
	out := new(AccessCode)
	err := c.cc.Invoke(ctx, "/AccessCodeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeServiceClient) Create(ctx context.Context, in *AccessCode, opts ...grpc.CallOption) (*AccessCode, error) {
	out := new(AccessCode)
	err := c.cc.Invoke(ctx, "/AccessCodeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeServiceClient) Update(ctx context.Context, in *AccessCode, opts ...grpc.CallOption) (*AccessCode, error) {
	out := new(AccessCode)
	err := c.cc.Invoke(ctx, "/AccessCodeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessCodeServiceClient) Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AccessCodeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessCodeServiceServer is the server API for AccessCodeService service.
type AccessCodeServiceServer interface {
	List(context.Context, *Empty) (*AccessCodeList, error)
	Get(context.Context, *ID) (*AccessCode, error)
	Create(context.Context, *AccessCode) (*AccessCode, error)
	Update(context.Context, *AccessCode) (*AccessCode, error)
	Delete(context.Context, *ID) (*Empty, error)
}

// UnimplementedAccessCodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccessCodeServiceServer struct {
}

func (*UnimplementedAccessCodeServiceServer) List(context.Context, *Empty) (*AccessCodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAccessCodeServiceServer) Get(context.Context, *ID) (*AccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccessCodeServiceServer) Create(context.Context, *AccessCode) (*AccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccessCodeServiceServer) Update(context.Context, *AccessCode) (*AccessCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAccessCodeServiceServer) Delete(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAccessCodeServiceServer(s *grpc.Server, srv AccessCodeServiceServer) {
	s.RegisterService(&_AccessCodeService_serviceDesc, srv)
}

func _AccessCodeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccessCodeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccessCodeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeServiceServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccessCodeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeServiceServer).Create(ctx, req.(*AccessCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccessCodeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeServiceServer).Update(ctx, req.(*AccessCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessCodeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCodeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccessCodeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCodeServiceServer).Delete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessCodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AccessCodeService",
	HandlerType: (*AccessCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AccessCodeService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccessCodeService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccessCodeService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccessCodeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccessCodeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hobbyfarm/accesscode_service.proto",
}
