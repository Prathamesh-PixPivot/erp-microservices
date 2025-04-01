package repository

import (
	"errors"
	"time"

	"itsm/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IncidentRepository interface {
	CreateIncident(incident *models.Incident) (*models.Incident, error)
	GetIncident(id string) (*models.Incident, error)
	UpdateIncident(incident *models.Incident) error
	DeleteIncident(id string) error
	ListIncidents() ([]*models.Incident, error)
}

type gormIncidentRepo struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) IncidentRepository {
	db.AutoMigrate(&models.Incident{})
	return &gormIncidentRepo{db: db}
}

func (r *gormIncidentRepo) CreateIncident(incident *models.Incident) (*models.Incident, error) {
	incident.ID = uuid.New().String()
	now := time.Now()
	incident.CreatedAt = now
	incident.UpdatedAt = now
	incident.Status = "open"
	if err := r.db.Create(incident).Error; err != nil {
		return nil, err
	}
	return incident, nil
}

func (r *gormIncidentRepo) GetIncident(id string) (*models.Incident, error) {
	var incident models.Incident
	if err := r.db.First(&incident, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("incident not found")
		}
		return nil, err
	}
	return &incident, nil
}

func (r *gormIncidentRepo) UpdateIncident(incident *models.Incident) error {
	incident.UpdatedAt = time.Now()
	return r.db.Save(incident).Error
}

func (r *gormIncidentRepo) DeleteIncident(id string) error {
	return r.db.Delete(&models.Incident{}, "id = ?", id).Error
}

func (r *gormIncidentRepo) ListIncidents() ([]*models.Incident, error) {
	var incidents []*models.Incident
	if err := r.db.Find(&incidents).Error; err != nil {
		return nil, err
	}
	return incidents, nil
}
