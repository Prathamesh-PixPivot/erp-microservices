// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: internal/api/protobufs/gstr9c.proto

package protobufs

import (
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

type GSTR9CRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gstin                   string `protobuf:"bytes,1,opt,name=gstin,proto3" json:"gstin,omitempty"`
	ReturnPeriod            string `protobuf:"bytes,2,opt,name=return_period,json=returnPeriod,proto3" json:"return_period,omitempty"`
	AuditDetails            string `protobuf:"bytes,3,opt,name=audit_details,json=auditDetails,proto3" json:"audit_details,omitempty"`
	ReconciliationStatement string `protobuf:"bytes,4,opt,name=reconciliation_statement,json=reconciliationStatement,proto3" json:"reconciliation_statement,omitempty"`
}

func (x *GSTR9CRequest) Reset() {
	*x = GSTR9CRequest{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CRequest) ProtoMessage() {}

func (x *GSTR9CRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CRequest.ProtoReflect.Descriptor instead.
func (*GSTR9CRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{0}
}

func (x *GSTR9CRequest) GetGstin() string {
	if x != nil {
		return x.Gstin
	}
	return ""
}

func (x *GSTR9CRequest) GetReturnPeriod() string {
	if x != nil {
		return x.ReturnPeriod
	}
	return ""
}

func (x *GSTR9CRequest) GetAuditDetails() string {
	if x != nil {
		return x.AuditDetails
	}
	return ""
}

func (x *GSTR9CRequest) GetReconciliationStatement() string {
	if x != nil {
		return x.ReconciliationStatement
	}
	return ""
}

type GSTR9CResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefId  string `protobuf:"bytes,1,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GSTR9CResponse) Reset() {
	*x = GSTR9CResponse{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CResponse) ProtoMessage() {}

func (x *GSTR9CResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CResponse.ProtoReflect.Descriptor instead.
func (*GSTR9CResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{1}
}

func (x *GSTR9CResponse) GetRefId() string {
	if x != nil {
		return x.RefId
	}
	return ""
}

func (x *GSTR9CResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GSTR9CSubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gstin        string `protobuf:"bytes,1,opt,name=gstin,proto3" json:"gstin,omitempty"`
	ReturnPeriod string `protobuf:"bytes,2,opt,name=return_period,json=returnPeriod,proto3" json:"return_period,omitempty"`
}

func (x *GSTR9CSubmitRequest) Reset() {
	*x = GSTR9CSubmitRequest{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CSubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CSubmitRequest) ProtoMessage() {}

func (x *GSTR9CSubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CSubmitRequest.ProtoReflect.Descriptor instead.
func (*GSTR9CSubmitRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{2}
}

func (x *GSTR9CSubmitRequest) GetGstin() string {
	if x != nil {
		return x.Gstin
	}
	return ""
}

func (x *GSTR9CSubmitRequest) GetReturnPeriod() string {
	if x != nil {
		return x.ReturnPeriod
	}
	return ""
}

type GSTR9CSubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arn    string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GSTR9CSubmitResponse) Reset() {
	*x = GSTR9CSubmitResponse{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CSubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CSubmitResponse) ProtoMessage() {}

func (x *GSTR9CSubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CSubmitResponse.ProtoReflect.Descriptor instead.
func (*GSTR9CSubmitResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{3}
}

func (x *GSTR9CSubmitResponse) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *GSTR9CSubmitResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GSTR9CFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gstin        string `protobuf:"bytes,1,opt,name=gstin,proto3" json:"gstin,omitempty"`
	ReturnPeriod string `protobuf:"bytes,2,opt,name=return_period,json=returnPeriod,proto3" json:"return_period,omitempty"`
	Arn          string `protobuf:"bytes,3,opt,name=arn,proto3" json:"arn,omitempty"`
}

func (x *GSTR9CFileRequest) Reset() {
	*x = GSTR9CFileRequest{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CFileRequest) ProtoMessage() {}

func (x *GSTR9CFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CFileRequest.ProtoReflect.Descriptor instead.
func (*GSTR9CFileRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{4}
}

func (x *GSTR9CFileRequest) GetGstin() string {
	if x != nil {
		return x.Gstin
	}
	return ""
}

func (x *GSTR9CFileRequest) GetReturnPeriod() string {
	if x != nil {
		return x.ReturnPeriod
	}
	return ""
}

func (x *GSTR9CFileRequest) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

type GSTR9CFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilingStatus string `protobuf:"bytes,1,opt,name=filing_status,json=filingStatus,proto3" json:"filing_status,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GSTR9CFileResponse) Reset() {
	*x = GSTR9CFileResponse{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CFileResponse) ProtoMessage() {}

func (x *GSTR9CFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CFileResponse.ProtoReflect.Descriptor instead.
func (*GSTR9CFileResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{5}
}

func (x *GSTR9CFileResponse) GetFilingStatus() string {
	if x != nil {
		return x.FilingStatus
	}
	return ""
}

func (x *GSTR9CFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GSTR9CStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gstin        string `protobuf:"bytes,1,opt,name=gstin,proto3" json:"gstin,omitempty"`
	ReturnPeriod string `protobuf:"bytes,2,opt,name=return_period,json=returnPeriod,proto3" json:"return_period,omitempty"`
	Arn          string `protobuf:"bytes,3,opt,name=arn,proto3" json:"arn,omitempty"`
}

func (x *GSTR9CStatusRequest) Reset() {
	*x = GSTR9CStatusRequest{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CStatusRequest) ProtoMessage() {}

func (x *GSTR9CStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CStatusRequest.ProtoReflect.Descriptor instead.
func (*GSTR9CStatusRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{6}
}

func (x *GSTR9CStatusRequest) GetGstin() string {
	if x != nil {
		return x.Gstin
	}
	return ""
}

func (x *GSTR9CStatusRequest) GetReturnPeriod() string {
	if x != nil {
		return x.ReturnPeriod
	}
	return ""
}

func (x *GSTR9CStatusRequest) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

type GSTR9CStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GSTR9CStatusResponse) Reset() {
	*x = GSTR9CStatusResponse{}
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GSTR9CStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GSTR9CStatusResponse) ProtoMessage() {}

func (x *GSTR9CStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_protobufs_gstr9c_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GSTR9CStatusResponse.ProtoReflect.Descriptor instead.
func (*GSTR9CStatusResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_protobufs_gstr9c_proto_rawDescGZIP(), []int{7}
}

func (x *GSTR9CStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GSTR9CStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_api_protobufs_gstr9c_proto protoreflect.FileDescriptor

var file_internal_api_protobufs_gstr9c_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x67, 0x73, 0x74, 0x72, 0x39, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0d, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x73, 0x74, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x73, 0x74, 0x69, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x18, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x3f, 0x0a, 0x0e, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x13, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x73,
	0x74, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x73, 0x74, 0x69, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x60, 0x0a, 0x11, 0x47, 0x53, 0x54, 0x52, 0x39,
	0x43, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x73, 0x74, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x73, 0x74,
	0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x22, 0x53, 0x0a, 0x12, 0x47, 0x53, 0x54,
	0x52, 0x39, 0x43, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x62,
	0x0a, 0x13, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x73, 0x74, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x73, 0x74, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x72, 0x6e, 0x22, 0x48, 0x0a, 0x14, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf2, 0x01, 0x0a,
	0x0d, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d,
	0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x12, 0x0e, 0x2e, 0x47,
	0x53, 0x54, 0x52, 0x39, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47,
	0x53, 0x54, 0x52, 0x39, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x12, 0x14, 0x2e,
	0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x65, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x12, 0x12, 0x2e, 0x47, 0x53, 0x54, 0x52, 0x39,
	0x43, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47,
	0x53, 0x54, 0x52, 0x39, 0x43, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x2e, 0x47, 0x53, 0x54, 0x52, 0x39, 0x43, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x53, 0x54,
	0x52, 0x39, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x67, 0x73, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_protobufs_gstr9c_proto_rawDescOnce sync.Once
	file_internal_api_protobufs_gstr9c_proto_rawDescData = file_internal_api_protobufs_gstr9c_proto_rawDesc
)

func file_internal_api_protobufs_gstr9c_proto_rawDescGZIP() []byte {
	file_internal_api_protobufs_gstr9c_proto_rawDescOnce.Do(func() {
		file_internal_api_protobufs_gstr9c_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_protobufs_gstr9c_proto_rawDescData)
	})
	return file_internal_api_protobufs_gstr9c_proto_rawDescData
}

var file_internal_api_protobufs_gstr9c_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_api_protobufs_gstr9c_proto_goTypes = []any{
	(*GSTR9CRequest)(nil),        // 0: GSTR9CRequest
	(*GSTR9CResponse)(nil),       // 1: GSTR9CResponse
	(*GSTR9CSubmitRequest)(nil),  // 2: GSTR9CSubmitRequest
	(*GSTR9CSubmitResponse)(nil), // 3: GSTR9CSubmitResponse
	(*GSTR9CFileRequest)(nil),    // 4: GSTR9CFileRequest
	(*GSTR9CFileResponse)(nil),   // 5: GSTR9CFileResponse
	(*GSTR9CStatusRequest)(nil),  // 6: GSTR9CStatusRequest
	(*GSTR9CStatusResponse)(nil), // 7: GSTR9CStatusResponse
}
var file_internal_api_protobufs_gstr9c_proto_depIdxs = []int32{
	0, // 0: GSTR9CService.SaveGSTR9C:input_type -> GSTR9CRequest
	2, // 1: GSTR9CService.SubmitGSTR9C:input_type -> GSTR9CSubmitRequest
	4, // 2: GSTR9CService.FileGSTR9C:input_type -> GSTR9CFileRequest
	6, // 3: GSTR9CService.GetGSTR9CStatus:input_type -> GSTR9CStatusRequest
	1, // 4: GSTR9CService.SaveGSTR9C:output_type -> GSTR9CResponse
	3, // 5: GSTR9CService.SubmitGSTR9C:output_type -> GSTR9CSubmitResponse
	5, // 6: GSTR9CService.FileGSTR9C:output_type -> GSTR9CFileResponse
	7, // 7: GSTR9CService.GetGSTR9CStatus:output_type -> GSTR9CStatusResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_api_protobufs_gstr9c_proto_init() }
func file_internal_api_protobufs_gstr9c_proto_init() {
	if File_internal_api_protobufs_gstr9c_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_api_protobufs_gstr9c_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_protobufs_gstr9c_proto_goTypes,
		DependencyIndexes: file_internal_api_protobufs_gstr9c_proto_depIdxs,
		MessageInfos:      file_internal_api_protobufs_gstr9c_proto_msgTypes,
	}.Build()
	File_internal_api_protobufs_gstr9c_proto = out.File
	file_internal_api_protobufs_gstr9c_proto_rawDesc = nil
	file_internal_api_protobufs_gstr9c_proto_goTypes = nil
	file_internal_api_protobufs_gstr9c_proto_depIdxs = nil
}
