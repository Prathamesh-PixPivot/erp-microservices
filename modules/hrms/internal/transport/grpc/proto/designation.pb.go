// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/designation.proto

package hrms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Designation struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title          string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Level          string                 `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	HierarchyLevel uint32                 `protobuf:"varint,4,opt,name=hierarchy_level,json=hierarchyLevel,proto3" json:"hierarchy_level,omitempty"`
	DepartmentId   uint64                 `protobuf:"varint,5,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Designation) Reset() {
	*x = Designation{}
	mi := &file_proto_designation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Designation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Designation) ProtoMessage() {}

func (x *Designation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Designation.ProtoReflect.Descriptor instead.
func (*Designation) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{0}
}

func (x *Designation) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Designation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Designation) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Designation) GetHierarchyLevel() uint32 {
	if x != nil {
		return x.HierarchyLevel
	}
	return 0
}

func (x *Designation) GetDepartmentId() uint64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

type CreateDesignationRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Title          string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Level          string                 `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	HierarchyLevel uint32                 `protobuf:"varint,3,opt,name=hierarchy_level,json=hierarchyLevel,proto3" json:"hierarchy_level,omitempty"`
	DepartmentId   uint64                 `protobuf:"varint,4,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateDesignationRequest) Reset() {
	*x = CreateDesignationRequest{}
	mi := &file_proto_designation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDesignationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDesignationRequest) ProtoMessage() {}

func (x *CreateDesignationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDesignationRequest.ProtoReflect.Descriptor instead.
func (*CreateDesignationRequest) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDesignationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateDesignationRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *CreateDesignationRequest) GetHierarchyLevel() uint32 {
	if x != nil {
		return x.HierarchyLevel
	}
	return 0
}

func (x *CreateDesignationRequest) GetDepartmentId() uint64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

type GetDesignationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDesignationRequest) Reset() {
	*x = GetDesignationRequest{}
	mi := &file_proto_designation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDesignationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDesignationRequest) ProtoMessage() {}

func (x *GetDesignationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDesignationRequest.ProtoReflect.Descriptor instead.
func (*GetDesignationRequest) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{2}
}

func (x *GetDesignationRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateDesignationRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title          *string                `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Level          *string                `protobuf:"bytes,3,opt,name=level,proto3,oneof" json:"level,omitempty"`
	HierarchyLevel *uint32                `protobuf:"varint,4,opt,name=hierarchy_level,json=hierarchyLevel,proto3,oneof" json:"hierarchy_level,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateDesignationRequest) Reset() {
	*x = UpdateDesignationRequest{}
	mi := &file_proto_designation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDesignationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDesignationRequest) ProtoMessage() {}

func (x *UpdateDesignationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDesignationRequest.ProtoReflect.Descriptor instead.
func (*UpdateDesignationRequest) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDesignationRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDesignationRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateDesignationRequest) GetLevel() string {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return ""
}

func (x *UpdateDesignationRequest) GetHierarchyLevel() uint32 {
	if x != nil && x.HierarchyLevel != nil {
		return *x.HierarchyLevel
	}
	return 0
}

type DeleteDesignationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDesignationRequest) Reset() {
	*x = DeleteDesignationRequest{}
	mi := &file_proto_designation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDesignationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDesignationRequest) ProtoMessage() {}

func (x *DeleteDesignationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDesignationRequest.ProtoReflect.Descriptor instead.
func (*DeleteDesignationRequest) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDesignationRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListDesignationsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DepartmentId  uint64                 `protobuf:"varint,1,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	Limit         uint32                 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        uint32                 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Search        *string                `protobuf:"bytes,4,opt,name=search,proto3,oneof" json:"search,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDesignationsRequest) Reset() {
	*x = ListDesignationsRequest{}
	mi := &file_proto_designation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDesignationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDesignationsRequest) ProtoMessage() {}

