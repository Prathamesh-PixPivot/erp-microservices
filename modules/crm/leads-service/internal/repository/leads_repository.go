package repository

import (
	"fmt"
	"leads-service/internal/models"

	"gorm.io/gorm"
)

type LeadRepository interface {
	Create(lead *models.Lead) (*models.Lead, error)
	GetByID(id uint) (*models.Lead, error)
	GetByEmail(email string) (*models.Lead, error) // New method for finding leads by email
	Update(lead *models.Lead) (*models.Lead, error)
	Delete(id uint) error
	GetAll() ([]models.Lead, error)
}

type leadRepository struct {
	db *gorm.DB
}

func NewLeadRepository(db *gorm.DB) LeadRepository {
	return &leadRepository{db}
}

func (r *leadRepository) Create(lead *models.Lead) (*models.Lead, error) {
	if err := r.db.Create(lead).Error; err != nil {
		return nil, err
	}
	fmt.Println("Lead created successfully")
	return lead, nil
}

func (r *leadRepository) GetByID(id uint) (*models.Lead, error) {
	var lead models.Lead
	if err := r.db.First(&lead, id).Error; err != nil {
		return nil, err
	}
	return &lead, nil
}

// New function to get lead by email
func (r *leadRepository) GetByEmail(email string) (*models.Lead, error) {
	var lead models.Lead
	if err := r.db.Where("email = ?", email).First(&lead).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &lead, nil
}

func (r *leadRepository) Update(lead *models.Lead) (*models.Lead, error) {
	// if err := r.db.Save(lead).Error; err != nil {
	// 	return nil, err
	// }
	// return lead, nil

	if err := r.db.Where("id=?", lead.ID).Updates(lead).Error; err != nil {
		print("failed to update")
		return nil, err
	}

	var updatedLead models.Lead

	if err := r.db.First(&updatedLead, "id=?", lead.ID).Error; err != nil {
		print("failed to reload the updated lead")
		return lead, nil
	}

	return &updatedLead, nil
}

func (r *leadRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Lead{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *leadRepository) GetAll() ([]models.Lead, error) {
	var leads []models.Lead
	if err := r.db.Find(&leads).Error; err != nil {
		return nil, err
	}
	return leads, nil
}
