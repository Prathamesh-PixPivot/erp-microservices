// zabbix_client.go
package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type ZabbixClient struct {
	APIURL    string
	AuthToken string
	Timeout   time.Duration
}

type Request struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth,omitempty"`
	ID      int         `json:"id"`
}

func (c *ZabbixClient) Call(method string, params interface{}) (json.RawMessage, error) {
	reqData := Request{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		Auth:    c.AuthToken,
		ID:      1,
	}
	data, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: c.Timeout}
	var resp *http.Response
	// Simple retry logic: try up to 3 times.
	for i := 0; i < 3; i++ {
		resp, err = client.Post(c.APIURL, "application/json", bytes.NewBuffer(data))
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res struct {
		Result json.RawMessage `json:"result"`
		Error  interface{}     `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, errors.New("zabbix API error")
	}
	return res.Result, nil
}
