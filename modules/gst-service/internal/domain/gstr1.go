package domain

// GSTR1Request represents the structure of GSTR1 data
type GSTR1Request struct {
	ID           uint      `gorm:"primaryKey"`
	GSTIN        string    `json:"gstin"`
	ReturnPeriod string    `json:"return_period"`
	Invoices     []Invoice `gorm:"many2many:gstr1_invoices"` // ✅ Many-to-Many relation with Invoice
}

// GSTR1SubmitRequest represents the request for submitting GSTR1
type GSTR1SubmitRequest struct {
	GSTIN        string `json:"gstin"`
	ReturnPeriod string `json:"return_period"`
}

// GSTR1StatusRequest represents the request to get the status of GSTR1
type GSTR1StatusRequest struct {
	GSTIN        string `json:"gstin"`
	ReturnPeriod string `json:"return_period"`
	ARN          string `json:"arn"` // ARN (Application Reference Number)
}

// GSTR1FileRequest represents the request to file GSTR1
type GSTR1FileRequest struct {
	GSTIN        string `json:"gstin"`
	ReturnPeriod string `json:"return_period"`
	ARN          string `json:"arn"`
}

// GSTR1Response represents the response after saving or submitting GSTR1
type GSTR1Response struct {
	RefID  string `json:"ref_id"`
	Status string `json:"status"`
}

// GSTR1StatusResponse represents the response for the GSTR1 status
type GSTR1StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// GSTR1ReconcileRequest represents the request for reconciling GSTR1 with GSTR2A
type GSTR1ReconcileRequest struct {
	GSTIN        string    `json:"gstin"`
	ReturnPeriod string    `json:"return_period"`
	Invoices     []Invoice `json:"invoices"`
}

// GSTR1ReconcileResponse represents the response for reconciliation results
type GSTR1ReconcileResponse struct {
	Status                string    `json:"status"`
	ReconciliationDetails []Invoice `json:"reconciliation_details"`
}
