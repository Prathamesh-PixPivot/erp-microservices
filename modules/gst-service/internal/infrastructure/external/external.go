package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gst-service/internal/infrastructure/config"
	"gst-service/internal/infrastructure/logger"
)

// GSTClient handles communication with the GST Portal API
type GSTClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewGSTClient creates a new GST API client
func NewGSTClient(cfg *config.Config) *GSTClient {
	return &GSTClient{
		baseURL: cfg.ExternalAPI.GSTPortalBaseURL,
		apiKey:  cfg.ExternalAPI.APIKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// MakeRequest handles API requests to the GST portal
func (gc *GSTClient) MakeRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", gc.baseURL, endpoint)

	var reqBody []byte
	var err error
	if payload != nil {
		reqBody, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to encode request payload: %w", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gc.apiKey))

	res, err := gc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-OK status: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %w", err)
	}

	return body, nil
}

// SubmitGSTR1 submits GSTR1 data to the GST Portal
func (gc *GSTClient) SubmitGSTR1(gstin string, returnPeriod string, invoices []map[string]interface{}) (string, error) {
	endpoint := "gst/submitGSTR1"
	payload := map[string]interface{}{
		"gstin":         gstin,
		"return_period": returnPeriod,
		"invoices":      invoices,
	}

	response, err := gc.MakeRequest(http.MethodPost, endpoint, payload)
	if err != nil {
		logger.SISLogger.Error("🚨 Failed to submit GSTR1", "error", err)
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if arn, ok := result["arn"].(string); ok {
		return arn, nil
	}
	return "", errors.New("invalid response from GST portal")
}

// Generic function to get GST filing status
func (gc *GSTClient) GetGSTStatus(gstin, returnPeriod, arn, endpoint string) (string, error) {
	url := fmt.Sprintf("%s/%s?gstin=%s&return_period=%s&arn=%s", gc.baseURL, endpoint, gstin, returnPeriod, arn)

	res, err := gc.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch GST status: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned non-OK status: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read API response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if status, ok := result["status"].(string); ok {
		return status, nil
	}
	return "", errors.New("invalid status response from GST portal")
}

// Submit GSTR1A
func (gc *GSTClient) SubmitGSTR1A(gstin, returnPeriod string, invoices []map[string]interface{}) (string, error) {
	return gc.SubmitGSTR1(gstin, returnPeriod, invoices)
}

// Get GSTR1A Status
func (gc *GSTClient) GetGSTR1AStatus(gstin, returnPeriod, arn string) (string, error) {
	return gc.GetGSTStatus(gstin, returnPeriod, arn, "gst/statusGSTR1A")
}

// Submit GSTR2A
func (gc *GSTClient) SubmitGSTR2A(gstin, returnPeriod string, invoices []map[string]interface{}) (string, error) {
	return gc.SubmitGSTR1(gstin, returnPeriod, invoices)
}

// Get GSTR2A Status
func (gc *GSTClient) GetGSTR2AStatus(gstin, returnPeriod, arn string) (string, error) {
	return gc.GetGSTStatus(gstin, returnPeriod, arn, "gst/statusGSTR2A")
}

// Submit GSTR3B
func (gc *GSTClient) SubmitGSTR3B(gstin, returnPeriod string, taxDetails map[string]interface{}) (string, error) {
	endpoint := "gst/submitGSTR3B"
	payload := map[string]interface{}{
		"gstin":         gstin,
		"return_period": returnPeriod,
		"tax_details":   taxDetails,
	}

	return gc.makeRequestAndExtractARN(endpoint, payload)
}

// Get GSTR3B Status
func (gc *GSTClient) GetGSTR3BStatus(gstin, returnPeriod, arn string) (string, error) {
	return gc.GetGSTStatus(gstin, returnPeriod, arn, "gst/statusGSTR3B")
}

// Submit GSTR9
func (gc *GSTClient) SubmitGSTR9(gstin, returnPeriod string, summary map[string]interface{}) (string, error) {
	endpoint := "gst/submitGSTR9"
	payload := map[string]interface{}{
		"gstin":         gstin,
		"return_period": returnPeriod,
		"summary":       summary,
	}

	return gc.makeRequestAndExtractARN(endpoint, payload)
}

// Get GSTR9 Status
func (gc *GSTClient) GetGSTR9Status(gstin, returnPeriod, arn string) (string, error) {
	return gc.GetGSTStatus(gstin, returnPeriod, arn, "gst/statusGSTR9")
}

// Submit GSTR9C
func (gc *GSTClient) SubmitGSTR9C(gstin, returnPeriod string, auditReport map[string]interface{}) (string, error) {
	endpoint := "gst/submitGSTR9C"
	payload := map[string]interface{}{
		"gstin":         gstin,
		"return_period": returnPeriod,
		"audit_report":  auditReport,
	}

	return gc.makeRequestAndExtractARN(endpoint, payload)
}

// Get GSTR9C Status
func (gc *GSTClient) GetGSTR9CStatus(gstin, returnPeriod, arn string) (string, error) {
	return gc.GetGSTStatus(gstin, returnPeriod, arn, "gst/statusGSTR9C")
}

// Reconcile GSTR1 with GSTR2A
func (gc *GSTClient) ReconcileGSTR1WithGSTR2A(gstin, returnPeriod string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/gst/reconcile?gstin=%s&return_period=%s", gc.baseURL, gstin, returnPeriod)

	res, err := gc.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reconciliation data: %w", err)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	return result, nil
}

// Util function to make request and extract ARN
// makeRequestAndExtractARN sends an API request and extracts the ARN from the response
func (gc *GSTClient) makeRequestAndExtractARN(endpoint string, payload interface{}) (string, error) {
	response, err := gc.MakeRequest(http.MethodPost, endpoint, payload)
	if err != nil {
		logger.SISLogger.Error("🚨 Failed to submit request", "endpoint", endpoint, "error", err)
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if arn, ok := result["arn"].(string); ok {
		return arn, nil
	}
	return "", errors.New("invalid response from GST portal")
}
