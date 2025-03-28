// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/performance_kpi.proto

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

// Performance KPI message
type PerformanceKPI struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReviewId      uint64                 `protobuf:"varint,2,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	KpiName       string                 `protobuf:"bytes,3,opt,name=kpi_name,json=kpiName,proto3" json:"kpi_name,omitempty"`
	Score         float64                `protobuf:"fixed64,4,opt,name=score,proto3" json:"score,omitempty"`
	Comments      string                 `protobuf:"bytes,5,opt,name=comments,proto3" json:"comments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PerformanceKPI) Reset() {
	*x = PerformanceKPI{}
	mi := &file_proto_performance_kpi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PerformanceKPI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformanceKPI) ProtoMessage() {}

func (x *PerformanceKPI) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformanceKPI.ProtoReflect.Descriptor instead.
func (*PerformanceKPI) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{0}
}

func (x *PerformanceKPI) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PerformanceKPI) GetReviewId() uint64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *PerformanceKPI) GetKpiName() string {
	if x != nil {
		return x.KpiName
	}
	return ""
}

func (x *PerformanceKPI) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *PerformanceKPI) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

// Request to create a KPI entry
type CreatePerformanceKPIRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewId      uint64                 `protobuf:"varint,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	KpiName       string                 `protobuf:"bytes,2,opt,name=kpi_name,json=kpiName,proto3" json:"kpi_name,omitempty"`
	Score         float64                `protobuf:"fixed64,3,opt,name=score,proto3" json:"score,omitempty"`
	Comments      string                 `protobuf:"bytes,4,opt,name=comments,proto3" json:"comments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePerformanceKPIRequest) Reset() {
	*x = CreatePerformanceKPIRequest{}
	mi := &file_proto_performance_kpi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePerformanceKPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePerformanceKPIRequest) ProtoMessage() {}

func (x *CreatePerformanceKPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePerformanceKPIRequest.ProtoReflect.Descriptor instead.
func (*CreatePerformanceKPIRequest) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePerformanceKPIRequest) GetReviewId() uint64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *CreatePerformanceKPIRequest) GetKpiName() string {
	if x != nil {
		return x.KpiName
	}
	return ""
}

func (x *CreatePerformanceKPIRequest) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CreatePerformanceKPIRequest) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

// Response containing the created KPI details
type PerformanceKPIResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kpi           *PerformanceKPI        `protobuf:"bytes,1,opt,name=kpi,proto3" json:"kpi,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PerformanceKPIResponse) Reset() {
	*x = PerformanceKPIResponse{}
	mi := &file_proto_performance_kpi_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PerformanceKPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformanceKPIResponse) ProtoMessage() {}

func (x *PerformanceKPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformanceKPIResponse.ProtoReflect.Descriptor instead.
func (*PerformanceKPIResponse) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{2}
}

func (x *PerformanceKPIResponse) GetKpi() *PerformanceKPI {
	if x != nil {
		return x.Kpi
	}
	return nil
}

// Request to get a KPI by ID
type GetPerformanceKPIRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	KpiId         uint64                 `protobuf:"varint,1,opt,name=kpi_id,json=kpiId,proto3" json:"kpi_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPerformanceKPIRequest) Reset() {
	*x = GetPerformanceKPIRequest{}
	mi := &file_proto_performance_kpi_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPerformanceKPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPerformanceKPIRequest) ProtoMessage() {}

func (x *GetPerformanceKPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPerformanceKPIRequest.ProtoReflect.Descriptor instead.
func (*GetPerformanceKPIRequest) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{3}
}

func (x *GetPerformanceKPIRequest) GetKpiId() uint64 {
	if x != nil {
		return x.KpiId
	}
	return 0
}

