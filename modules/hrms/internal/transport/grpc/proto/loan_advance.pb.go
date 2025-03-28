// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/loan_advance.proto

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

// Enum for loan/advance status (renamed to prevent conflicts)
type LoanStatus int32

const (
	LoanStatus_LOAN_PENDING  LoanStatus = 0
	LoanStatus_LOAN_APPROVED LoanStatus = 1
	LoanStatus_LOAN_REJECTED LoanStatus = 2
)

// Enum value maps for LoanStatus.
var (
	LoanStatus_name = map[int32]string{
		0: "LOAN_PENDING",
		1: "LOAN_APPROVED",
		2: "LOAN_REJECTED",
	}
	LoanStatus_value = map[string]int32{
		"LOAN_PENDING":  0,
		"LOAN_APPROVED": 1,
		"LOAN_REJECTED": 2,
	}
)

func (x LoanStatus) Enum() *LoanStatus {
	p := new(LoanStatus)
	*p = x
	return p
}

func (x LoanStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoanStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_loan_advance_proto_enumTypes[0].Descriptor()
}

func (LoanStatus) Type() protoreflect.EnumType {
	return &file_proto_loan_advance_proto_enumTypes[0]
}

func (x LoanStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoanStatus.Descriptor instead.
func (LoanStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{0}
}

// LoanAdvance request message
type LoanAdvance struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EmployeeId      uint64                 `protobuf:"varint,2,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	Amount          float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Purpose         string                 `protobuf:"bytes,4,opt,name=purpose,proto3" json:"purpose,omitempty"`
	Status          LoanStatus             `protobuf:"varint,5,opt,name=status,proto3,enum=hrms.LoanStatus" json:"status,omitempty"`
	ApprovedBy      *uint64                `protobuf:"varint,6,opt,name=approved_by,json=approvedBy,proto3,oneof" json:"approved_by,omitempty"`
	ApprovalDate    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=approval_date,json=approvalDate,proto3,oneof" json:"approval_date,omitempty"`
	RepaymentStart  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=repayment_start,json=repaymentStart,proto3,oneof" json:"repayment_start,omitempty"`
	RepaymentMonths int32                  `protobuf:"varint,9,opt,name=repayment_months,json=repaymentMonths,proto3" json:"repayment_months,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *LoanAdvance) Reset() {
	*x = LoanAdvance{}
	mi := &file_proto_loan_advance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanAdvance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanAdvance) ProtoMessage() {}

func (x *LoanAdvance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanAdvance.ProtoReflect.Descriptor instead.
func (*LoanAdvance) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{0}
}

func (x *LoanAdvance) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoanAdvance) GetEmployeeId() uint64 {
	if x != nil {
		return x.EmployeeId
	}
	return 0
}

func (x *LoanAdvance) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *LoanAdvance) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

func (x *LoanAdvance) GetStatus() LoanStatus {
	if x != nil {
		return x.Status
	}
	return LoanStatus_LOAN_PENDING
}

func (x *LoanAdvance) GetApprovedBy() uint64 {
	if x != nil && x.ApprovedBy != nil {
		return *x.ApprovedBy
	}
	return 0
}

func (x *LoanAdvance) GetApprovalDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ApprovalDate
	}
	return nil
}

func (x *LoanAdvance) GetRepaymentStart() *timestamppb.Timestamp {
	if x != nil {
		return x.RepaymentStart
	}
	return nil
}

func (x *LoanAdvance) GetRepaymentMonths() int32 {
	if x != nil {
		return x.RepaymentMonths
	}
	return 0
}

func (x *LoanAdvance) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// Request to submit a new loan/advance request
type RequestLoanAdvanceRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	EmployeeId      uint64                 `protobuf:"varint,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	Amount          float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Purpose         string                 `protobuf:"bytes,3,opt,name=purpose,proto3" json:"purpose,omitempty"`
	RepaymentMonths int32                  `protobuf:"varint,4,opt,name=repayment_months,json=repaymentMonths,proto3" json:"repayment_months,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RequestLoanAdvanceRequest) Reset() {
	*x = RequestLoanAdvanceRequest{}
	mi := &file_proto_loan_advance_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestLoanAdvanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLoanAdvanceRequest) ProtoMessage() {}

func (x *RequestLoanAdvanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLoanAdvanceRequest.ProtoReflect.Descriptor instead.
func (*RequestLoanAdvanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{1}
}

func (x *RequestLoanAdvanceRequest) GetEmployeeId() uint64 {
	if x != nil {
		return x.EmployeeId
	}
	return 0
}

func (x *RequestLoanAdvanceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RequestLoanAdvanceRequest) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

func (x *RequestLoanAdvanceRequest) GetRepaymentMonths() int32 {
	if x != nil {
		return x.RepaymentMonths
	}
	return 0
}

// Response containing loan details
type LoanAdvanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Loan          *LoanAdvance           `protobuf:"bytes,1,opt,name=loan,proto3" json:"loan,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoanAdvanceResponse) Reset() {
	*x = LoanAdvanceResponse{}
	mi := &file_proto_loan_advance_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanAdvanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanAdvanceResponse) ProtoMessage() {}

