package domain

// GSTR9Request represents the structure of GSTR9 data
type GSTR9Request struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    Turnover     string `json:"total_turnover"`   // Total turnover for the period
    TaxPayable   string `json:"tax_payable"`      // Tax payable for the period
}

// GSTR9Response represents the response for saving GSTR9 data
type GSTR9Response struct {
    RefID  string `json:"ref_id"`
    Status string `json:"status"`
}

// GSTR9SubmitRequest represents the request for submitting GSTR9
type GSTR9SubmitRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
}

// GSTR9SubmitResponse represents the response after submitting GSTR9
type GSTR9SubmitResponse struct {
    ARN    string `json:"arn"`
    Status string `json:"status"`
}

// GSTR9StatusRequest represents the request to get the status of GSTR9
type GSTR9StatusRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR9StatusResponse represents the response for GSTR9 status
type GSTR9StatusResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// GSTR9FileRequest represents the request to file GSTR9
type GSTR9FileRequest struct {
    GSTIN        string `json:"gstin"`
    ReturnPeriod string `json:"return_period"`
    ARN          string `json:"arn"`
}

// GSTR9FileResponse represents the response for filing GSTR9
type GSTR9FileResponse struct {
    FilingStatus string `json:"filing_status"`
    Message      string `json:"message"`
}
