// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/skill_development.proto

package hrms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ✅ SkillDevelopment now lives here
type SkillDevelopment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReviewId      uint64                 `protobuf:"varint,2,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Skill         string                 `protobuf:"bytes,3,opt,name=skill,proto3" json:"skill,omitempty"`
	Progress      string                 `protobuf:"bytes,4,opt,name=progress,proto3" json:"progress,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SkillDevelopment) Reset() {
	*x = SkillDevelopment{}
	mi := &file_proto_skill_development_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SkillDevelopment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillDevelopment) ProtoMessage() {}

func (x *SkillDevelopment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillDevelopment.ProtoReflect.Descriptor instead.
func (*SkillDevelopment) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{0}
}

func (x *SkillDevelopment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkillDevelopment) GetReviewId() uint64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *SkillDevelopment) GetSkill() string {
	if x != nil {
		return x.Skill
	}
	return ""
}

func (x *SkillDevelopment) GetProgress() string {
	if x != nil {
		return x.Progress
	}
	return ""
}

// Request to create a Skill Development entry
type CreateSkillDevelopmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewId      uint64                 `protobuf:"varint,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Skill         string                 `protobuf:"bytes,2,opt,name=skill,proto3" json:"skill,omitempty"`
	Progress      string                 `protobuf:"bytes,3,opt,name=progress,proto3" json:"progress,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSkillDevelopmentRequest) Reset() {
	*x = CreateSkillDevelopmentRequest{}
	mi := &file_proto_skill_development_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSkillDevelopmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSkillDevelopmentRequest) ProtoMessage() {}

func (x *CreateSkillDevelopmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSkillDevelopmentRequest.ProtoReflect.Descriptor instead.
func (*CreateSkillDevelopmentRequest) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSkillDevelopmentRequest) GetReviewId() uint64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *CreateSkillDevelopmentRequest) GetSkill() string {
	if x != nil {
		return x.Skill
	}
	return ""
}

func (x *CreateSkillDevelopmentRequest) GetProgress() string {
	if x != nil {
		return x.Progress
	}
	return ""
}

// Response containing Skill Development details
type SkillDevelopmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SkillDev      *SkillDevelopment      `protobuf:"bytes,1,opt,name=skill_dev,json=skillDev,proto3" json:"skill_dev,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SkillDevelopmentResponse) Reset() {
	*x = SkillDevelopmentResponse{}
	mi := &file_proto_skill_development_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SkillDevelopmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillDevelopmentResponse) ProtoMessage() {}

func (x *SkillDevelopmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillDevelopmentResponse.ProtoReflect.Descriptor instead.
func (*SkillDevelopmentResponse) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{2}
}

func (x *SkillDevelopmentResponse) GetSkillDev() *SkillDevelopment {
	if x != nil {
		return x.SkillDev
	}
	return nil
}

// Request to get a Skill Development entry by ID
type GetSkillDevelopmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SkillDevId    uint64                 `protobuf:"varint,1,opt,name=skill_dev_id,json=skillDevId,proto3" json:"skill_dev_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSkillDevelopmentRequest) Reset() {
	*x = GetSkillDevelopmentRequest{}
	mi := &file_proto_skill_development_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSkillDevelopmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkillDevelopmentRequest) ProtoMessage() {}

func (x *GetSkillDevelopmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkillDevelopmentRequest.ProtoReflect.Descriptor instead.
func (*GetSkillDevelopmentRequest) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{3}
}

func (x *GetSkillDevelopmentRequest) GetSkillDevId() uint64 {
	if x != nil {
		return x.SkillDevId
	}
	return 0
}