func (x *LoanAdvanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanAdvanceResponse.ProtoReflect.Descriptor instead.
func (*LoanAdvanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{2}
}

func (x *LoanAdvanceResponse) GetLoan() *LoanAdvance {
	if x != nil {
		return x.Loan
	}
	return nil
}

// Request to approve a loan/advance request
type ApproveLoanAdvanceRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	LoanId         uint64                 `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	ApproverId     uint64                 `protobuf:"varint,2,opt,name=approver_id,json=approverId,proto3" json:"approver_id,omitempty"`
	ApprovalDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=approval_date,json=approvalDate,proto3" json:"approval_date,omitempty"`
	RepaymentStart *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=repayment_start,json=repaymentStart,proto3" json:"repayment_start,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ApproveLoanAdvanceRequest) Reset() {
	*x = ApproveLoanAdvanceRequest{}
	mi := &file_proto_loan_advance_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApproveLoanAdvanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveLoanAdvanceRequest) ProtoMessage() {}

func (x *ApproveLoanAdvanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveLoanAdvanceRequest.ProtoReflect.Descriptor instead.
func (*ApproveLoanAdvanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{3}
}

func (x *ApproveLoanAdvanceRequest) GetLoanId() uint64 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

func (x *ApproveLoanAdvanceRequest) GetApproverId() uint64 {
	if x != nil {
		return x.ApproverId
	}
	return 0
}

func (x *ApproveLoanAdvanceRequest) GetApprovalDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ApprovalDate
	}
	return nil
}

func (x *ApproveLoanAdvanceRequest) GetRepaymentStart() *timestamppb.Timestamp {
	if x != nil {
		return x.RepaymentStart
	}
	return nil
}

// Request to reject a loan/advance request
type RejectLoanAdvanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LoanId        uint64                 `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	ApproverId    uint64                 `protobuf:"varint,2,opt,name=approver_id,json=approverId,proto3" json:"approver_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RejectLoanAdvanceRequest) Reset() {
	*x = RejectLoanAdvanceRequest{}
	mi := &file_proto_loan_advance_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectLoanAdvanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectLoanAdvanceRequest) ProtoMessage() {}

func (x *RejectLoanAdvanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectLoanAdvanceRequest.ProtoReflect.Descriptor instead.
func (*RejectLoanAdvanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{4}
}

func (x *RejectLoanAdvanceRequest) GetLoanId() uint64 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

func (x *RejectLoanAdvanceRequest) GetApproverId() uint64 {
	if x != nil {
		return x.ApproverId
	}
	return 0
}

// Request to fetch a loan request by ID
type GetLoanAdvanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LoanId        uint64                 `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLoanAdvanceRequest) Reset() {
	*x = GetLoanAdvanceRequest{}
	mi := &file_proto_loan_advance_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLoanAdvanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoanAdvanceRequest) ProtoMessage() {}

func (x *GetLoanAdvanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoanAdvanceRequest.ProtoReflect.Descriptor instead.
func (*GetLoanAdvanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{5}
}

func (x *GetLoanAdvanceRequest) GetLoanId() uint64 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

// Request to list loan/advance requests with optional filters
type ListLoanAdvancesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *LoanStatus            `protobuf:"varint,1,opt,name=status,proto3,enum=hrms.LoanStatus,oneof" json:"status,omitempty"`
	EmployeeId    *uint64                `protobuf:"varint,2,opt,name=employee_id,json=employeeId,proto3,oneof" json:"employee_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLoanAdvancesRequest) Reset() {
	*x = ListLoanAdvancesRequest{}
	mi := &file_proto_loan_advance_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLoanAdvancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoanAdvancesRequest) ProtoMessage() {}

