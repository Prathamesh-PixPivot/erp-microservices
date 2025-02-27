package domain

// GSTR2ARequest represents the structure of GSTR2A data
type GSTR2ARequest struct {
    ID           uint       `gorm:"primaryKey"`
    GSTIN        string     `json:"gstin"`
    ReturnPeriod string     `json:"return_period"`
    Invoices     []Invoice  `gorm:"many2many:gstr2a_invoices"` // ✅ Many-to-Many relation with Invoice
}

// GSTR2ASubmitRequest represents the request for submitting GSTR2A
type GSTR2ASubmitRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
}

// GSTR2AResponse represents the response for saving GSTR2A data
type GSTR2AResponse struct {
    RefID  string `json:"ref_id"`
    Status string `json:"status"`
}

// GSTR2AStatusRequest represents the request to get the status of GSTR2A
type GSTR2AStatusRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR2AStatusResponse represents the response for GSTR2A status
type GSTR2AStatusResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// GSTR2AFileRequest represents the request to file GSTR2A
type GSTR2AFileRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR2AFileResponse represents the response for filing GSTR2A
type GSTR2AFileResponse struct {
    FilingStatus string `json:"filing_status"`
    Message      string `json:"message"`
}
