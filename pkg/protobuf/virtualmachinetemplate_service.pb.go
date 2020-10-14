// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: hobbyfarm/virtualmachinetemplate_service.proto

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

type VirtualMachineTemplateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtualMachineTemplates []*VirtualMachineTemplate `protobuf:"bytes,1,rep,name=VirtualMachineTemplates,proto3" json:"VirtualMachineTemplates,omitempty"`
}

func (x *VirtualMachineTemplateList) Reset() {
	*x = VirtualMachineTemplateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hobbyfarm_virtualmachinetemplate_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineTemplateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineTemplateList) ProtoMessage() {}

func (x *VirtualMachineTemplateList) ProtoReflect() protoreflect.Message {
	mi := &file_hobbyfarm_virtualmachinetemplate_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineTemplateList.ProtoReflect.Descriptor instead.
func (*VirtualMachineTemplateList) Descriptor() ([]byte, []int) {
	return file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualMachineTemplateList) GetVirtualMachineTemplates() []*VirtualMachineTemplate {
	if x != nil {
		return x.VirtualMachineTemplates
	}
	return nil
}

var File_hobbyfarm_virtualmachinetemplate_service_proto protoreflect.FileDescriptor

var file_hobbyfarm_virtualmachinetemplate_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x66,
	0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x1a, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x17, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x17, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x32, 0xe4, 0x03, 0x0a, 0x1d, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x51, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x17,
	0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12,
	0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x12, 0x66, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x17, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x1a, 0x17, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2b, 0x1a, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x2f, 0x7b, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x43, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x2a, 0x24, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x49, 0x44,
	0x7d, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x67, 0x61, 0x6e,
	0x74, 0x75, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescOnce sync.Once
	file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescData = file_hobbyfarm_virtualmachinetemplate_service_proto_rawDesc
)

func file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescGZIP() []byte {
	file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescOnce.Do(func() {
		file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescData)
	})
	return file_hobbyfarm_virtualmachinetemplate_service_proto_rawDescData
}

var file_hobbyfarm_virtualmachinetemplate_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hobbyfarm_virtualmachinetemplate_service_proto_goTypes = []interface{}{
	(*VirtualMachineTemplateList)(nil), // 0: VirtualMachineTemplateList
	(*VirtualMachineTemplate)(nil),     // 1: VirtualMachineTemplate
	(*Empty)(nil),                      // 2: Empty
	(*ID)(nil),                         // 3: ID
}
var file_hobbyfarm_virtualmachinetemplate_service_proto_depIdxs = []int32{
	1, // 0: VirtualMachineTemplateList.VirtualMachineTemplates:type_name -> VirtualMachineTemplate
	2, // 1: VirtualMachineTemplateService.List:input_type -> Empty
	3, // 2: VirtualMachineTemplateService.Get:input_type -> ID
	1, // 3: VirtualMachineTemplateService.Create:input_type -> VirtualMachineTemplate
	1, // 4: VirtualMachineTemplateService.Update:input_type -> VirtualMachineTemplate
	3, // 5: VirtualMachineTemplateService.Delete:input_type -> ID
	0, // 6: VirtualMachineTemplateService.List:output_type -> VirtualMachineTemplateList
	1, // 7: VirtualMachineTemplateService.Get:output_type -> VirtualMachineTemplate
	1, // 8: VirtualMachineTemplateService.Create:output_type -> VirtualMachineTemplate
	1, // 9: VirtualMachineTemplateService.Update:output_type -> VirtualMachineTemplate
	2, // 10: VirtualMachineTemplateService.Delete:output_type -> Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hobbyfarm_virtualmachinetemplate_service_proto_init() }
func file_hobbyfarm_virtualmachinetemplate_service_proto_init() {
	if File_hobbyfarm_virtualmachinetemplate_service_proto != nil {
		return
	}
	file_hobbyfarm_hobbyfarm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hobbyfarm_virtualmachinetemplate_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineTemplateList); i {
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
			RawDescriptor: file_hobbyfarm_virtualmachinetemplate_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hobbyfarm_virtualmachinetemplate_service_proto_goTypes,
		DependencyIndexes: file_hobbyfarm_virtualmachinetemplate_service_proto_depIdxs,
		MessageInfos:      file_hobbyfarm_virtualmachinetemplate_service_proto_msgTypes,
	}.Build()
	File_hobbyfarm_virtualmachinetemplate_service_proto = out.File
	file_hobbyfarm_virtualmachinetemplate_service_proto_rawDesc = nil
	file_hobbyfarm_virtualmachinetemplate_service_proto_goTypes = nil
	file_hobbyfarm_virtualmachinetemplate_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VirtualMachineTemplateServiceClient is the client API for VirtualMachineTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineTemplateServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VirtualMachineTemplateList, error)
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*VirtualMachineTemplate, error)
	Create(ctx context.Context, in *VirtualMachineTemplate, opts ...grpc.CallOption) (*VirtualMachineTemplate, error)
	Update(ctx context.Context, in *VirtualMachineTemplate, opts ...grpc.CallOption) (*VirtualMachineTemplate, error)
	Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type virtualMachineTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVirtualMachineTemplateServiceClient(cc grpc.ClientConnInterface) VirtualMachineTemplateServiceClient {
	return &virtualMachineTemplateServiceClient{cc}
}

