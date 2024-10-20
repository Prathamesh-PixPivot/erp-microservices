package services

import (
	"contacts-service/internal/models"
	"contacts-service/internal/repository"
	"errors"
	"regexp"
	"time"
)

var (
	ErrContactNotFound    = errors.New("contact not found")
	ErrInvalidContactData = errors.New("invalid contact data")
	ErrContactExists      = errors.New("contact with this email already exists")
)

type ContactService interface {
	CreateContact(contact *models.Contact) (*models.Contact, error)
	GetContact(id uint) (*models.Contact, error)
	UpdateContact(contact *models.Contact) (*models.Contact, error)
	DeleteContact(id uint) error
	ListContacts(pageNumber uint, pageSize uint, sortBy string, ascending bool) ([]models.Contact, error)
}

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) ContactService {
	return &contactService{repo: repo}
}

// CreateContact validates and creates a new contact.
func (s *contactService) CreateContact(contact *models.Contact) (*models.Contact, error) {
	// Validate required fields
	if contact.FirstName == "" || contact.LastName == "" || contact.Email == "" {
		return nil, ErrInvalidContactData
	}

	// Validate email format
	if !isValidEmail(contact.Email) {
		return nil, errors.New("invalid email format")
	}

	// Set timestamps
	now := time.Now()
	contact.CreatedAt = now
	contact.UpdatedAt = now

	// Attempt to create the contact
	createdContact, err := s.repo.Create(contact)
	if err != nil {
		if errors.Is(err, repository.ErrContactExists) {
			return nil, ErrContactExists
		}
		return nil, err
	}

	return createdContact, nil
}

// GetContact retrieves a contact by ID.
func (s *contactService) GetContact(id uint) (*models.Contact, error) {
	contact, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrContactNotFound) {
			return nil, ErrContactNotFound
		}
		return nil, err
	}
	return contact, nil
}

// UpdateContact validates and updates an existing contact.
func (s *contactService) UpdateContact(contact *models.Contact) (*models.Contact, error) {
	// Validate contact ID
	if contact.ID == 0 {
		return nil, ErrInvalidContactData
	}

	// Validate email format if provided
	if contact.Email != "" && !isValidEmail(contact.Email) {
		return nil, errors.New("invalid email format")
	}

	// Set the updated_at timestamp
	contact.UpdatedAt = time.Now()

	// Update the contact
	updatedContact, err := s.repo.Update(contact)
	if err != nil {
		if errors.Is(err, repository.ErrContactNotFound) {
			return nil, ErrContactNotFound
		}
		if errors.Is(err, repository.ErrContactExists) {
			return nil, ErrContactExists
		}
		return nil, err
	}

	return updatedContact, nil
}

// DeleteContact removes a contact by ID.
func (s *contactService) DeleteContact(id uint) error {
	// Check if the contact exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrContactNotFound) {
			return ErrContactNotFound
		}
		return err
	}

	// Delete the contact
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

// ListContacts retrieves contacts with pagination and sorting.
func (s *contactService) ListContacts(pageNumber uint, pageSize uint, sortBy string, ascending bool) ([]models.Contact, error) {
	// Validate pagination parameters
	if pageNumber == 0 {
		pageNumber = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	// Validate sortBy field
	validSortFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"created_at": true,
		"updated_at": true,
	}
	if sortBy != "" && !validSortFields[sortBy] {
		return nil, errors.New("invalid sort field")
	}

	contacts, err := s.repo.List(pageNumber, pageSize, sortBy, ascending)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

// Helper function to validate email format using regex.
func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
