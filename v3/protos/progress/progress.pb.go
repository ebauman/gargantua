// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: progress/progress.proto

package progresspb

import (
	general "github.com/hobbyfarm/gargantua/v3/protos/general"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentStep uint32            `protobuf:"varint,1,opt,name=current_step,json=currentStep,proto3" json:"current_step,omitempty"`
	MaxStep     uint32            `protobuf:"varint,2,opt,name=max_step,json=maxStep,proto3" json:"max_step,omitempty"`
	TotalStep   uint32            `protobuf:"varint,3,opt,name=total_step,json=totalStep,proto3" json:"total_step,omitempty"`
	Scenario    string            `protobuf:"bytes,4,opt,name=scenario,proto3" json:"scenario,omitempty"`
	Course      string            `protobuf:"bytes,5,opt,name=course,proto3" json:"course,omitempty"`
	User        string            `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Labels      map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateProgressRequest) Reset() {
	*x = CreateProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_progress_progress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProgressRequest) ProtoMessage() {}

func (x *CreateProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_progress_progress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProgressRequest.ProtoReflect.Descriptor instead.
func (*CreateProgressRequest) Descriptor() ([]byte, []int) {
	return file_progress_progress_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProgressRequest) GetCurrentStep() uint32 {
	if x != nil {
		return x.CurrentStep
	}
	return 0
}

func (x *CreateProgressRequest) GetMaxStep() uint32 {
	if x != nil {
		return x.MaxStep
	}
	return 0
}

func (x *CreateProgressRequest) GetTotalStep() uint32 {
	if x != nil {
		return x.TotalStep
	}
	return 0
}

func (x *CreateProgressRequest) GetScenario() string {
	if x != nil {
		return x.Scenario
	}
	return ""
}

func (x *CreateProgressRequest) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *CreateProgressRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateProgressRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Progress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid               string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	CurrentStep       uint32                 `protobuf:"varint,3,opt,name=current_step,json=currentStep,proto3" json:"current_step,omitempty"`
	MaxStep           uint32                 `protobuf:"varint,4,opt,name=max_step,json=maxStep,proto3" json:"max_step,omitempty"`
	TotalStep         uint32                 `protobuf:"varint,5,opt,name=total_step,json=totalStep,proto3" json:"total_step,omitempty"`
	Scenario          string                 `protobuf:"bytes,6,opt,name=scenario,proto3" json:"scenario,omitempty"`
	Course            string                 `protobuf:"bytes,7,opt,name=course,proto3" json:"course,omitempty"`
	User              string                 `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	Started           string                 `protobuf:"bytes,9,opt,name=started,proto3" json:"started,omitempty"`
	LastUpdate        string                 `protobuf:"bytes,10,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	Finished          string                 `protobuf:"bytes,11,opt,name=finished,proto3" json:"finished,omitempty"`
	Steps             []*ProgressStep        `protobuf:"bytes,12,rep,name=steps,proto3" json:"steps,omitempty"`
	Labels            map[string]string      `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreationTimestamp *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
}

func (x *Progress) Reset() {
	*x = Progress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_progress_progress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress) ProtoMessage() {}

func (x *Progress) ProtoReflect() protoreflect.Message {
	mi := &file_progress_progress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress.ProtoReflect.Descriptor instead.
func (*Progress) Descriptor() ([]byte, []int) {
	return file_progress_progress_proto_rawDescGZIP(), []int{1}
}

func (x *Progress) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Progress) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Progress) GetCurrentStep() uint32 {
	if x != nil {
		return x.CurrentStep
	}
	return 0
}

func (x *Progress) GetMaxStep() uint32 {
	if x != nil {
		return x.MaxStep
	}
	return 0
}

func (x *Progress) GetTotalStep() uint32 {
	if x != nil {
		return x.TotalStep
	}
	return 0
}

func (x *Progress) GetScenario() string {
	if x != nil {
		return x.Scenario
	}
	return ""
}

