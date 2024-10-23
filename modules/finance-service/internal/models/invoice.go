package models

import (
	"time"

	"github.com/google/uuid"
)

// Invoice model for Sales, Purchase, Proforma, and Challan Invoices
type Invoice struct {
	ID             uuid.UUID     `gorm:"type:uuid;primary_key;" json:"id"`
	InvoiceNumber  string        `json:"invoice_number"`                    // Auto-generated or custom pattern
	Type           string        `json:"type"`                              // "sales", "proforma", "challan", "purchase"
	VendorID       *uuid.UUID    `json:"vendor_id"`                         // Vendor ID for purchase invoices
	CustomerID     *uuid.UUID    `json:"customer_id"`                       // Customer ID for sales and other invoices
	OrganizationID uuid.UUID     `json:"organization_id"`                   // Organization ID for generating invoice number
	Items          []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"` // Linked items
	TotalAmount    float64       `json:"total_amount"`                      // Calculated automatically based on items
	CGST           float64       `json:"cgst"`                              // Calculated automatically
	SGST           float64       `json:"sgst"`                              // Calculated automatically
	IGST           float64       `json:"igst"`                              // Calculated automatically
	Status         string        `json:"status"`                            // e.g., "paid", "pending", etc.
	InvoiceDate    time.Time     `json:"invoice_date"`                      // Date of the invoice
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

// InvoiceItem model for individual items in an invoice
type InvoiceItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	InvoiceID uuid.UUID `json:"invoice_id"`
	ItemID    string    `json:"item_id"`  // Item identifier (pulled from inventory service)
	Name      string    `json:"name"`     // Name of the item
	Price     float64   `json:"price"`    // Price per unit
	Quantity  int       `json:"quantity"` // Quantity of the item
	Total     float64   `json:"total"`    // Calculated automatically: Price * Quantity
}