// Request to list Skill Development entries for a specific review
type ListSkillDevelopmentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewId      uint64                 `protobuf:"varint,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSkillDevelopmentsRequest) Reset() {
	*x = ListSkillDevelopmentsRequest{}
	mi := &file_proto_skill_development_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSkillDevelopmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSkillDevelopmentsRequest) ProtoMessage() {}

func (x *ListSkillDevelopmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSkillDevelopmentsRequest.ProtoReflect.Descriptor instead.
func (*ListSkillDevelopmentsRequest) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{4}
}

func (x *ListSkillDevelopmentsRequest) GetReviewId() uint64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *ListSkillDevelopmentsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListSkillDevelopmentsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Response containing a paginated list of Skill Development entries
type ListSkillDevelopmentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	SkillDevs     []*SkillDevelopment    `protobuf:"bytes,4,rep,name=skill_devs,json=skillDevs,proto3" json:"skill_devs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSkillDevelopmentsResponse) Reset() {
	*x = ListSkillDevelopmentsResponse{}
	mi := &file_proto_skill_development_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSkillDevelopmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSkillDevelopmentsResponse) ProtoMessage() {}

func (x *ListSkillDevelopmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSkillDevelopmentsResponse.ProtoReflect.Descriptor instead.
func (*ListSkillDevelopmentsResponse) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{5}
}

func (x *ListSkillDevelopmentsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListSkillDevelopmentsResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListSkillDevelopmentsResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListSkillDevelopmentsResponse) GetSkillDevs() []*SkillDevelopment {
	if x != nil {
		return x.SkillDevs
	}
	return nil
}

// Request to update a Skill Development entry
type UpdateSkillDevelopmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SkillDevId    uint64                 `protobuf:"varint,1,opt,name=skill_dev_id,json=skillDevId,proto3" json:"skill_dev_id,omitempty"`
	Skill         *string                `protobuf:"bytes,2,opt,name=skill,proto3,oneof" json:"skill,omitempty"`
	Progress      *string                `protobuf:"bytes,3,opt,name=progress,proto3,oneof" json:"progress,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSkillDevelopmentRequest) Reset() {
	*x = UpdateSkillDevelopmentRequest{}
	mi := &file_proto_skill_development_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSkillDevelopmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSkillDevelopmentRequest) ProtoMessage() {}

func (x *UpdateSkillDevelopmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSkillDevelopmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateSkillDevelopmentRequest) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSkillDevelopmentRequest) GetSkillDevId() uint64 {
	if x != nil {
		return x.SkillDevId
	}
	return 0
}

func (x *UpdateSkillDevelopmentRequest) GetSkill() string {
	if x != nil && x.Skill != nil {
		return *x.Skill
	}
	return ""
}

func (x *UpdateSkillDevelopmentRequest) GetProgress() string {
	if x != nil && x.Progress != nil {
		return *x.Progress
	}
	return ""
}

// Request to delete a Skill Development entry
type DeleteSkillDevelopmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SkillDevId    uint64                 `protobuf:"varint,1,opt,name=skill_dev_id,json=skillDevId,proto3" json:"skill_dev_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSkillDevelopmentRequest) Reset() {
	*x = DeleteSkillDevelopmentRequest{}
	mi := &file_proto_skill_development_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSkillDevelopmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSkillDevelopmentRequest) ProtoMessage() {}

func (x *DeleteSkillDevelopmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_skill_development_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSkillDevelopmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteSkillDevelopmentRequest) Descriptor() ([]byte, []int) {
	return file_proto_skill_development_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSkillDevelopmentRequest) GetSkillDevId() uint64 {
	if x != nil {
		return x.SkillDevId
	}
	return 0
}

var File_proto_skill_development_proto protoreflect.FileDescriptor

var file_proto_skill_development_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x68, 0x72, 0x6d, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x71, 0x0a, 0x10, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6e, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x18, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x76, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x22, 0x3e, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65,
	0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x44, 0x65, 0x76, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x9a, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x5f, 0x64, 0x65, 0x76, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x72,
	0x6d, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x73, 0x22, 0x94,
	0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x41, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f,
	0x64, 0x65, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x49, 0x64, 0x32, 0xe1, 0x03, 0x0a, 0x17, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x68, 0x72, 0x6d,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68,
	0x72, 0x6d, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x72, 0x6d, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x11, 0x5a, 0x0f,
	0x68, 0x72, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x72, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_skill_development_proto_rawDescOnce sync.Once
	file_proto_skill_development_proto_rawDescData []byte
)