func (x *Progress) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Progress) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Progress) GetStarted() string {
	if x != nil {
		return x.Started
	}
	return ""
}

func (x *Progress) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

func (x *Progress) GetFinished() string {
	if x != nil {
		return x.Finished
	}
	return ""
}

func (x *Progress) GetSteps() []*ProgressStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Progress) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Progress) GetCreationTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTimestamp
	}
	return nil
}

type ProgressStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step      uint32 `protobuf:"varint,1,opt,name=step,proto3" json:"step,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ProgressStep) Reset() {
	*x = ProgressStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_progress_progress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressStep) ProtoMessage() {}

func (x *ProgressStep) ProtoReflect() protoreflect.Message {
	mi := &file_progress_progress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressStep.ProtoReflect.Descriptor instead.
func (*ProgressStep) Descriptor() ([]byte, []int) {
	return file_progress_progress_proto_rawDescGZIP(), []int{2}
}

func (x *ProgressStep) GetStep() uint32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *ProgressStep) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type UpdateProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CurrentStep *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=current_step,json=currentStep,proto3" json:"current_step,omitempty"`
	MaxStep     *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_step,json=maxStep,proto3" json:"max_step,omitempty"`
	TotalStep   *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=total_step,json=totalStep,proto3" json:"total_step,omitempty"`
	LastUpdate  string                  `protobuf:"bytes,5,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	Finished    string                  `protobuf:"bytes,6,opt,name=finished,proto3" json:"finished,omitempty"`
	Steps       []*ProgressStep         `protobuf:"bytes,7,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *UpdateProgressRequest) Reset() {
	*x = UpdateProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_progress_progress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProgressRequest) ProtoMessage() {}

func (x *UpdateProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_progress_progress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProgressRequest.ProtoReflect.Descriptor instead.
func (*UpdateProgressRequest) Descriptor() ([]byte, []int) {
	return file_progress_progress_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProgressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProgressRequest) GetCurrentStep() *wrapperspb.UInt32Value {
	if x != nil {
		return x.CurrentStep
	}
	return nil
}

func (x *UpdateProgressRequest) GetMaxStep() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxStep
	}
	return nil
}

func (x *UpdateProgressRequest) GetTotalStep() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TotalStep
	}
	return nil
}

func (x *UpdateProgressRequest) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

func (x *UpdateProgressRequest) GetFinished() string {
	if x != nil {
		return x.Finished
	}
	return ""
}

func (x *UpdateProgressRequest) GetSteps() []*ProgressStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

type UpdateCollectionProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labelselector string                  `protobuf:"bytes,1,opt,name=labelselector,proto3" json:"labelselector,omitempty"`
	CurrentStep   *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=current_step,json=currentStep,proto3" json:"current_step,omitempty"`
	MaxStep       *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=max_step,json=maxStep,proto3" json:"max_step,omitempty"`
	TotalStep     *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=total_step,json=totalStep,proto3" json:"total_step,omitempty"`
	LastUpdate    string                  `protobuf:"bytes,5,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	Finished      string                  `protobuf:"bytes,6,opt,name=finished,proto3" json:"finished,omitempty"`
	Steps         []*ProgressStep         `protobuf:"bytes,7,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *UpdateCollectionProgressRequest) Reset() {
	*x = UpdateCollectionProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_progress_progress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCollectionProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCollectionProgressRequest) ProtoMessage() {}

func (x *UpdateCollectionProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_progress_progress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCollectionProgressRequest.ProtoReflect.Descriptor instead.
func (*UpdateCollectionProgressRequest) Descriptor() ([]byte, []int) {
	return file_progress_progress_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCollectionProgressRequest) GetLabelselector() string {
	if x != nil {
		return x.Labelselector
	}
	return ""
}

func (x *UpdateCollectionProgressRequest) GetCurrentStep() *wrapperspb.UInt32Value {
	if x != nil {
		return x.CurrentStep
	}
	return nil
}

func (x *UpdateCollectionProgressRequest) GetMaxStep() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxStep
	}
	return nil
}

