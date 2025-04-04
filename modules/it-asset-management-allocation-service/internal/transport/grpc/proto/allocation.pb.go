// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc2
// source: allocation.proto

package proto

import (
	"amaa/internal/transport/grpc/proto/common"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AllocateAssetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       string                 `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AssignedTo    string                 `protobuf:"bytes,2,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocateAssetRequest) Reset() {
	*x = AllocateAssetRequest{}
	mi := &file_allocation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateAssetRequest) ProtoMessage() {}

func (x *AllocateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allocation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateAssetRequest.ProtoReflect.Descriptor instead.
func (*AllocateAssetRequest) Descriptor() ([]byte, []int) {
	return file_allocation_proto_rawDescGZIP(), []int{0}
}

func (x *AllocateAssetRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *AllocateAssetRequest) GetAssignedTo() string {
	if x != nil {
		return x.AssignedTo
	}
	return ""
}

type AllocationResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetId        string                 `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AssignedTo     string                 `protobuf:"bytes,3,opt,name=assigned_to,json=assignedTo,proto3" json:"assigned_to,omitempty"`
	AssignmentDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=assignment_date,json=assignmentDate,proto3" json:"assignment_date,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *AllocationResponse) Reset() {
	*x = AllocationResponse{}
	mi := &file_allocation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationResponse) ProtoMessage() {}

func (x *AllocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_allocation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationResponse.ProtoReflect.Descriptor instead.
func (*AllocationResponse) Descriptor() ([]byte, []int) {
	return file_allocation_proto_rawDescGZIP(), []int{1}
}

func (x *AllocationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AllocationResponse) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *AllocationResponse) GetAssignedTo() string {
	if x != nil {
		return x.AssignedTo
	}
	return ""
}

func (x *AllocationResponse) GetAssignmentDate() *timestamppb.Timestamp {
	if x != nil {
		return x.AssignmentDate
	}
	return nil
}

type ReallocateAssetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       string                 `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	NewAssignedTo string                 `protobuf:"bytes,2,opt,name=new_assigned_to,json=newAssignedTo,proto3" json:"new_assigned_to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReallocateAssetRequest) Reset() {
	*x = ReallocateAssetRequest{}
	mi := &file_allocation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReallocateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReallocateAssetRequest) ProtoMessage() {}

func (x *ReallocateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allocation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReallocateAssetRequest.ProtoReflect.Descriptor instead.
func (*ReallocateAssetRequest) Descriptor() ([]byte, []int) {
	return file_allocation_proto_rawDescGZIP(), []int{2}
}

func (x *ReallocateAssetRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *ReallocateAssetRequest) GetNewAssignedTo() string {
	if x != nil {
		return x.NewAssignedTo
	}
	return ""
}

type DeallocateAssetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AllocationId  string                 `protobuf:"bytes,1,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeallocateAssetRequest) Reset() {
	*x = DeallocateAssetRequest{}
	mi := &file_allocation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeallocateAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeallocateAssetRequest) ProtoMessage() {}

func (x *DeallocateAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_allocation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeallocateAssetRequest.ProtoReflect.Descriptor instead.
func (*DeallocateAssetRequest) Descriptor() ([]byte, []int) {
	return file_allocation_proto_rawDescGZIP(), []int{3}
}

func (x *DeallocateAssetRequest) GetAllocationId() string {
	if x != nil {
		return x.AllocationId
	}
	return ""
}

var File_allocation_proto protoreflect.FileDescriptor

var file_allocation_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x14, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54,
	0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x54, 0x6f, 0x12, 0x43, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x5b, 0x0a, 0x16, 0x52, 0x65, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x22, 0x3d, 0x0a, 0x16, 0x44, 0x65, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0x86, 0x02, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0d, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x0f, 0x52, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x12, 0x22, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x0f, 0x44, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x12, 0x22, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_allocation_proto_rawDescOnce sync.Once
	file_allocation_proto_rawDescData []byte
)

func file_allocation_proto_rawDescGZIP() []byte {
	file_allocation_proto_rawDescOnce.Do(func() {
		file_allocation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_allocation_proto_rawDesc), len(file_allocation_proto_rawDesc)))
	})
	return file_allocation_proto_rawDescData
}

var file_allocation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_allocation_proto_goTypes = []any{
	(*AllocateAssetRequest)(nil),   // 0: allocation.AllocateAssetRequest
	(*AllocationResponse)(nil),     // 1: allocation.AllocationResponse
	(*ReallocateAssetRequest)(nil), // 2: allocation.ReallocateAssetRequest
	(*DeallocateAssetRequest)(nil), // 3: allocation.DeallocateAssetRequest
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*common.GenericResponse)(nil), // 5: common.GenericResponse
}
var file_allocation_proto_depIdxs = []int32{
	4, // 0: allocation.AllocationResponse.assignment_date:type_name -> google.protobuf.Timestamp
	0, // 1: allocation.AllocationService.AllocateAsset:input_type -> allocation.AllocateAssetRequest
	2, // 2: allocation.AllocationService.ReallocateAsset:input_type -> allocation.ReallocateAssetRequest
	3, // 3: allocation.AllocationService.DeallocateAsset:input_type -> allocation.DeallocateAssetRequest
	1, // 4: allocation.AllocationService.AllocateAsset:output_type -> allocation.AllocationResponse
	5, // 5: allocation.AllocationService.ReallocateAsset:output_type -> common.GenericResponse
	5, // 6: allocation.AllocationService.DeallocateAsset:output_type -> common.GenericResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_allocation_proto_init() }
func file_allocation_proto_init() {
	if File_allocation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_allocation_proto_rawDesc), len(file_allocation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_allocation_proto_goTypes,
		DependencyIndexes: file_allocation_proto_depIdxs,
		MessageInfos:      file_allocation_proto_msgTypes,
	}.Build()
	File_allocation_proto = out.File
	file_allocation_proto_goTypes = nil
	file_allocation_proto_depIdxs = nil
}