// Request to list KPIs for a specific review
type ListPerformanceKPIsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewId      uint64                 `protobuf:"varint,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPerformanceKPIsRequest) Reset() {
	*x = ListPerformanceKPIsRequest{}
	mi := &file_proto_performance_kpi_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPerformanceKPIsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPerformanceKPIsRequest) ProtoMessage() {}

func (x *ListPerformanceKPIsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPerformanceKPIsRequest.ProtoReflect.Descriptor instead.
func (*ListPerformanceKPIsRequest) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{4}
}

func (x *ListPerformanceKPIsRequest) GetReviewId() uint64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *ListPerformanceKPIsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPerformanceKPIsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Response containing a paginated list of KPIs
type ListPerformanceKPIsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Kpis          []*PerformanceKPI      `protobuf:"bytes,4,rep,name=kpis,proto3" json:"kpis,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPerformanceKPIsResponse) Reset() {
	*x = ListPerformanceKPIsResponse{}
	mi := &file_proto_performance_kpi_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPerformanceKPIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPerformanceKPIsResponse) ProtoMessage() {}

func (x *ListPerformanceKPIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPerformanceKPIsResponse.ProtoReflect.Descriptor instead.
func (*ListPerformanceKPIsResponse) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{5}
}

func (x *ListPerformanceKPIsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPerformanceKPIsResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPerformanceKPIsResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPerformanceKPIsResponse) GetKpis() []*PerformanceKPI {
	if x != nil {
		return x.Kpis
	}
	return nil
}

// Request to update a KPI
type UpdatePerformanceKPIRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	KpiId         uint64                 `protobuf:"varint,1,opt,name=kpi_id,json=kpiId,proto3" json:"kpi_id,omitempty"`
	KpiName       *string                `protobuf:"bytes,2,opt,name=kpi_name,json=kpiName,proto3,oneof" json:"kpi_name,omitempty"`
	Score         *float64               `protobuf:"fixed64,3,opt,name=score,proto3,oneof" json:"score,omitempty"`
	Comments      *string                `protobuf:"bytes,4,opt,name=comments,proto3,oneof" json:"comments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePerformanceKPIRequest) Reset() {
	*x = UpdatePerformanceKPIRequest{}
	mi := &file_proto_performance_kpi_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePerformanceKPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePerformanceKPIRequest) ProtoMessage() {}

func (x *UpdatePerformanceKPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePerformanceKPIRequest.ProtoReflect.Descriptor instead.
func (*UpdatePerformanceKPIRequest) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePerformanceKPIRequest) GetKpiId() uint64 {
	if x != nil {
		return x.KpiId
	}
	return 0
}

func (x *UpdatePerformanceKPIRequest) GetKpiName() string {
	if x != nil && x.KpiName != nil {
		return *x.KpiName
	}
	return ""
}

func (x *UpdatePerformanceKPIRequest) GetScore() float64 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *UpdatePerformanceKPIRequest) GetComments() string {
	if x != nil && x.Comments != nil {
		return *x.Comments
	}
	return ""
}

// Request to delete a KPI
type DeletePerformanceKPIRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	KpiId         uint64                 `protobuf:"varint,1,opt,name=kpi_id,json=kpiId,proto3" json:"kpi_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePerformanceKPIRequest) Reset() {
	*x = DeletePerformanceKPIRequest{}
	mi := &file_proto_performance_kpi_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePerformanceKPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePerformanceKPIRequest) ProtoMessage() {}

func (x *DeletePerformanceKPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_performance_kpi_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePerformanceKPIRequest.ProtoReflect.Descriptor instead.
func (*DeletePerformanceKPIRequest) Descriptor() ([]byte, []int) {
	return file_proto_performance_kpi_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePerformanceKPIRequest) GetKpiId() uint64 {
	if x != nil {
		return x.KpiId
	}
	return 0
}

var File_proto_performance_kpi_proto protoreflect.FileDescriptor

var file_proto_performance_kpi_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x6b, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68,
	0x72, 0x6d, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x4b, 0x50, 0x49, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6b, 0x70, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x70, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x87, 0x01,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x70,
	0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x70,
	0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x16, 0x50, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x03, 0x6b, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x4b, 0x50, 0x49, 0x52, 0x03, 0x6b, 0x70, 0x69, 0x22, 0x31, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6b, 0x70, 0x69, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x1a,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b,
	0x50, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6b, 0x70, 0x69,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x04, 0x6b,
	0x70, 0x69, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6b, 0x70, 0x69, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x6b, 0x70,
	0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x6b, 0x70, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6b, 0x70, 0x69, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x34, 0x0a, 0x1b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b,
	0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x70, 0x69,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6b, 0x70, 0x69, 0x49, 0x64,
	0x32, 0xc5, 0x03, 0x0a, 0x15, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x4b, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b,
	0x50, 0x49, 0x12, 0x21, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x50, 0x65, 0x72,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x12, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50,
	0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x73, 0x12, 0x20, 0x2e,
	0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x12, 0x21, 0x2e, 0x68, 0x72, 0x6d,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x12, 0x21, 0x2e,
	0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x68, 0x72, 0x6d, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x72, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_proto_performance_kpi_proto_rawDescOnce sync.Once
	file_proto_performance_kpi_proto_rawDescData []byte
)