func (x *ListDesignationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDesignationsRequest.ProtoReflect.Descriptor instead.
func (*ListDesignationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{5}
}

func (x *ListDesignationsRequest) GetDepartmentId() uint64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

func (x *ListDesignationsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDesignationsRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListDesignationsRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

type DesignationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Designation   *Designation           `protobuf:"bytes,1,opt,name=designation,proto3" json:"designation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DesignationResponse) Reset() {
	*x = DesignationResponse{}
	mi := &file_proto_designation_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DesignationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesignationResponse) ProtoMessage() {}

func (x *DesignationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesignationResponse.ProtoReflect.Descriptor instead.
func (*DesignationResponse) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{6}
}

func (x *DesignationResponse) GetDesignation() *Designation {
	if x != nil {
		return x.Designation
	}
	return nil
}

type ListDesignationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Designations  []*Designation         `protobuf:"bytes,1,rep,name=designations,proto3" json:"designations,omitempty"`
	TotalCount    int64                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDesignationsResponse) Reset() {
	*x = ListDesignationsResponse{}
	mi := &file_proto_designation_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDesignationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDesignationsResponse) ProtoMessage() {}

func (x *ListDesignationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDesignationsResponse.ProtoReflect.Descriptor instead.
func (*ListDesignationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{7}
}

func (x *ListDesignationsResponse) GetDesignations() []*Designation {
	if x != nil {
		return x.Designations
	}
	return nil
}

func (x *ListDesignationsResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type EmptyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	mi := &file_proto_designation_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designation_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_proto_designation_proto_rawDescGZIP(), []int{8}
}

var File_proto_designation_proto protoreflect.FileDescriptor

var file_proto_designation_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x72, 0x6d, 0x73, 0x22,
	0x97, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x68,
	0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x68, 0x69, 0x65,
	0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb6, 0x01, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f,
	0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0e, 0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63,
	0x68, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x94,
	0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x4a, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x72, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0c, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x99, 0x03, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x68, 0x72, 0x6d,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x68, 0x72, 0x6d,
	0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x51, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x68, 0x72, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x68, 0x72, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_designation_proto_rawDescOnce sync.Once
	file_proto_designation_proto_rawDescData []byte
)

func file_proto_designation_proto_rawDescGZIP() []byte {
	file_proto_designation_proto_rawDescOnce.Do(func() {
		file_proto_designation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_designation_proto_rawDesc), len(file_proto_designation_proto_rawDesc)))
	})
	return file_proto_designation_proto_rawDescData
}

var file_proto_designation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_designation_proto_goTypes = []any{
	(*Designation)(nil),              // 0: hrms.Designation
	(*CreateDesignationRequest)(nil), // 1: hrms.CreateDesignationRequest
	(*GetDesignationRequest)(nil),    // 2: hrms.GetDesignationRequest
	(*UpdateDesignationRequest)(nil), // 3: hrms.UpdateDesignationRequest
	(*DeleteDesignationRequest)(nil), // 4: hrms.DeleteDesignationRequest
	(*ListDesignationsRequest)(nil),  // 5: hrms.ListDesignationsRequest
	(*DesignationResponse)(nil),      // 6: hrms.DesignationResponse
	(*ListDesignationsResponse)(nil), // 7: hrms.ListDesignationsResponse
	(*EmptyResponse)(nil),            // 8: hrms.EmptyResponse
}
var file_proto_designation_proto_depIdxs = []int32{
	0, // 0: hrms.DesignationResponse.designation:type_name -> hrms.Designation
	0, // 1: hrms.ListDesignationsResponse.designations:type_name -> hrms.Designation
	1, // 2: hrms.DesignationService.CreateDesignation:input_type -> hrms.CreateDesignationRequest
	2, // 3: hrms.DesignationService.GetDesignationByID:input_type -> hrms.GetDesignationRequest
	3, // 4: hrms.DesignationService.UpdateDesignation:input_type -> hrms.UpdateDesignationRequest
	4, // 5: hrms.DesignationService.DeleteDesignation:input_type -> hrms.DeleteDesignationRequest
	5, // 6: hrms.DesignationService.ListDesignations:input_type -> hrms.ListDesignationsRequest
	6, // 7: hrms.DesignationService.CreateDesignation:output_type -> hrms.DesignationResponse
	6, // 8: hrms.DesignationService.GetDesignationByID:output_type -> hrms.DesignationResponse
	8, // 9: hrms.DesignationService.UpdateDesignation:output_type -> hrms.EmptyResponse
	8, // 10: hrms.DesignationService.DeleteDesignation:output_type -> hrms.EmptyResponse
	7, // 11: hrms.DesignationService.ListDesignations:output_type -> hrms.ListDesignationsResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_designation_proto_init() }
func file_proto_designation_proto_init() {
	if File_proto_designation_proto != nil {
		return
	}
	file_proto_designation_proto_msgTypes[3].OneofWrappers = []any{}
	file_proto_designation_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_designation_proto_rawDesc), len(file_proto_designation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_designation_proto_goTypes,
		DependencyIndexes: file_proto_designation_proto_depIdxs,
		MessageInfos:      file_proto_designation_proto_msgTypes,
	}.Build()
	File_proto_designation_proto = out.File
	file_proto_designation_proto_goTypes = nil
	file_proto_designation_proto_depIdxs = nil
}
