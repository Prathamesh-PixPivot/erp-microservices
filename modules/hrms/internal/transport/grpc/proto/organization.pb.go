// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/organization.proto

package hrms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Organization model
type Organization struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address       string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Organization) Reset() {
	*x = Organization{}
	mi := &file_proto_organization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Organization) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Organization) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Organization) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Organization) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Organization) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Request for creating an organization
type CreateOrganizationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrganizationRequest) Reset() {
	*x = CreateOrganizationRequest{}
	mi := &file_proto_organization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrganizationRequest) ProtoMessage() {}

func (x *CreateOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrganizationRequest.ProtoReflect.Descriptor instead.
func (*CreateOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrganizationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOrganizationRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateOrganizationRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateOrganizationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Response containing organization details
type OrganizationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Organization  *Organization          `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationResponse) Reset() {
	*x = OrganizationResponse{}
	mi := &file_proto_organization_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationResponse) ProtoMessage() {}

func (x *OrganizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationResponse.ProtoReflect.Descriptor instead.
func (*OrganizationResponse) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationResponse) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

// Request for fetching an organization by ID
type GetOrganizationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgId         uint64                 `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrganizationRequest) Reset() {
	*x = GetOrganizationRequest{}
	mi := &file_proto_organization_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrganizationRequest) ProtoMessage() {}

func (x *GetOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrganizationRequest.ProtoReflect.Descriptor instead.
func (*GetOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrganizationRequest) GetOrgId() uint64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

// Request for updating an organization
type UpdateOrganizationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgId         uint64                 `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Address       *string                `protobuf:"bytes,3,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Phone         *string                `protobuf:"bytes,4,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Email         *string                `protobuf:"bytes,5,opt,name=email,proto3,oneof" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrganizationRequest) Reset() {
	*x = UpdateOrganizationRequest{}
	mi := &file_proto_organization_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrganizationRequest) ProtoMessage() {}

func (x *UpdateOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrganizationRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOrganizationRequest) GetOrgId() uint64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *UpdateOrganizationRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateOrganizationRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *UpdateOrganizationRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateOrganizationRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

// Request for deleting an organization
type DeleteOrganizationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgId         uint64                 `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOrganizationRequest) Reset() {
	*x = DeleteOrganizationRequest{}
	mi := &file_proto_organization_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrganizationRequest) ProtoMessage() {}

func (x *DeleteOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrganizationRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteOrganizationRequest) GetOrgId() uint64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

// Request for listing organizations with pagination and search filter
type ListOrganizationsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search        *string                `protobuf:"bytes,3,opt,name=search,proto3,oneof" json:"search,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrganizationsRequest) Reset() {
	*x = ListOrganizationsRequest{}
	mi := &file_proto_organization_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrganizationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsRequest) ProtoMessage() {}

func (x *ListOrganizationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsRequest.ProtoReflect.Descriptor instead.
func (*ListOrganizationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrganizationsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListOrganizationsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListOrganizationsRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

// Response containing a paginated list of organizations
type ListOrganizationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Organizations []*Organization        `protobuf:"bytes,4,rep,name=organizations,proto3" json:"organizations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrganizationsResponse) Reset() {
	*x = ListOrganizationsResponse{}
	mi := &file_proto_organization_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrganizationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsResponse) ProtoMessage() {}

func (x *ListOrganizationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_organization_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsResponse.ProtoReflect.Descriptor instead.
func (*ListOrganizationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_organization_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrganizationsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListOrganizationsResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListOrganizationsResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListOrganizationsResponse) GetOrganizations() []*Organization {
	if x != nil {
		return x.Organizations
	}
	return nil
}

var File_proto_organization_proto protoreflect.FileDescriptor

var file_proto_organization_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x72, 0x6d, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee,
	0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x75, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4e, 0x0a, 0x14, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x32, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x99, 0x01, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xa9, 0x03, 0x0a, 0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x68, 0x72, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x68, 0x72, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_organization_proto_rawDescOnce sync.Once
	file_proto_organization_proto_rawDescData []byte
)

func file_proto_organization_proto_rawDescGZIP() []byte {
	file_proto_organization_proto_rawDescOnce.Do(func() {
		file_proto_organization_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_organization_proto_rawDesc), len(file_proto_organization_proto_rawDesc)))
	})
	return file_proto_organization_proto_rawDescData
}

var file_proto_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_organization_proto_goTypes = []any{
	(*Organization)(nil),              // 0: hrms.Organization
	(*CreateOrganizationRequest)(nil), // 1: hrms.CreateOrganizationRequest
	(*OrganizationResponse)(nil),      // 2: hrms.OrganizationResponse
	(*GetOrganizationRequest)(nil),    // 3: hrms.GetOrganizationRequest
	(*UpdateOrganizationRequest)(nil), // 4: hrms.UpdateOrganizationRequest
	(*DeleteOrganizationRequest)(nil), // 5: hrms.DeleteOrganizationRequest
	(*ListOrganizationsRequest)(nil),  // 6: hrms.ListOrganizationsRequest
	(*ListOrganizationsResponse)(nil), // 7: hrms.ListOrganizationsResponse
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 9: google.protobuf.Empty
}
var file_proto_organization_proto_depIdxs = []int32{
	8, // 0: hrms.Organization.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: hrms.Organization.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: hrms.OrganizationResponse.organization:type_name -> hrms.Organization
	0, // 3: hrms.ListOrganizationsResponse.organizations:type_name -> hrms.Organization
	1, // 4: hrms.OrganizationService.CreateOrganization:input_type -> hrms.CreateOrganizationRequest
	3, // 5: hrms.OrganizationService.GetOrganization:input_type -> hrms.GetOrganizationRequest
	4, // 6: hrms.OrganizationService.UpdateOrganization:input_type -> hrms.UpdateOrganizationRequest
	5, // 7: hrms.OrganizationService.DeleteOrganization:input_type -> hrms.DeleteOrganizationRequest
	6, // 8: hrms.OrganizationService.ListOrganizations:input_type -> hrms.ListOrganizationsRequest
	2, // 9: hrms.OrganizationService.CreateOrganization:output_type -> hrms.OrganizationResponse
	2, // 10: hrms.OrganizationService.GetOrganization:output_type -> hrms.OrganizationResponse
	9, // 11: hrms.OrganizationService.UpdateOrganization:output_type -> google.protobuf.Empty
	9, // 12: hrms.OrganizationService.DeleteOrganization:output_type -> google.protobuf.Empty
	7, // 13: hrms.OrganizationService.ListOrganizations:output_type -> hrms.ListOrganizationsResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_organization_proto_init() }
func file_proto_organization_proto_init() {
	if File_proto_organization_proto != nil {
		return
	}
	file_proto_organization_proto_msgTypes[4].OneofWrappers = []any{}
	file_proto_organization_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_organization_proto_rawDesc), len(file_proto_organization_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_organization_proto_goTypes,
		DependencyIndexes: file_proto_organization_proto_depIdxs,
		MessageInfos:      file_proto_organization_proto_msgTypes,
	}.Build()
	File_proto_organization_proto = out.File
	file_proto_organization_proto_goTypes = nil
	file_proto_organization_proto_depIdxs = nil
}