func file_proto_performance_kpi_proto_rawDescGZIP() []byte {
	file_proto_performance_kpi_proto_rawDescOnce.Do(func() {
		file_proto_performance_kpi_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_performance_kpi_proto_rawDesc), len(file_proto_performance_kpi_proto_rawDesc)))
	})
	return file_proto_performance_kpi_proto_rawDescData
}

var file_proto_performance_kpi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_performance_kpi_proto_goTypes = []any{
	(*PerformanceKPI)(nil),              // 0: hrms.PerformanceKPI
	(*CreatePerformanceKPIRequest)(nil), // 1: hrms.CreatePerformanceKPIRequest
	(*PerformanceKPIResponse)(nil),      // 2: hrms.PerformanceKPIResponse
	(*GetPerformanceKPIRequest)(nil),    // 3: hrms.GetPerformanceKPIRequest
	(*ListPerformanceKPIsRequest)(nil),  // 4: hrms.ListPerformanceKPIsRequest
	(*ListPerformanceKPIsResponse)(nil), // 5: hrms.ListPerformanceKPIsResponse
	(*UpdatePerformanceKPIRequest)(nil), // 6: hrms.UpdatePerformanceKPIRequest
	(*DeletePerformanceKPIRequest)(nil), // 7: hrms.DeletePerformanceKPIRequest
	(*emptypb.Empty)(nil),               // 8: google.protobuf.Empty
}
var file_proto_performance_kpi_proto_depIdxs = []int32{
	0, // 0: hrms.PerformanceKPIResponse.kpi:type_name -> hrms.PerformanceKPI
	0, // 1: hrms.ListPerformanceKPIsResponse.kpis:type_name -> hrms.PerformanceKPI
	1, // 2: hrms.PerformanceKPIService.CreatePerformanceKPI:input_type -> hrms.CreatePerformanceKPIRequest
	3, // 3: hrms.PerformanceKPIService.GetPerformanceKPI:input_type -> hrms.GetPerformanceKPIRequest
	4, // 4: hrms.PerformanceKPIService.ListPerformanceKPIs:input_type -> hrms.ListPerformanceKPIsRequest
	6, // 5: hrms.PerformanceKPIService.UpdatePerformanceKPI:input_type -> hrms.UpdatePerformanceKPIRequest
	7, // 6: hrms.PerformanceKPIService.DeletePerformanceKPI:input_type -> hrms.DeletePerformanceKPIRequest
	2, // 7: hrms.PerformanceKPIService.CreatePerformanceKPI:output_type -> hrms.PerformanceKPIResponse
	2, // 8: hrms.PerformanceKPIService.GetPerformanceKPI:output_type -> hrms.PerformanceKPIResponse
	5, // 9: hrms.PerformanceKPIService.ListPerformanceKPIs:output_type -> hrms.ListPerformanceKPIsResponse
	8, // 10: hrms.PerformanceKPIService.UpdatePerformanceKPI:output_type -> google.protobuf.Empty
	8, // 11: hrms.PerformanceKPIService.DeletePerformanceKPI:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_performance_kpi_proto_init() }
func file_proto_performance_kpi_proto_init() {
	if File_proto_performance_kpi_proto != nil {
		return
	}
	file_proto_performance_kpi_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_performance_kpi_proto_rawDesc), len(file_proto_performance_kpi_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_performance_kpi_proto_goTypes,
		DependencyIndexes: file_proto_performance_kpi_proto_depIdxs,
		MessageInfos:      file_proto_performance_kpi_proto_msgTypes,
	}.Build()
	File_proto_performance_kpi_proto = out.File
	file_proto_performance_kpi_proto_goTypes = nil
	file_proto_performance_kpi_proto_depIdxs = nil
}
