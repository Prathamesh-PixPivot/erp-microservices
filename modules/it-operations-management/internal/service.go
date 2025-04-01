// service.go
package internal

import (
	"io"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (s *Server) Start() error {
	http.HandleFunc("/webhook/alert", s.handleAlert)
	log.Printf("ITOM server listening on %s", s.Addr)
	if s.Config.UseTLS {
		return http.ListenAndServeTLS(s.Addr, s.Config.TLSCertFile, s.Config.TLSKeyFile, nil)
	}
	return http.ListenAndServe(s.Addr, nil)
}

func (s *Server) handleAlert(w http.ResponseWriter, r *http.Request) {
	if err := s.ITOMService.ProcessAlert(r); err != nil {
		http.Error(w, "failed to process alert", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "alert processed"}`))
}

func (s *ITOMService) ProcessAlert(r *http.Request) error {
	_, span := s.Tracer.Start(r.Context(), "ProcessAlert")
	defer span.End()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.Logger.Error("failed to read alert request body", zap.Error(err))
		span.RecordError(err)
		return err
	}

	s.Logger.Info("Received alert", zap.ByteString("payload", body))

	// Validate incoming request (e.g., check a shared secret header)
	secret := r.Header.Get("X-Shared-Secret")
	if secret != "expected-secret" { // compare to secure config
		s.Logger.Error("invalid shared secret")
		return http.ErrNoCookie
	}

	// Process the alert payload
	if err := HandleAlert(s.Logger, body); err != nil {
		s.Logger.Error("failed to process alert", zap.Error(err))
		span.RecordError(err)
		return err
	}

	// Optionally, trigger auto remediation logic or ITSM incident creation.
	s.Logger.Info("Alert processed successfully", zap.String("time", time.Now().Format(time.RFC3339)))
	return nil
}
