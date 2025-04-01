package internal

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Alert represents a simplified alert structure.
type Alert struct {
	HostID      string `json:"hostid"`
	AlertType   string `json:"alert_type"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

// HandleAlert unmarshals and processes the alert payload.
func HandleAlert(logger *zap.Logger, payload []byte) error {
	var alert Alert
	if err := json.Unmarshal(payload, &alert); err != nil {
		logger.Error("failed to unmarshal alert", zap.Error(err))
		return err
	}
	logger.Info("Processed alert", zap.String("host_id", alert.HostID), zap.String("alert_type", alert.AlertType))
	// TODO: Map alert to auto remediation or ITSM incident creation.
	return nil
}
