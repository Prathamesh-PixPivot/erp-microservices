package domain

// GSTR3BRequest represents the structure of GSTR3B data
type GSTR3BRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    TaxableValue string `json:"taxable_value"` // Total taxable value
    TaxLiability string `json:"tax_liability"` // Total tax liability
    ITCClaimed   string `json:"itc_claimed"`   // Input Tax Credit claimed
}

// GSTR3BResponse represents the response for saving GSTR3B data
type GSTR3BResponse struct {
    RefID  string `json:"ref_id"`
    Status string `json:"status"`
}

// GSTR3BSubmitRequest represents the request for submitting GSTR3B
type GSTR3BSubmitRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
}

// GSTR3BSubmitResponse represents the response after submitting GSTR3B
type GSTR3BSubmitResponse struct {
    ARN    string `json:"arn"`
    Status string `json:"status"`
}

// GSTR3BStatusRequest represents the request to get the status of GSTR3B
type GSTR3BStatusRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR3BStatusResponse represents the response for GSTR3B status
type GSTR3BStatusResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// GSTR3BFileRequest represents the request to file GSTR3B
type GSTR3BFileRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR3BFileResponse represents the response for filing GSTR3B
type GSTR3BFileResponse struct {
    FilingStatus string `json:"filing_status"`
    Message      string `json:"message"`
}
