// internal/models/contact.go

package models

// MinimalContact represents essential contact details needed by Activity.
type MinimalContact struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
}
