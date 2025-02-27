package domain

// GSTR9CRequest represents the structure of GSTR9C data
type GSTR9CRequest struct {
    GSTIN             string `json:"gstin"`
    ReturnPeriod      string `json:"return_period"`
    AuditDetails      string `json:"audit_details"`       // Auditor’s details
    ReconciliationStmt string `json:"reconciliation_statement"` // Reconciliation statement
}

// GSTR9CResponse represents the response for saving GSTR9C data
type GSTR9CResponse struct {
    RefID  string `json:"ref_id"`
    Status string `json:"status"`
}

// GSTR9CSubmitRequest represents the request for submitting GSTR9C
type GSTR9CSubmitRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
}

// GSTR9CSubmitResponse represents the response after submitting GSTR9C
type GSTR9CSubmitResponse struct {
    ARN    string `json:"arn"`
    Status string `json:"status"`
}

// GSTR9CStatusRequest represents the request to get the status of GSTR9C
type GSTR9CStatusRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR9CStatusResponse represents the response for GSTR9C status
type GSTR9CStatusResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// GSTR9CFileRequest represents the request to file GSTR9C
type GSTR9CFileRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR9CFileResponse represents the response for filing GSTR9C
type GSTR9CFileResponse struct {
    FilingStatus string `json:"filing_status"`
    Message      string `json:"message"`
}
