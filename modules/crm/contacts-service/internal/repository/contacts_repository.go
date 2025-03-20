package repository

import (
	"contacts-service/internal/models"
	"errors"
	"log"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

var (
	ErrContactExists   = errors.New("contact with this email already exists")
	ErrContactNotFound = errors.New("contact not found")
)

// ContactRepository defines the contract for contact CRUD operations.
type ContactRepository interface {
	Create(contact *models.Contact) (*models.Contact, error)
	GetByID(id uint) (*models.Contact, error)
	Update(contact *models.Contact) (*models.Contact, error)
	Delete(id uint) error
	List(pageNumber uint, pageSize uint, sortBy string, ascending bool) ([]models.Contact, error)
}

type contactRepository struct {
	db *gorm.DB
}

// NewContactRepository creates a new instance of contactRepository.
func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

// Create inserts a new unified contact (individual or company) into the database.
func (r *contactRepository) Create(contact *models.Contact) (*models.Contact, error) {
	if err := r.db.Create(contact).Error; err != nil {
		if isUniqueConstraintError(err, "contacts_email_key") {
			return nil, ErrContactExists
		}
		return nil, err
	}
	return contact, nil
}

// GetByID retrieves a unified contact by its ID.
func (r *contactRepository) GetByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	if err := r.db.First(&contact, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	return &contact, nil
}

// Update modifies an existing unified contact.
func (r *contactRepository) Update(contact *models.Contact) (*models.Contact, error) {
	result := r.db.Model(&models.Contact{}).Where("id = ?", contact.ID).Updates(contact)
	if result.Error != nil {
		if isUniqueConstraintError(result.Error, "contacts_email_key") {
			return nil, ErrContactExists
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrContactNotFound
	}
	var updatedContact models.Contact
	if err := r.db.First(&updatedContact, "id=?", contact.ID).Error; err != nil {
		log.Printf("failed to reload updated contact")
		return nil, err
	}
	return &updatedContact, nil
}

// Delete removes a unified contact by its ID.
func (r *contactRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrContactNotFound
	}
	return nil
}

// List retrieves contacts with pagination and sorting.
// This function works with the unified contact model that may include additional fields
// like ContactType, CompanyName, and TaxationDetailID.
func (r *contactRepository) List(pageNumber uint, pageSize uint, sortBy string, ascending bool) ([]models.Contact, error) {
	var contacts []models.Contact

	query := r.db.Model(&models.Contact{})

	// Apply sorting if provided.
	if sortBy != "" {
		order := sortBy
		if ascending {
			order += " ASC"
		} else {
			order += " DESC"
		}
		query = query.Order(order)
	}

	// Apply pagination.
	offset := (pageNumber - 1) * pageSize
	query = query.Offset(int(offset)).Limit(int(pageSize))

	if err := query.Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

// isUniqueConstraintError checks for PostgreSQL unique constraint violations.
func isUniqueConstraintError(err error, constraintName string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" && pgErr.ConstraintName == constraintName {
			return true
		}
	}
	return false
}
