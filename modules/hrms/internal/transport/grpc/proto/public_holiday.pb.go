// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/public_holiday.proto

package hrms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Public Holiday message structure.
type PublicHoliday struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationId uint64                 `protobuf:"varint,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Date           *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PublicHoliday) Reset() {
	*x = PublicHoliday{}
	mi := &file_proto_public_holiday_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicHoliday) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicHoliday) ProtoMessage() {}

func (x *PublicHoliday) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicHoliday.ProtoReflect.Descriptor instead.
func (*PublicHoliday) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{0}
}

func (x *PublicHoliday) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PublicHoliday) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *PublicHoliday) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublicHoliday) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

// Request & Response messages.
type CreatePublicHolidayRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OrganizationId uint64                 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date           *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreatePublicHolidayRequest) Reset() {
	*x = CreatePublicHolidayRequest{}
	mi := &file_proto_public_holiday_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePublicHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePublicHolidayRequest) ProtoMessage() {}

func (x *CreatePublicHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePublicHolidayRequest.ProtoReflect.Descriptor instead.
func (*CreatePublicHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePublicHolidayRequest) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *CreatePublicHolidayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePublicHolidayRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type PublicHolidayResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Holiday       *PublicHoliday         `protobuf:"bytes,1,opt,name=holiday,proto3" json:"holiday,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicHolidayResponse) Reset() {
	*x = PublicHolidayResponse{}
	mi := &file_proto_public_holiday_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicHolidayResponse) ProtoMessage() {}

func (x *PublicHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicHolidayResponse.ProtoReflect.Descriptor instead.
func (*PublicHolidayResponse) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{2}
}

func (x *PublicHolidayResponse) GetHoliday() *PublicHoliday {
	if x != nil {
		return x.Holiday
	}
	return nil
}

type GetPublicHolidayRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HolidayId     uint64                 `protobuf:"varint,1,opt,name=holiday_id,json=holidayId,proto3" json:"holiday_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPublicHolidayRequest) Reset() {
	*x = GetPublicHolidayRequest{}
	mi := &file_proto_public_holiday_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPublicHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicHolidayRequest) ProtoMessage() {}

func (x *GetPublicHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicHolidayRequest.ProtoReflect.Descriptor instead.
func (*GetPublicHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{3}
}

func (x *GetPublicHolidayRequest) GetHolidayId() uint64 {
	if x != nil {
		return x.HolidayId
	}
	return 0
}

type ListPublicHolidaysRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OrganizationId uint64                 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Year           *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"` // Optional year filter
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListPublicHolidaysRequest) Reset() {
	*x = ListPublicHolidaysRequest{}
	mi := &file_proto_public_holiday_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPublicHolidaysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicHolidaysRequest) ProtoMessage() {}

func (x *ListPublicHolidaysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicHolidaysRequest.ProtoReflect.Descriptor instead.
func (*ListPublicHolidaysRequest) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{4}
}

func (x *ListPublicHolidaysRequest) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *ListPublicHolidaysRequest) GetYear() *wrapperspb.Int32Value {
	if x != nil {
		return x.Year
	}
	return nil
}

type ListPublicHolidaysResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Holidays      []*PublicHoliday       `protobuf:"bytes,1,rep,name=holidays,proto3" json:"holidays,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPublicHolidaysResponse) Reset() {
	*x = ListPublicHolidaysResponse{}
	mi := &file_proto_public_holiday_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPublicHolidaysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicHolidaysResponse) ProtoMessage() {}

func (x *ListPublicHolidaysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicHolidaysResponse.ProtoReflect.Descriptor instead.
func (*ListPublicHolidaysResponse) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{5}
}

func (x *ListPublicHolidaysResponse) GetHolidays() []*PublicHoliday {
	if x != nil {
		return x.Holidays
	}
	return nil
}

type UpdatePublicHolidayRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HolidayId     uint64                 `protobuf:"varint,1,opt,name=holiday_id,json=holidayId,proto3" json:"holiday_id,omitempty"`
	Updates       map[string]string      `protobuf:"bytes,2,rep,name=updates,proto3" json:"updates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // Dynamic field updates
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePublicHolidayRequest) Reset() {
	*x = UpdatePublicHolidayRequest{}
	mi := &file_proto_public_holiday_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePublicHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePublicHolidayRequest) ProtoMessage() {}

