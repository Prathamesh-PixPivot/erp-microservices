// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Activity struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description *string        `json:"description,omitempty"`
	Type        string         `json:"type"`
	Status      ActivityStatus `json:"status"`
	DueDate     *string        `json:"dueDate,omitempty"`
	CreatedAt   string         `json:"createdAt"`
	UpdatedAt   string         `json:"updatedAt"`
	Contact     *Contact       `json:"contact"`
	Tasks       []*Task        `json:"tasks"`
}

type AddPaymentDueInput struct {
	InvoiceID string  `json:"invoice_id"`
	AmountDue float64 `json:"amount_due"`
	DueDate   string  `json:"due_date"`
	Status    string  `json:"status"`
}

type Contact struct {
	ID                  string  `json:"id"`
	FirstName           string  `json:"firstName"`
	LastName            string  `json:"lastName"`
	Email               string  `json:"email"`
	Phone               *string `json:"phone,omitempty"`
	Address             *string `json:"address,omitempty"`
	City                *string `json:"city,omitempty"`
	State               *string `json:"state,omitempty"`
	Country             *string `json:"country,omitempty"`
	ZipCode             *string `json:"zipCode,omitempty"`
	Company             *string `json:"company,omitempty"`
	Position            *string `json:"position,omitempty"`
	SocialMediaProfiles *string `json:"socialMediaProfiles,omitempty"`
	Notes               *string `json:"notes,omitempty"`
	CreatedAt           *string `json:"createdAt,omitempty"`
	UpdatedAt           *string `json:"updatedAt,omitempty"`
}

type CreateActivityInput struct {
	Title       string         `json:"title"`
	Description *string        `json:"description,omitempty"`
	Type        string         `json:"type"`
	Status      ActivityStatus `json:"status"`
	DueDate     *string        `json:"dueDate,omitempty"`
	ContactID   string         `json:"contactId"`
}