func (x *ListLoanAdvancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoanAdvancesRequest.ProtoReflect.Descriptor instead.
func (*ListLoanAdvancesRequest) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{6}
}

func (x *ListLoanAdvancesRequest) GetStatus() LoanStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return LoanStatus_LOAN_PENDING
}

func (x *ListLoanAdvancesRequest) GetEmployeeId() uint64 {
	if x != nil && x.EmployeeId != nil {
		return *x.EmployeeId
	}
	return 0
}

// Response containing a list of loan/advance requests
type ListLoanAdvancesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Loans         []*LoanAdvance         `protobuf:"bytes,1,rep,name=loans,proto3" json:"loans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLoanAdvancesResponse) Reset() {
	*x = ListLoanAdvancesResponse{}
	mi := &file_proto_loan_advance_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLoanAdvancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoanAdvancesResponse) ProtoMessage() {}

func (x *ListLoanAdvancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoanAdvancesResponse.ProtoReflect.Descriptor instead.
func (*ListLoanAdvancesResponse) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{7}
}

func (x *ListLoanAdvancesResponse) GetLoans() []*LoanAdvance {
	if x != nil {
		return x.Loans
	}
	return nil
}

// Request to delete a loan/advance request
type DeleteLoanAdvanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LoanId        uint64                 `protobuf:"varint,1,opt,name=loan_id,json=loanId,proto3" json:"loan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteLoanAdvanceRequest) Reset() {
	*x = DeleteLoanAdvanceRequest{}
	mi := &file_proto_loan_advance_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteLoanAdvanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLoanAdvanceRequest) ProtoMessage() {}

func (x *DeleteLoanAdvanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_loan_advance_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLoanAdvanceRequest.ProtoReflect.Descriptor instead.
func (*DeleteLoanAdvanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_loan_advance_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteLoanAdvanceRequest) GetLoanId() uint64 {
	if x != nil {
		return x.LoanId
	}
	return 0
}

var File_proto_loan_advance_proto protoreflect.FileDescriptor

var file_proto_loan_advance_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x72, 0x6d, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec,
	0x03, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f,
	0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x00, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x44, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0e,
	0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x72, 0x65,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x99, 0x01,
	0x0a, 0x19, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x22, 0x3c, 0x0a, 0x13, 0x4c, 0x6f, 0x61,
	0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x6e, 0x22, 0xdb, 0x01, 0x0a, 0x19, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x3f, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x43, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x22, 0x54, 0x0a, 0x18, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4c,
	0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x89, 0x01,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x68, 0x72, 0x6d, 0x73,
	0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52,
	0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x6f, 0x61, 0x6e,
	0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x22, 0x33,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x6f, 0x61,
	0x6e, 0x49, 0x64, 0x2a, 0x44, 0x0a, 0x0a, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4f, 0x41, 0x4e, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x41, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x52,
	0x4f, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x41, 0x4e, 0x5f, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xec, 0x03, 0x0a, 0x12, 0x4c, 0x6f,
	0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x50, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c,
	0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x61,
	0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x4b, 0x0a, 0x11, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x1b, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x68, 0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x68,
	0x72, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x72,
	0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x1e, 0x2e, 0x68, 0x72, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f,
	0x61, 0x6e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x68, 0x72, 0x6d, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x72, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_proto_loan_advance_proto_rawDescOnce sync.Once
	file_proto_loan_advance_proto_rawDescData []byte
)

func file_proto_loan_advance_proto_rawDescGZIP() []byte {
	file_proto_loan_advance_proto_rawDescOnce.Do(func() {
		file_proto_loan_advance_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_loan_advance_proto_rawDesc), len(file_proto_loan_advance_proto_rawDesc)))
	})
	return file_proto_loan_advance_proto_rawDescData
}