func (x *UpdatePublicHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePublicHolidayRequest.ProtoReflect.Descriptor instead.
func (*UpdatePublicHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePublicHolidayRequest) GetHolidayId() uint64 {
	if x != nil {
		return x.HolidayId
	}
	return 0
}

func (x *UpdatePublicHolidayRequest) GetUpdates() map[string]string {
	if x != nil {
		return x.Updates
	}
	return nil
}

type UpdatePublicHolidayResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePublicHolidayResponse) Reset() {
	*x = UpdatePublicHolidayResponse{}
	mi := &file_proto_public_holiday_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePublicHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePublicHolidayResponse) ProtoMessage() {}

func (x *UpdatePublicHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePublicHolidayResponse.ProtoReflect.Descriptor instead.
func (*UpdatePublicHolidayResponse) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePublicHolidayResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeletePublicHolidayRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HolidayId     uint64                 `protobuf:"varint,1,opt,name=holiday_id,json=holidayId,proto3" json:"holiday_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePublicHolidayRequest) Reset() {
	*x = DeletePublicHolidayRequest{}
	mi := &file_proto_public_holiday_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePublicHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePublicHolidayRequest) ProtoMessage() {}

func (x *DeletePublicHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePublicHolidayRequest.ProtoReflect.Descriptor instead.
func (*DeletePublicHolidayRequest) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePublicHolidayRequest) GetHolidayId() uint64 {
	if x != nil {
		return x.HolidayId
	}
	return 0
}

type DeletePublicHolidayResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePublicHolidayResponse) Reset() {
	*x = DeletePublicHolidayResponse{}
	mi := &file_proto_public_holiday_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePublicHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePublicHolidayResponse) ProtoMessage() {}

func (x *DeletePublicHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_public_holiday_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePublicHolidayResponse.ProtoReflect.Descriptor instead.
func (*DeletePublicHolidayResponse) Descriptor() ([]byte, []int) {
	return file_proto_public_holiday_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePublicHolidayResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_public_holiday_proto protoreflect.FileDescriptor

var file_proto_public_holiday_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x68,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x72,
	0x6d, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x46,
	0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x68, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x07, 0x68,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x22, 0x38, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x49, 0x64,
	0x22, 0x75, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x4d, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x08, 0x68, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x68, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x1b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x3b, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x49, 0x64, 0x22,
	0x37, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xcd, 0x03, 0x0a, 0x14, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x54, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x20, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x72, 0x6d,
	0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x1d, 0x2e, 0x68, 0x72,
	0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x72, 0x6d,
	0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x12, 0x1f, 0x2e,
	0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x20, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x72, 0x6d, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69,
	0x64, 0x61, 0x79, 0x12, 0x20, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x68, 0x72, 0x6d, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x72, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_proto_public_holiday_proto_rawDescOnce sync.Once
	file_proto_public_holiday_proto_rawDescData []byte
)

func file_proto_public_holiday_proto_rawDescGZIP() []byte {
	file_proto_public_holiday_proto_rawDescOnce.Do(func() {
		file_proto_public_holiday_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_public_holiday_proto_rawDesc), len(file_proto_public_holiday_proto_rawDesc)))
	})
	return file_proto_public_holiday_proto_rawDescData
}

var file_proto_public_holiday_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_public_holiday_proto_goTypes = []any{
	(*PublicHoliday)(nil),               // 0: hrms.PublicHoliday
	(*CreatePublicHolidayRequest)(nil),  // 1: hrms.CreatePublicHolidayRequest
	(*PublicHolidayResponse)(nil),       // 2: hrms.PublicHolidayResponse
	(*GetPublicHolidayRequest)(nil),     // 3: hrms.GetPublicHolidayRequest
	(*ListPublicHolidaysRequest)(nil),   // 4: hrms.ListPublicHolidaysRequest
	(*ListPublicHolidaysResponse)(nil),  // 5: hrms.ListPublicHolidaysResponse
	(*UpdatePublicHolidayRequest)(nil),  // 6: hrms.UpdatePublicHolidayRequest
	(*UpdatePublicHolidayResponse)(nil), // 7: hrms.UpdatePublicHolidayResponse
	(*DeletePublicHolidayRequest)(nil),  // 8: hrms.DeletePublicHolidayRequest
	(*DeletePublicHolidayResponse)(nil), // 9: hrms.DeletePublicHolidayResponse
	nil,                                 // 10: hrms.UpdatePublicHolidayRequest.UpdatesEntry
	(*timestamppb.Timestamp)(nil),       // 11: google.protobuf.Timestamp
	(*wrapperspb.Int32Value)(nil),       // 12: google.protobuf.Int32Value
}
var file_proto_public_holiday_proto_depIdxs = []int32{
	11, // 0: hrms.PublicHoliday.date:type_name -> google.protobuf.Timestamp
	11, // 1: hrms.CreatePublicHolidayRequest.date:type_name -> google.protobuf.Timestamp
	0,  // 2: hrms.PublicHolidayResponse.holiday:type_name -> hrms.PublicHoliday
	12, // 3: hrms.ListPublicHolidaysRequest.year:type_name -> google.protobuf.Int32Value
	0,  // 4: hrms.ListPublicHolidaysResponse.holidays:type_name -> hrms.PublicHoliday
	10, // 5: hrms.UpdatePublicHolidayRequest.updates:type_name -> hrms.UpdatePublicHolidayRequest.UpdatesEntry
	1,  // 6: hrms.PublicHolidayService.CreatePublicHoliday:input_type -> hrms.CreatePublicHolidayRequest
	3,  // 7: hrms.PublicHolidayService.GetPublicHoliday:input_type -> hrms.GetPublicHolidayRequest
	4,  // 8: hrms.PublicHolidayService.ListPublicHolidays:input_type -> hrms.ListPublicHolidaysRequest
	6,  // 9: hrms.PublicHolidayService.UpdatePublicHoliday:input_type -> hrms.UpdatePublicHolidayRequest
	8,  // 10: hrms.PublicHolidayService.DeletePublicHoliday:input_type -> hrms.DeletePublicHolidayRequest
	2,  // 11: hrms.PublicHolidayService.CreatePublicHoliday:output_type -> hrms.PublicHolidayResponse
	2,  // 12: hrms.PublicHolidayService.GetPublicHoliday:output_type -> hrms.PublicHolidayResponse
	5,  // 13: hrms.PublicHolidayService.ListPublicHolidays:output_type -> hrms.ListPublicHolidaysResponse
	7,  // 14: hrms.PublicHolidayService.UpdatePublicHoliday:output_type -> hrms.UpdatePublicHolidayResponse
	9,  // 15: hrms.PublicHolidayService.DeletePublicHoliday:output_type -> hrms.DeletePublicHolidayResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_public_holiday_proto_init() }
func file_proto_public_holiday_proto_init() {
	if File_proto_public_holiday_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_public_holiday_proto_rawDesc), len(file_proto_public_holiday_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_public_holiday_proto_goTypes,
		DependencyIndexes: file_proto_public_holiday_proto_depIdxs,
		MessageInfos:      file_proto_public_holiday_proto_msgTypes,
	}.Build()
	File_proto_public_holiday_proto = out.File
	file_proto_public_holiday_proto_goTypes = nil
	file_proto_public_holiday_proto_depIdxs = nil
}
