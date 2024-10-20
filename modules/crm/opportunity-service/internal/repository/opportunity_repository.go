package repository

import (
	"opportunity-service/internal/models"
	"time"

	"gorm.io/gorm"
)

type OpportunityRepository interface {
	Create(opportunity *models.Opportunity) (*models.Opportunity, error)
	GetByID(id uint) (*models.Opportunity, error)
	Update(opportunity *models.Opportunity) (*models.Opportunity, error)
	Delete(id uint) error
	List(ownerID uint) ([]models.Opportunity, error)
	UpdateSelective(opportunity *models.Opportunity) error
}

type opportunityRepository struct {
	db *gorm.DB
}

func NewOpportunityRepository(db *gorm.DB) OpportunityRepository {
	return &opportunityRepository{db}
}

func (r *opportunityRepository) Create(opportunity *models.Opportunity) (*models.Opportunity, error) {
	if err := r.db.Create(opportunity).Error; err != nil {
		return nil, err
	}
	return opportunity, nil
}

func (r *opportunityRepository) GetByID(id uint) (*models.Opportunity, error) {
	var opportunity models.Opportunity
	if err := r.db.First(&opportunity, id).Error; err != nil {
		return nil, err
	}
	return &opportunity, nil
}

func (r *opportunityRepository) Update(opportunity *models.Opportunity) (*models.Opportunity, error) {
	if err := r.db.Save(opportunity).Error; err != nil {
		return nil, err
	}
	return opportunity, nil
}

func (r *opportunityRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Opportunity{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *opportunityRepository) List(ownerID uint) ([]models.Opportunity, error) {
	var opportunities []models.Opportunity
	query := r.db
	if ownerID != 0 {
		query = query.Where("owner_id = ?", ownerID)
	}
	if err := query.Find(&opportunities).Error; err != nil {
		return nil, err
	}
	return opportunities, nil
}

func (r *opportunityRepository) UpdateSelective(opportunity *models.Opportunity) error {
	// Create a map of fields to update
	updates := map[string]interface{}{}

	if opportunity.Name != nil {
		updates["name"] = opportunity.Name
	}
	if opportunity.Description != nil {
		updates["description"] = opportunity.Description
	}
	if opportunity.Stage != nil {
		updates["stage"] = opportunity.Stage
	}
	if opportunity.Amount != nil {
		updates["amount"] = opportunity.Amount
	}
	if opportunity.CloseDate != nil {
		updates["close_date"] = opportunity.CloseDate
	}
	if opportunity.Probability != nil {
		updates["probability"] = opportunity.Probability
	}
	if opportunity.LeadID != nil {
		updates["lead_id"] = opportunity.LeadID
	}
	if opportunity.AccountID != nil {
		updates["account_id"] = opportunity.AccountID
	}
	if opportunity.OwnerID != nil {
		updates["owner_id"] = opportunity.OwnerID
	}

	updates["updated_at"] = time.Now()

	return r.db.Model(&models.Opportunity{}).Where("id = ?", opportunity.ID).Updates(updates).Error
}