func (c *virtualMachineTemplateServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VirtualMachineTemplateList, error) {
	out := new(VirtualMachineTemplateList)
	err := c.cc.Invoke(ctx, "/VirtualMachineTemplateService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineTemplateServiceClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*VirtualMachineTemplate, error) {
	out := new(VirtualMachineTemplate)
	err := c.cc.Invoke(ctx, "/VirtualMachineTemplateService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineTemplateServiceClient) Create(ctx context.Context, in *VirtualMachineTemplate, opts ...grpc.CallOption) (*VirtualMachineTemplate, error) {
	out := new(VirtualMachineTemplate)
	err := c.cc.Invoke(ctx, "/VirtualMachineTemplateService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineTemplateServiceClient) Update(ctx context.Context, in *VirtualMachineTemplate, opts ...grpc.CallOption) (*VirtualMachineTemplate, error) {
	out := new(VirtualMachineTemplate)
	err := c.cc.Invoke(ctx, "/VirtualMachineTemplateService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineTemplateServiceClient) Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/VirtualMachineTemplateService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineTemplateServiceServer is the server API for VirtualMachineTemplateService service.
type VirtualMachineTemplateServiceServer interface {
	List(context.Context, *Empty) (*VirtualMachineTemplateList, error)
	Get(context.Context, *ID) (*VirtualMachineTemplate, error)
	Create(context.Context, *VirtualMachineTemplate) (*VirtualMachineTemplate, error)
	Update(context.Context, *VirtualMachineTemplate) (*VirtualMachineTemplate, error)
	Delete(context.Context, *ID) (*Empty, error)
}

// UnimplementedVirtualMachineTemplateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineTemplateServiceServer struct {
}

func (*UnimplementedVirtualMachineTemplateServiceServer) List(context.Context, *Empty) (*VirtualMachineTemplateList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedVirtualMachineTemplateServiceServer) Get(context.Context, *ID) (*VirtualMachineTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedVirtualMachineTemplateServiceServer) Create(context.Context, *VirtualMachineTemplate) (*VirtualMachineTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedVirtualMachineTemplateServiceServer) Update(context.Context, *VirtualMachineTemplate) (*VirtualMachineTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedVirtualMachineTemplateServiceServer) Delete(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterVirtualMachineTemplateServiceServer(s *grpc.Server, srv VirtualMachineTemplateServiceServer) {
	s.RegisterService(&_VirtualMachineTemplateService_serviceDesc, srv)
}

func _VirtualMachineTemplateService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineTemplateServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineTemplateService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineTemplateServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineTemplateService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineTemplateServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineTemplateService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineTemplateServiceServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineTemplateService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineTemplateServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineTemplateService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineTemplateServiceServer).Create(ctx, req.(*VirtualMachineTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineTemplateService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineTemplateServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineTemplateService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineTemplateServiceServer).Update(ctx, req.(*VirtualMachineTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineTemplateService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineTemplateServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VirtualMachineTemplateService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineTemplateServiceServer).Delete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineTemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "VirtualMachineTemplateService",
	HandlerType: (*VirtualMachineTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _VirtualMachineTemplateService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _VirtualMachineTemplateService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _VirtualMachineTemplateService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VirtualMachineTemplateService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VirtualMachineTemplateService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hobbyfarm/virtualmachinetemplate_service.proto",
}
