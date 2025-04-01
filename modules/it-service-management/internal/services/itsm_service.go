// internal/service/itsm_service.go
package services

import (
	"itsm/internal/models"
	"itsm/internal/repository"
)

type ITSMService interface {
	// Incident operations
	CreateIncident(incident *models.Incident) (*models.Incident, error)
	GetIncident(id string) (*models.Incident, error)
	UpdateIncident(incident *models.Incident) error
	ListIncidents() ([]*models.Incident, error)
	// Change operations
	CreateChange(change *models.Change) (*models.Change, error)
	GetChange(id string) (*models.Change, error)
	UpdateChange(change *models.Change) error
	ListChanges() ([]*models.Change, error)
	// Service Request operations
	CreateServiceRequest(sr *models.ServiceRequest) (*models.ServiceRequest, error)
	GetServiceRequest(id string) (*models.ServiceRequest, error)
	UpdateServiceRequest(sr *models.ServiceRequest) error
	ListServiceRequests() ([]*models.ServiceRequest, error)
}

type itsmService struct {
	incidentRepo       repository.IncidentRepository
	changeRepo         repository.ChangeRepository
	serviceRequestRepo repository.ServiceRequestRepository
}

func NewITSMService(incidentRepo repository.IncidentRepository, changeRepo repository.ChangeRepository, srRepo repository.ServiceRequestRepository) ITSMService {
	return &itsmService{
		incidentRepo:       incidentRepo,
		changeRepo:         changeRepo,
		serviceRequestRepo: srRepo,
	}
}

// Incident methods
func (s *itsmService) CreateIncident(incident *models.Incident) (*models.Incident, error) {
	return s.incidentRepo.CreateIncident(incident)
}

func (s *itsmService) GetIncident(id string) (*models.Incident, error) {
	return s.incidentRepo.GetIncident(id)
}

func (s *itsmService) UpdateIncident(incident *models.Incident) error {
	return s.incidentRepo.UpdateIncident(incident)
}

func (s *itsmService) ListIncidents() ([]*models.Incident, error) {
	return s.incidentRepo.ListIncidents()
}

// Change methods
func (s *itsmService) CreateChange(change *models.Change) (*models.Change, error) {
	return s.changeRepo.CreateChange(change)
}

func (s *itsmService) GetChange(id string) (*models.Change, error) {
	return s.changeRepo.GetChange(id)
}

func (s *itsmService) UpdateChange(change *models.Change) error {
	return s.changeRepo.UpdateChange(change)
}

func (s *itsmService) ListChanges() ([]*models.Change, error) {
	return s.changeRepo.ListChanges()
}

// Service Request methods
func (s *itsmService) CreateServiceRequest(sr *models.ServiceRequest) (*models.ServiceRequest, error) {
	return s.serviceRequestRepo.CreateServiceRequest(sr)
}

func (s *itsmService) GetServiceRequest(id string) (*models.ServiceRequest, error) {
	return s.serviceRequestRepo.GetServiceRequest(id)
}

func (s *itsmService) UpdateServiceRequest(sr *models.ServiceRequest) error {
	return s.serviceRequestRepo.UpdateServiceRequest(sr)
}

func (s *itsmService) ListServiceRequests() ([]*models.ServiceRequest, error) {
	return s.serviceRequestRepo.ListServiceRequests()
}
