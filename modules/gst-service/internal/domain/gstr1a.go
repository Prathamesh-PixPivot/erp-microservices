package domain

// GSTR1ARequest represents the structure of GSTR1A data
type GSTR1ARequest struct {
    ID           uint       `gorm:"primaryKey"`
    GSTIN        string     `json:"gstin"`
    ReturnPeriod string     `json:"return_period"`
    Invoices     []Invoice  `gorm:"many2many:gstr1a_invoices"` // ✅ Many-to-Many relation with Invoice
}

// GSTR1ASubmitRequest represents the request for submitting GSTR1A
type GSTR1ASubmitRequest struct {
	GSTIN        string `json:"gstin"`
	ReturnPeriod string `json:"return_period"`
}

// GSTR1AResponse represents the response for saving GSTR1A data
type GSTR1AResponse struct {
	RefID  string `json:"ref_id"`
	Status string `json:"status"`
}

// GSTR1AStatusRequest represents the request to get the status of GSTR1A
type GSTR1AStatusRequest struct {
	GSTIN        string `json:"gstin"`
	ReturnPeriod string `json:"return_period"`
	ARN          string `json:"arn"`
}

// GSTR1AStatusResponse represents the response for GSTR1A status
type GSTR1AStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// GSTR1AFileRequest represents the request to file GSTR1A
type GSTR1AFileRequest struct {
	GSTIN        string `json:"gstin"`
	ReturnPeriod string `json:"return_period"`
	ARN          string `json:"arn"`
}

// GSTR1AFileResponse represents the response for filing GSTR1A
type GSTR1AFileResponse struct {
	FilingStatus string `json:"filing_status"`
	Message      string `json:"message"`
}