func (x *UpdateCollectionProgressRequest) GetTotalStep() *wrapperspb.UInt32Value {
	if x != nil {
		return x.TotalStep
	}
	return nil
}

func (x *UpdateCollectionProgressRequest) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

func (x *UpdateCollectionProgressRequest) GetFinished() string {
	if x != nil {
		return x.Finished
	}
	return ""
}

func (x *UpdateCollectionProgressRequest) GetSteps() []*ProgressStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

type ListProgressesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progresses []*Progress `protobuf:"bytes,1,rep,name=progresses,proto3" json:"progresses,omitempty"`
}

func (x *ListProgressesResponse) Reset() {
	*x = ListProgressesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_progress_progress_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProgressesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProgressesResponse) ProtoMessage() {}

func (x *ListProgressesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_progress_progress_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProgressesResponse.ProtoReflect.Descriptor instead.
func (*ListProgressesResponse) Descriptor() ([]byte, []int) {
	return file_progress_progress_proto_rawDescGZIP(), []int{5}
}

func (x *ListProgressesResponse) GetProgresses() []*Progress {
	if x != nil {
		return x.Progresses
	}
	return nil
}

var File_progress_progress_proto protoreflect.FileDescriptor

var file_progress_progress_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x65, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x04, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78,
	0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x78,
	0x53, 0x74, 0x65, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x65, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x74, 0x65, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x12, 0x36, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x49, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x40,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0xc9, 0x02, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x37, 0x0a, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x78,
	0x53, 0x74, 0x65, 0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x65,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x2c,
	0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0xe9, 0x02, 0x0a,
	0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x37, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73,
	0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x65, 0x70,
	0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65,
	0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0x4c, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x32, 0x88, 0x04, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x76, 0x63, 0x12, 0x46, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x36,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x13, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x5d, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x48, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x6f, 0x62, 0x62, 0x79, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x67, 0x61, 0x6e,
	0x74, 0x75, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x3b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_progress_progress_proto_rawDescOnce sync.Once
	file_progress_progress_proto_rawDescData = file_progress_progress_proto_rawDesc
)

func file_progress_progress_proto_rawDescGZIP() []byte {
	file_progress_progress_proto_rawDescOnce.Do(func() {
		file_progress_progress_proto_rawDescData = protoimpl.X.CompressGZIP(file_progress_progress_proto_rawDescData)
	})
	return file_progress_progress_proto_rawDescData
}