var file_proto_loan_advance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_loan_advance_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_loan_advance_proto_goTypes = []any{
	(LoanStatus)(0),                   // 0: hrms.LoanStatus
	(*LoanAdvance)(nil),               // 1: hrms.LoanAdvance
	(*RequestLoanAdvanceRequest)(nil), // 2: hrms.RequestLoanAdvanceRequest
	(*LoanAdvanceResponse)(nil),       // 3: hrms.LoanAdvanceResponse
	(*ApproveLoanAdvanceRequest)(nil), // 4: hrms.ApproveLoanAdvanceRequest
	(*RejectLoanAdvanceRequest)(nil),  // 5: hrms.RejectLoanAdvanceRequest
	(*GetLoanAdvanceRequest)(nil),     // 6: hrms.GetLoanAdvanceRequest
	(*ListLoanAdvancesRequest)(nil),   // 7: hrms.ListLoanAdvancesRequest
	(*ListLoanAdvancesResponse)(nil),  // 8: hrms.ListLoanAdvancesResponse
	(*DeleteLoanAdvanceRequest)(nil),  // 9: hrms.DeleteLoanAdvanceRequest
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 11: google.protobuf.Empty
}
var file_proto_loan_advance_proto_depIdxs = []int32{
	0,  // 0: hrms.LoanAdvance.status:type_name -> hrms.LoanStatus
	10, // 1: hrms.LoanAdvance.approval_date:type_name -> google.protobuf.Timestamp
	10, // 2: hrms.LoanAdvance.repayment_start:type_name -> google.protobuf.Timestamp
	10, // 3: hrms.LoanAdvance.created_at:type_name -> google.protobuf.Timestamp
	1,  // 4: hrms.LoanAdvanceResponse.loan:type_name -> hrms.LoanAdvance
	10, // 5: hrms.ApproveLoanAdvanceRequest.approval_date:type_name -> google.protobuf.Timestamp
	10, // 6: hrms.ApproveLoanAdvanceRequest.repayment_start:type_name -> google.protobuf.Timestamp
	0,  // 7: hrms.ListLoanAdvancesRequest.status:type_name -> hrms.LoanStatus
	1,  // 8: hrms.ListLoanAdvancesResponse.loans:type_name -> hrms.LoanAdvance
	2,  // 9: hrms.LoanAdvanceService.RequestLoanAdvance:input_type -> hrms.RequestLoanAdvanceRequest
	4,  // 10: hrms.LoanAdvanceService.ApproveLoanAdvance:input_type -> hrms.ApproveLoanAdvanceRequest
	5,  // 11: hrms.LoanAdvanceService.RejectLoanAdvance:input_type -> hrms.RejectLoanAdvanceRequest
	6,  // 12: hrms.LoanAdvanceService.GetLoanAdvance:input_type -> hrms.GetLoanAdvanceRequest
	7,  // 13: hrms.LoanAdvanceService.ListLoanAdvances:input_type -> hrms.ListLoanAdvancesRequest
	9,  // 14: hrms.LoanAdvanceService.DeleteLoanAdvance:input_type -> hrms.DeleteLoanAdvanceRequest
	3,  // 15: hrms.LoanAdvanceService.RequestLoanAdvance:output_type -> hrms.LoanAdvanceResponse
	11, // 16: hrms.LoanAdvanceService.ApproveLoanAdvance:output_type -> google.protobuf.Empty
	11, // 17: hrms.LoanAdvanceService.RejectLoanAdvance:output_type -> google.protobuf.Empty
	3,  // 18: hrms.LoanAdvanceService.GetLoanAdvance:output_type -> hrms.LoanAdvanceResponse
	8,  // 19: hrms.LoanAdvanceService.ListLoanAdvances:output_type -> hrms.ListLoanAdvancesResponse
	11, // 20: hrms.LoanAdvanceService.DeleteLoanAdvance:output_type -> google.protobuf.Empty
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_loan_advance_proto_init() }
func file_proto_loan_advance_proto_init() {
	if File_proto_loan_advance_proto != nil {
		return
	}
	file_proto_loan_advance_proto_msgTypes[0].OneofWrappers = []any{}
	file_proto_loan_advance_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_loan_advance_proto_rawDesc), len(file_proto_loan_advance_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_loan_advance_proto_goTypes,
		DependencyIndexes: file_proto_loan_advance_proto_depIdxs,
		EnumInfos:         file_proto_loan_advance_proto_enumTypes,
		MessageInfos:      file_proto_loan_advance_proto_msgTypes,
	}.Build()
	File_proto_loan_advance_proto = out.File
	file_proto_loan_advance_proto_goTypes = nil
	file_proto_loan_advance_proto_depIdxs = nil
}