type CreateContactInput struct {
	FirstName           string  `json:"firstName"`
	LastName            string  `json:"lastName"`
	Email               string  `json:"email"`
	Phone               *string `json:"phone,omitempty"`
	Address             *string `json:"address,omitempty"`
	City                *string `json:"city,omitempty"`
	State               *string `json:"state,omitempty"`
	Country             *string `json:"country,omitempty"`
	ZipCode             *string `json:"zipCode,omitempty"`
	Company             *string `json:"company,omitempty"`
	Position            *string `json:"position,omitempty"`
	SocialMediaProfiles *string `json:"socialMediaProfiles,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

type CreateCreditDebitNoteInput struct {
	Type      string  `json:"type"`
	InvoiceID string  `json:"invoice_id"`
	Amount    float64 `json:"amount"`
	Reason    string  `json:"reason"`
	Date      string  `json:"date"`
}

type CreateInvoiceInput struct {
	Type        string              `json:"type"`
	VendorID    *string             `json:"vendor_id,omitempty"`
	CustomerID  *string             `json:"customer_id,omitempty"`
	Items       []*InvoiceItemInput `json:"items"`
	InvoiceDate string              `json:"invoice_date"`
}

type CreateLeadInput struct {
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Email          string     `json:"email"`
	Phone          *string    `json:"phone,omitempty"`
	Status         LeadStatus `json:"status"`
	AssignedTo     string     `json:"assignedTo"`
	OrganizationID string     `json:"organizationId"`
}

type CreateLedgerEntryInput struct {
	TransactionID   string   `json:"transaction_id"`
	Description     string   `json:"description"`
	Debit           *float64 `json:"debit,omitempty"`
	Credit          *float64 `json:"credit,omitempty"`
	TransactionDate string   `json:"transaction_date"`
}

type CreateOpportunityInput struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Stage       string   `json:"stage"`
	Amount      float64  `json:"amount"`
	CloseDate   string   `json:"closeDate"`
	Probability *float64 `json:"probability,omitempty"`
	LeadID      string   `json:"leadId"`
	AccountID   *string  `json:"accountId,omitempty"`
	OwnerID     string   `json:"ownerId"`
}

type CreateTaskInput struct {
	Title       string       `json:"title"`
	Description *string      `json:"description,omitempty"`
	Status      TaskStatus   `json:"status"`
	Priority    TaskPriority `json:"priority"`
	DueDate     *string      `json:"dueDate,omitempty"`
	ActivityID  string       `json:"activityId"`
}

type CreditDebitNote struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	InvoiceID string  `json:"invoice_id"`
	Amount    float64 `json:"amount"`
	Reason    string  `json:"reason"`
	Date      string  `json:"date"`
}

type Invoice struct {
	ID            string         `json:"id"`
	InvoiceNumber string         `json:"invoice_number"`
	Type          string         `json:"type"`
	VendorID      *string        `json:"vendor_id,omitempty"`
	CustomerID    *string        `json:"customer_id,omitempty"`
	TotalAmount   float64        `json:"total_amount"`
	Cgst          float64        `json:"cgst"`
	Sgst          float64        `json:"sgst"`
	Igst          float64        `json:"igst"`
	Status        string         `json:"status"`
	InvoiceDate   string         `json:"invoice_date"`
	Items         []*InvoiceItem `json:"items"`
}

type InvoiceItem struct {
	ID       string  `json:"id"`
	ItemID   string  `json:"item_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

type InvoiceItemInput struct {
	ItemID   string  `json:"item_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Lead struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	Email        string        `json:"email"`
	Phone        *string       `json:"phone,omitempty"`
	Status       LeadStatus    `json:"status"`
	AssignedTo   string        `json:"assignedTo"`
	Organization *Organization `json:"organization"`
}

type LedgerEntry struct {
	ID              string   `json:"id"`
	TransactionID   string   `json:"transaction_id"`
	Description     string   `json:"description"`
	Debit           *float64 `json:"debit,omitempty"`
	Credit          *float64 `json:"credit,omitempty"`
	Balance         *float64 `json:"balance,omitempty"`
	TransactionDate string   `json:"transaction_date"`
}

type Mutation struct {
}

type Opportunity struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description *string       `json:"description,omitempty"`
	Stage       string        `json:"stage"`
	Amount      float64       `json:"amount"`
	CloseDate   string        `json:"closeDate"`
	Probability *float64      `json:"probability,omitempty"`
	Lead        *Lead         `json:"lead"`
	Account     *Organization `json:"account,omitempty"`
	Owner       *User         `json:"owner"`
}

type Organization struct {
	ID       string `json:"id"`
	GstIn    string `json:"gstIn"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Zipcode  string `json:"zipcode"`
	Website  string `json:"website"`
	Industry string `json:"industry"`
}

type Payment struct {
	ID              string  `json:"id"`
	PurchaseOrderID string  `json:"purchaseOrderId"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	PaymentTerms    *string `json:"paymentTerms,omitempty"`
	PaidAt          *string `json:"paidAt,omitempty"`
}

type PaymentDue struct {
	ID        string  `json:"id"`
	InvoiceID string  `json:"invoice_id"`
	AmountDue float64 `json:"amount_due"`
	DueDate   string  `json:"due_date"`
	Status    string  `json:"status"`
}

type PurchaseOrder struct {
	ID           string  `json:"id"`
	VendorID     string  `json:"vendorId"`
	OrderDetails string  `json:"orderDetails"`
	Status       string  `json:"status"`
	DeliveryDate *string `json:"deliveryDate,omitempty"`
	ReceivedDate *string `json:"receivedDate,omitempty"`
}

type Query struct {
}

type SigninResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupResponse struct {
	Message        string `json:"message"`
	UserID         string `json:"userId"`
	OrganizationID string `json:"organizationId"`
}

type Task struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description *string      `json:"description,omitempty"`
	Status      TaskStatus   `json:"status"`
	Priority    TaskPriority `json:"priority"`
	DueDate     *string      `json:"dueDate,omitempty"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
	Activity    *Activity    `json:"activity"`
}

type UpdateActivityInput struct {
	ID          string          `json:"id"`
	Title       *string         `json:"title,omitempty"`
	Description *string         `json:"description,omitempty"`
	Type        *string         `json:"type,omitempty"`
	Status      *ActivityStatus `json:"status,omitempty"`
	DueDate     *string         `json:"dueDate,omitempty"`
	ContactID   *string         `json:"contactId,omitempty"`
}

type UpdateContactInput struct {
	ID                  string  `json:"id"`
	FirstName           *string `json:"firstName,omitempty"`
	LastName            *string `json:"lastName,omitempty"`
	Email               *string `json:"email,omitempty"`
	Phone               *string `json:"phone,omitempty"`
	Address             *string `json:"address,omitempty"`
	City                *string `json:"city,omitempty"`
	State               *string `json:"state,omitempty"`
	Country             *string `json:"country,omitempty"`
	ZipCode             *string `json:"zipCode,omitempty"`
	Company             *string `json:"company,omitempty"`
	Position            *string `json:"position,omitempty"`
	SocialMediaProfiles *string `json:"socialMediaProfiles,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

type UpdateInvoiceInput struct {
	InvoiceID   string              `json:"invoice_id"`
	Status      *string             `json:"status,omitempty"`
	Items       []*InvoiceItemInput `json:"items,omitempty"`
	InvoiceDate *string             `json:"invoice_date,omitempty"`
}

type UpdateLeadInput struct {
	ID             string      `json:"id"`
	FirstName      *string     `json:"firstName,omitempty"`
	LastName       *string     `json:"lastName,omitempty"`
	Email          *string     `json:"email,omitempty"`
	Phone          *string     `json:"phone,omitempty"`
	Status         *LeadStatus `json:"status,omitempty"`
	AssignedTo     string      `json:"assignedTo"`
	OrganizationID *string     `json:"organizationId,omitempty"`
}

type UpdateOpportunityInput struct {
	ID          string   `json:"id"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Stage       *string  `json:"stage,omitempty"`
	Amount      *float64 `json:"amount,omitempty"`
	CloseDate   *string  `json:"closeDate,omitempty"`
	Probability *float64 `json:"probability,omitempty"`
	LeadID      *string  `json:"leadId,omitempty"`
	AccountID   *string  `json:"accountId,omitempty"`
	OwnerID     *string  `json:"ownerId,omitempty"`
}

type UpdateTaskInput struct {
	ID          string        `json:"id"`
	Title       *string       `json:"title,omitempty"`
	Description *string       `json:"description,omitempty"`
	Status      *TaskStatus   `json:"status,omitempty"`
	Priority    *TaskPriority `json:"priority,omitempty"`
	DueDate     *string       `json:"dueDate,omitempty"`
	ActivityID  *string       `json:"activityId,omitempty"`
}

type User struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Role         string        `json:"role"`
	Organization *Organization `json:"organization"`
}