var file_progress_progress_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_progress_progress_proto_goTypes = []interface{}{
	(*CreateProgressRequest)(nil),           // 0: progress.CreateProgressRequest
	(*Progress)(nil),                        // 1: progress.Progress
	(*ProgressStep)(nil),                    // 2: progress.ProgressStep
	(*UpdateProgressRequest)(nil),           // 3: progress.UpdateProgressRequest
	(*UpdateCollectionProgressRequest)(nil), // 4: progress.UpdateCollectionProgressRequest
	(*ListProgressesResponse)(nil),          // 5: progress.ListProgressesResponse
	nil,                                     // 6: progress.CreateProgressRequest.LabelsEntry
	nil,                                     // 7: progress.Progress.LabelsEntry
	(*timestamppb.Timestamp)(nil),           // 8: google.protobuf.Timestamp
	(*wrapperspb.UInt32Value)(nil),          // 9: google.protobuf.UInt32Value
	(*general.GetRequest)(nil),              // 10: general.GetRequest
	(*general.ResourceId)(nil),              // 11: general.ResourceId
	(*general.ListOptions)(nil),             // 12: general.ListOptions
	(*emptypb.Empty)(nil),                   // 13: google.protobuf.Empty
}
var file_progress_progress_proto_depIdxs = []int32{
	6,  // 0: progress.CreateProgressRequest.labels:type_name -> progress.CreateProgressRequest.LabelsEntry
	2,  // 1: progress.Progress.steps:type_name -> progress.ProgressStep
	7,  // 2: progress.Progress.labels:type_name -> progress.Progress.LabelsEntry
	8,  // 3: progress.Progress.creation_timestamp:type_name -> google.protobuf.Timestamp
	9,  // 4: progress.UpdateProgressRequest.current_step:type_name -> google.protobuf.UInt32Value
	9,  // 5: progress.UpdateProgressRequest.max_step:type_name -> google.protobuf.UInt32Value
	9,  // 6: progress.UpdateProgressRequest.total_step:type_name -> google.protobuf.UInt32Value
	2,  // 7: progress.UpdateProgressRequest.steps:type_name -> progress.ProgressStep
	9,  // 8: progress.UpdateCollectionProgressRequest.current_step:type_name -> google.protobuf.UInt32Value
	9,  // 9: progress.UpdateCollectionProgressRequest.max_step:type_name -> google.protobuf.UInt32Value
	9,  // 10: progress.UpdateCollectionProgressRequest.total_step:type_name -> google.protobuf.UInt32Value
	2,  // 11: progress.UpdateCollectionProgressRequest.steps:type_name -> progress.ProgressStep
	1,  // 12: progress.ListProgressesResponse.progresses:type_name -> progress.Progress
	0,  // 13: progress.ProgressSvc.CreateProgress:input_type -> progress.CreateProgressRequest
	10, // 14: progress.ProgressSvc.GetProgress:input_type -> general.GetRequest
	3,  // 15: progress.ProgressSvc.UpdateProgress:input_type -> progress.UpdateProgressRequest
	4,  // 16: progress.ProgressSvc.UpdateCollectionProgress:input_type -> progress.UpdateCollectionProgressRequest
	11, // 17: progress.ProgressSvc.DeleteProgress:input_type -> general.ResourceId
	12, // 18: progress.ProgressSvc.DeleteCollectionProgress:input_type -> general.ListOptions
	12, // 19: progress.ProgressSvc.ListProgress:input_type -> general.ListOptions
	11, // 20: progress.ProgressSvc.CreateProgress:output_type -> general.ResourceId
	1,  // 21: progress.ProgressSvc.GetProgress:output_type -> progress.Progress
	13, // 22: progress.ProgressSvc.UpdateProgress:output_type -> google.protobuf.Empty
	13, // 23: progress.ProgressSvc.UpdateCollectionProgress:output_type -> google.protobuf.Empty
	13, // 24: progress.ProgressSvc.DeleteProgress:output_type -> google.protobuf.Empty
	13, // 25: progress.ProgressSvc.DeleteCollectionProgress:output_type -> google.protobuf.Empty
	5,  // 26: progress.ProgressSvc.ListProgress:output_type -> progress.ListProgressesResponse
	20, // [20:27] is the sub-list for method output_type
	13, // [13:20] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_progress_progress_proto_init() }
func file_progress_progress_proto_init() {
	if File_progress_progress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_progress_progress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProgressRequest); i {
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
		file_progress_progress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Progress); i {
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
		file_progress_progress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressStep); i {
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
		file_progress_progress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProgressRequest); i {
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
		file_progress_progress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCollectionProgressRequest); i {
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
		file_progress_progress_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProgressesResponse); i {
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
			RawDescriptor: file_progress_progress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_progress_progress_proto_goTypes,
		DependencyIndexes: file_progress_progress_proto_depIdxs,
		MessageInfos:      file_progress_progress_proto_msgTypes,
	}.Build()
	File_progress_progress_proto = out.File
	file_progress_progress_proto_rawDesc = nil
	file_progress_progress_proto_goTypes = nil
	file_progress_progress_proto_depIdxs = nil
}