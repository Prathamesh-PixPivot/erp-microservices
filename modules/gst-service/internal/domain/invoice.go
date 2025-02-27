package domain

type Invoice struct {
	ID            uint   `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	InvoiceNumber string `json:"invoice_number"`  // Invoice number
	InvoiceDate   string `json:"invoice_date"`    // Invoice date (YYYY-MM-DD)
	GSTINSupplier string `json:"gstin_supplier"`  // GSTIN of the supplier
	GSTINReceiver string `json:"gstin_receiver"`  // GSTIN of the receiver
	TaxableValue  string `json:"taxable_value"`   // Taxable value of the invoice
	TaxAmount     string `json:"tax_amount"`      // Tax amount
	HSNCode       string `json:"hsn_code"`        // HSN code for the product
	PlaceOfSupply string `json:"place_of_supply"` // Place of supply (state code)
}