type Vendor struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Category         string   `json:"category"`
	Service          string   `json:"service"`
	Industry         string   `json:"industry"`
	Gstin            string   `json:"gstin"`
	Certifications   *string  `json:"certifications,omitempty"`
	Licenses         *string  `json:"licenses,omitempty"`
	IsCompliant      *bool    `json:"isCompliant,omitempty"`
	PerformanceScore *float64 `json:"performanceScore,omitempty"`
	RiskAssessment   *string  `json:"riskAssessment,omitempty"`
}

type VendorPerformance struct {
	ID          string   `json:"id"`
	VendorID    string   `json:"vendorId"`
	Score       *float64 `json:"score,omitempty"`
	RiskLevel   *string  `json:"riskLevel,omitempty"`
	EvaluatedAt *string  `json:"evaluatedAt,omitempty"`
}

type ActivitySortField string

const (
	ActivitySortFieldTitle     ActivitySortField = "TITLE"
	ActivitySortFieldDuedate   ActivitySortField = "DUEDATE"
	ActivitySortFieldCreatedat ActivitySortField = "CREATEDAT"
	ActivitySortFieldUpdatedat ActivitySortField = "UPDATEDAT"
)

var AllActivitySortField = []ActivitySortField{
	ActivitySortFieldTitle,
	ActivitySortFieldDuedate,
	ActivitySortFieldCreatedat,
	ActivitySortFieldUpdatedat,
}

func (e ActivitySortField) IsValid() bool {
	switch e {
	case ActivitySortFieldTitle, ActivitySortFieldDuedate, ActivitySortFieldCreatedat, ActivitySortFieldUpdatedat:
		return true
	}
	return false
}

func (e ActivitySortField) String() string {
	return string(e)
}

func (e *ActivitySortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActivitySortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActivitySortField", str)
	}
	return nil
}

func (e ActivitySortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ActivityStatus string

const (
	ActivityStatusPending    ActivityStatus = "Pending"
	ActivityStatusInProgress ActivityStatus = "InProgress"
	ActivityStatusCompleted  ActivityStatus = "Completed"
	ActivityStatusCanceled   ActivityStatus = "Canceled"
	ActivityStatusScheduled  ActivityStatus = "Scheduled"
)

var AllActivityStatus = []ActivityStatus{
	ActivityStatusPending,
	ActivityStatusInProgress,
	ActivityStatusCompleted,
	ActivityStatusCanceled,
	ActivityStatusScheduled,
}

func (e ActivityStatus) IsValid() bool {
	switch e {
	case ActivityStatusPending, ActivityStatusInProgress, ActivityStatusCompleted, ActivityStatusCanceled, ActivityStatusScheduled:
		return true
	}
	return false
}

func (e ActivityStatus) String() string {
	return string(e)
}

func (e *ActivityStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActivityStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActivityStatus", str)
	}
	return nil
}

