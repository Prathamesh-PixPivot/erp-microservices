package repository

import (
	"user-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository with a database connection
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser stores a new user in the database
func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

// GetUserByID fetches a user by ID from the database
func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, id).Error
	return &user, err
}

// GetUsersByOrganizationID fetches all users belonging to an organization
func (repo *UserRepository) GetUsersByOrganizationID(orgID uint) ([]models.User, error) {
	var users []models.User
	err := repo.db.Where("organization_id = ?", orgID).Find(&users).Error
	return users, err
}