func file_proto_skill_development_proto_rawDescGZIP() []byte {
	file_proto_skill_development_proto_rawDescOnce.Do(func() {
		file_proto_skill_development_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_skill_development_proto_rawDesc), len(file_proto_skill_development_proto_rawDesc)))
	})
	return file_proto_skill_development_proto_rawDescData
}

var file_proto_skill_development_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_skill_development_proto_goTypes = []any{
	(*SkillDevelopment)(nil),              // 0: hrms.SkillDevelopment
	(*CreateSkillDevelopmentRequest)(nil), // 1: hrms.CreateSkillDevelopmentRequest
	(*SkillDevelopmentResponse)(nil),      // 2: hrms.SkillDevelopmentResponse
	(*GetSkillDevelopmentRequest)(nil),    // 3: hrms.GetSkillDevelopmentRequest
	(*ListSkillDevelopmentsRequest)(nil),  // 4: hrms.ListSkillDevelopmentsRequest
	(*ListSkillDevelopmentsResponse)(nil), // 5: hrms.ListSkillDevelopmentsResponse
	(*UpdateSkillDevelopmentRequest)(nil), // 6: hrms.UpdateSkillDevelopmentRequest
	(*DeleteSkillDevelopmentRequest)(nil), // 7: hrms.DeleteSkillDevelopmentRequest
	(*emptypb.Empty)(nil),                 // 8: google.protobuf.Empty
}
var file_proto_skill_development_proto_depIdxs = []int32{
	0, // 0: hrms.SkillDevelopmentResponse.skill_dev:type_name -> hrms.SkillDevelopment
	0, // 1: hrms.ListSkillDevelopmentsResponse.skill_devs:type_name -> hrms.SkillDevelopment
	1, // 2: hrms.SkillDevelopmentService.CreateSkillDevelopment:input_type -> hrms.CreateSkillDevelopmentRequest
	3, // 3: hrms.SkillDevelopmentService.GetSkillDevelopment:input_type -> hrms.GetSkillDevelopmentRequest
	4, // 4: hrms.SkillDevelopmentService.ListSkillDevelopments:input_type -> hrms.ListSkillDevelopmentsRequest
	6, // 5: hrms.SkillDevelopmentService.UpdateSkillDevelopment:input_type -> hrms.UpdateSkillDevelopmentRequest
	7, // 6: hrms.SkillDevelopmentService.DeleteSkillDevelopment:input_type -> hrms.DeleteSkillDevelopmentRequest
	2, // 7: hrms.SkillDevelopmentService.CreateSkillDevelopment:output_type -> hrms.SkillDevelopmentResponse
	2, // 8: hrms.SkillDevelopmentService.GetSkillDevelopment:output_type -> hrms.SkillDevelopmentResponse
	5, // 9: hrms.SkillDevelopmentService.ListSkillDevelopments:output_type -> hrms.ListSkillDevelopmentsResponse
	8, // 10: hrms.SkillDevelopmentService.UpdateSkillDevelopment:output_type -> google.protobuf.Empty
	8, // 11: hrms.SkillDevelopmentService.DeleteSkillDevelopment:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_skill_development_proto_init() }
func file_proto_skill_development_proto_init() {
	if File_proto_skill_development_proto != nil {
		return
	}
	file_proto_skill_development_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_skill_development_proto_rawDesc), len(file_proto_skill_development_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_skill_development_proto_goTypes,
		DependencyIndexes: file_proto_skill_development_proto_depIdxs,
		MessageInfos:      file_proto_skill_development_proto_msgTypes,
	}.Build()
	File_proto_skill_development_proto = out.File
	file_proto_skill_development_proto_goTypes = nil
	file_proto_skill_development_proto_depIdxs = nil
}