func (e ActivityStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContactSortField string

const (
	ContactSortFieldFirstname ContactSortField = "FIRSTNAME"
	ContactSortFieldLastname  ContactSortField = "LASTNAME"
	ContactSortFieldEmail     ContactSortField = "EMAIL"
	ContactSortFieldCreatedat ContactSortField = "CREATEDAT"
	ContactSortFieldUpdatedat ContactSortField = "UPDATEDAT"
)

var AllContactSortField = []ContactSortField{
	ContactSortFieldFirstname,
	ContactSortFieldLastname,
	ContactSortFieldEmail,
	ContactSortFieldCreatedat,
	ContactSortFieldUpdatedat,
}

func (e ContactSortField) IsValid() bool {
	switch e {
	case ContactSortFieldFirstname, ContactSortFieldLastname, ContactSortFieldEmail, ContactSortFieldCreatedat, ContactSortFieldUpdatedat:
		return true
	}
	return false
}

func (e ContactSortField) String() string {
	return string(e)
}

func (e *ContactSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContactSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactSortField", str)
	}
	return nil
}

func (e ContactSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeadStatus string

const (
	LeadStatusNew         LeadStatus = "New"
	LeadStatusContacted   LeadStatus = "Contacted"
	LeadStatusQualified   LeadStatus = "Qualified"
	LeadStatusConverted   LeadStatus = "Converted"
	LeadStatusUnqualified LeadStatus = "Unqualified"
)

var AllLeadStatus = []LeadStatus{
	LeadStatusNew,
	LeadStatusContacted,
	LeadStatusQualified,
	LeadStatusConverted,
	LeadStatusUnqualified,
}

func (e LeadStatus) IsValid() bool {
	switch e {
	case LeadStatusNew, LeadStatusContacted, LeadStatusQualified, LeadStatusConverted, LeadStatusUnqualified:
		return true
	}
	return false
}

func (e LeadStatus) String() string {
	return string(e)
}

func (e *LeadStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeadStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeadStatus", str)
	}
	return nil
}

func (e LeadStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OpportunitySortField string

const (
	OpportunitySortFieldName      OpportunitySortField = "NAME"
	OpportunitySortFieldAmount    OpportunitySortField = "AMOUNT"
	OpportunitySortFieldCreatedat OpportunitySortField = "CREATEDAT"
	OpportunitySortFieldUpdatedat OpportunitySortField = "UPDATEDAT"
)

var AllOpportunitySortField = []OpportunitySortField{
	OpportunitySortFieldName,
	OpportunitySortFieldAmount,
	OpportunitySortFieldCreatedat,
	OpportunitySortFieldUpdatedat,
}

func (e OpportunitySortField) IsValid() bool {
	switch e {
	case OpportunitySortFieldName, OpportunitySortFieldAmount, OpportunitySortFieldCreatedat, OpportunitySortFieldUpdatedat:
		return true
	}
	return false
}

func (e OpportunitySortField) String() string {
	return string(e)
}

func (e *OpportunitySortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OpportunitySortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OpportunitySortField", str)
	}
	return nil
}

func (e OpportunitySortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "Low"
	TaskPriorityMedium TaskPriority = "Medium"
	TaskPriorityHigh   TaskPriority = "High"
)

var AllTaskPriority = []TaskPriority{
	TaskPriorityLow,
	TaskPriorityMedium,
	TaskPriorityHigh,
}

func (e TaskPriority) IsValid() bool {
	switch e {
	case TaskPriorityLow, TaskPriorityMedium, TaskPriorityHigh:
		return true
	}
	return false
}

func (e TaskPriority) String() string {
	return string(e)
}

func (e *TaskPriority) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskPriority(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskPriority", str)
	}
	return nil
}

func (e TaskPriority) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskSortField string

const (
	TaskSortFieldTitle     TaskSortField = "TITLE"
	TaskSortFieldDuedate   TaskSortField = "DUEDATE"
	TaskSortFieldCreatedat TaskSortField = "CREATEDAT"
	TaskSortFieldUpdatedat TaskSortField = "UPDATEDAT"
)

var AllTaskSortField = []TaskSortField{
	TaskSortFieldTitle,
	TaskSortFieldDuedate,
	TaskSortFieldCreatedat,
	TaskSortFieldUpdatedat,
}

func (e TaskSortField) IsValid() bool {
	switch e {
	case TaskSortFieldTitle, TaskSortFieldDuedate, TaskSortFieldCreatedat, TaskSortFieldUpdatedat:
		return true
	}
	return false
}

func (e TaskSortField) String() string {
	return string(e)
}

func (e *TaskSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskSortField", str)
	}
	return nil
}

func (e TaskSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "Pending"
	TaskStatusInProgress TaskStatus = "InProgress"
	TaskStatusCompleted  TaskStatus = "Completed"
	TaskStatusCanceled   TaskStatus = "Canceled"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusPending,
	TaskStatusInProgress,
	TaskStatusCompleted,
	TaskStatusCanceled,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusPending, TaskStatusInProgress, TaskStatusCompleted, TaskStatusCanceled:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
