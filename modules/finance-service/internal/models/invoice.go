package models

import (
	"time"

	"github.com/google/uuid"
)

// Invoice model for Sales, Purchase, Proforma, and Challan Invoices
type Invoice struct {

	ID uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`

	InvoiceNumber  string        `json:"invoice_number"`
	InvoiceDate    time.Time     `json:"invoice_date"`                      // Date of the invoice
	Type           string        `json:"type"`                              // "sales", "proforma", "challan", "purchase"
	VendorId       *string       `json:"vendor_id"`                         // Vendor ID for purchase invoices
	CustomerId     *string       `json:"customer_id"`                       // Customer ID for sales and other invoices
	OrganizationId string        `json:"organization_id"`                   // Organization ID for generating invoice number
	Items          []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"` // Linked items

	DueDate              time.Time `json:"due_date"`
	DeliveryDate         time.Time `json:"delivery_date"`
	PoNumber             string    `json:"po_number"`
	EwayNumber           string    `json:"eway_number"`
	Status               string    `json:"status"`       // e.g., "paid", "pending", etc.
	PaymentType          string    `json:"payment_type"` // e.g., "paid", "pending", etc.
	ChequeNumber         string    `json:"cheque_number"`    // e.g., "paid", "pending", etc.
	ChallanNumber        string    `json:"challan_number"`
	ChallanDate          time.Time `json:"challan_date"`
	ReverseCharge        string    `json:"reverse_charge"`
	LrNumber             string    `json:"lr_number"`
	TransporterName      string    `json:"transporter_name"`
	TransporterId        string    `json:"transporter_id"`
	VehicleNumber        string    `json:"vehicle_number"`
	AgainstInvoiceNumber string    `json:"against_invoice_number"`
	AgainstInvoiceDate   time.Time `json:"against_invoice_date"`

	TotalAmount float64   `json:"total_amount"` // Calculated automatically based on items
	GstRate     float32   `json:"gst_rate"`     // Calculated automatically based on items
	CGST        float64   `json:"cgst"`         // Calculated automatically
	SGST        float64   `json:"sgst"`         // Calculated automatically
	IGST        float64   `json:"igst"`         // Calculated automatically
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// InvoiceItem model for individual items in an invoice
type InvoiceItem struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	InvoiceID   uuid.UUID `json:"invoice_id"`
	Name        string    `json:"name"` // Name of the item
	Description string    `json:"description"`
	Hsn         int       `json:"hsn"`
	Quantity    int       `json:"quantity"` // Quantity of the item
	Price       float64   `json:"price"`    // Price per unit
	Total       float64   `json:"total"`    // Calculated automatically: Price * Quantity
}
