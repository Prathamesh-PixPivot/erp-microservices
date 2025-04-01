// internal/repository/servicerequest_repository.go
package repository

import (
	"errors"
	"sync"
	"time"

	"itsm/internal/models"

	"github.com/google/uuid"
)

type ServiceRequestRepository interface {
	CreateServiceRequest(sr *models.ServiceRequest) (*models.ServiceRequest, error)
	GetServiceRequest(id string) (*models.ServiceRequest, error)
	UpdateServiceRequest(sr *models.ServiceRequest) error
	DeleteServiceRequest(id string) error
	ListServiceRequests() ([]*models.ServiceRequest, error)
}

type inMemoryServiceRequestRepo struct {
	mu       sync.RWMutex
	requests map[string]*models.ServiceRequest
}

func NewServiceRequestRepository() ServiceRequestRepository {
	return &inMemoryServiceRequestRepo{
		requests: make(map[string]*models.ServiceRequest),
	}
}

func (r *inMemoryServiceRequestRepo) CreateServiceRequest(sr *models.ServiceRequest) (*models.ServiceRequest, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	sr.ID = uuid.New().String()
	now := time.Now()
	sr.CreatedAt = now
	sr.UpdatedAt = now
	sr.Status = "open"
	r.requests[sr.ID] = sr
	return sr, nil
}

func (r *inMemoryServiceRequestRepo) GetServiceRequest(id string) (*models.ServiceRequest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sr, ok := r.requests[id]
	if !ok {
		return nil, errors.New("service request not found")
	}
	return sr, nil
}

func (r *inMemoryServiceRequestRepo) UpdateServiceRequest(sr *models.ServiceRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.requests[sr.ID]; !ok {
		return errors.New("service request not found")
	}
	sr.UpdatedAt = time.Now()
	r.requests[sr.ID] = sr
	return nil
}

func (r *inMemoryServiceRequestRepo) DeleteServiceRequest(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.requests[id]; !ok {
		return errors.New("service request not found")
	}
	delete(r.requests, id)
	return nil
}

func (r *inMemoryServiceRequestRepo) ListServiceRequests() ([]*models.ServiceRequest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*models.ServiceRequest, 0, len(r.requests))
	for _, sr := range r.requests {
		list = append(list, sr)
	}
	return list, nil
}
