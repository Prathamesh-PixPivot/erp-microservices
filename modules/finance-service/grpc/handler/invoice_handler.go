package handler

import (
	"context"
	"finance-service/grpc/finance_pb"
	"finance-service/internal/models"
	"finance-service/internal/services"
	"log"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InvoiceHandler struct {
	service services.InvoiceService
	finance_pb.UnimplementedInvoiceServiceServer
}

func NewInvoiceHandler(service services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) CreateInvoice(ctx context.Context, req *finance_pb.CreateInvoiceRequest) (*finance_pb.InvoiceResponse, error) {
	// Initialize invoice with required fields
	invoice := &models.Invoice{
		ID:             uuid.New(),
		Type:           req.Invoice.Type,
		TotalAmount:    req.Invoice.TotalAmount,
		CGST:           req.Invoice.Cgst,
		SGST:           req.Invoice.Sgst,
		IGST:           req.Invoice.Igst,
		Status:         req.Invoice.Status,
		InvoiceDate:    req.Invoice.InvoiceDate.AsTime(),
		VendorID:       &req.Invoice.VendorId,
		CustomerID:     &req.Invoice.CustomerId,
		OrganizationID: req.Invoice.OrganizationId, // Ensure this field exists in the proto definition
	}

	// Check if TotalAmount is calculated correctly
	if invoice.TotalAmount == 0 {
		log.Println("Warning: TotalAmount is zero; ensure calculation is correct.")
	}

	// Handle items for the invoice and calculate TotalAmount if needed
	var calculatedTotalAmount float64
	for _, item := range req.Invoice.Items {
		total := item.Price * float64(item.Quantity)
		calculatedTotalAmount += total

		invoice.Items = append(invoice.Items, models.InvoiceItem{
			ID:        uuid.New(),
			InvoiceID: invoice.ID,
			ItemID:    item.ItemId,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  int(item.Quantity),
			Total:     total,
		})
	}

	// Set calculated total amount if it's supposed to be done here
	if invoice.TotalAmount == 0 {
		invoice.TotalAmount = calculatedTotalAmount
	}

	// Calculate taxes if not provided
	if invoice.CGST == 0 && invoice.SGST == 0 && invoice.IGST == 0 {
		invoice.CGST, invoice.SGST, invoice.IGST = h.service.CalculateTaxes(invoice.TotalAmount, invoice.Type)
	}

	// Generate invoice number if it's not provided
	if invoice.InvoiceNumber == "" {
		var err error
		log.Println("Organization ID:", invoice.OrganizationID)
		invoice.InvoiceNumber, err = h.service.GenerateInvoiceNumber(invoice.OrganizationID)
		if err != nil {
			log.Printf("Error generating invoice number: %v", err)
			return nil, err
		}
	}

	// Call service to create the invoice
	err := h.service.CreateInvoice(invoice)
	if err != nil {
		log.Printf("Error creating invoice: %v", err)
		return nil, err
	}

	// Convert and return the response
	return &finance_pb.InvoiceResponse{
		Invoice: convertModelToProtoInvoice(invoice),
	}, nil
}

func (h *InvoiceHandler) GetInvoiceByID(ctx context.Context, req *finance_pb.GetInvoiceByIDRequest) (*finance_pb.InvoiceResponse, error) {
	invoice, err := h.service.GetInvoiceByID(uuid.MustParse(req.InvoiceId))
	if err != nil {
		log.Printf("Error fetching invoice by ID: %v", err)
		return nil, err
	}

	return &finance_pb.InvoiceResponse{
		Invoice: convertModelToProtoInvoice(invoice),
	}, nil
}

func (h *InvoiceHandler) UpdateInvoice(ctx context.Context, req *finance_pb.UpdateInvoiceRequest) (*finance_pb.InvoiceResponse, error) {
	invoiceID, err := uuid.Parse(req.Invoice.Id)
	if err != nil {
		log.Printf("Invalid invoice ID: %v", err)
		return nil, err
	}

	// Retrieve the existing invoice
	invoice, err := h.service.GetInvoiceByID(invoiceID)
	if err != nil {
		log.Printf("Error fetching invoice: %v", err)
		return nil, err
	}

	// Update status if provided
	if req.Invoice.Status != "" {
		invoice.Status = req.Invoice.Status
	}

	// Update invoice date if provided
	if req.Invoice.InvoiceDate != nil {
		invoice.InvoiceDate = req.Invoice.InvoiceDate.AsTime()
	}

	// Update items if provided and recalculate totals and taxes
	if len(req.Invoice.Items) > 0 {
		invoice.Items = []models.InvoiceItem{} // Reset the items

		// Loop through and add new items
		for _, item := range req.Invoice.Items {
			invoice.Items = append(invoice.Items, models.InvoiceItem{
				ID:        uuid.New(),
				InvoiceID: invoice.ID,
				ItemID:    item.ItemId,
				Name:      item.Name,
				Price:     item.Price,
				Quantity:  int(item.Quantity),
				Total:     item.Price * float64(item.Quantity), // Calculate the total for the item
			})
		}

		// Recalculate the total and taxes
		invoice.TotalAmount = h.service.CalculateTotalAmount(invoice.Items)
		invoice.CGST, invoice.SGST, invoice.IGST = h.service.CalculateTaxes(invoice.TotalAmount, invoice.Type)
	}

	// Generate invoice number if it's not provided
	if invoice.InvoiceNumber == "" {
		log.Println("Organization ID:", invoice.OrganizationID)
		invoice.InvoiceNumber, err = h.service.GenerateInvoiceNumber(invoice.OrganizationID)
		if err != nil {
			log.Printf("Error generating invoice number: %v", err)
			return nil, err
		}
	}

	// Save the updated invoice
	err = h.service.UpdateInvoice(invoice)
	if err != nil {
		log.Printf("Error updating invoice: %v", err)
		return nil, err
	}

	return &finance_pb.InvoiceResponse{
		Invoice: convertModelToProtoInvoice(invoice),
	}, nil
}

func (h *InvoiceHandler) DeleteInvoice(ctx context.Context, req *finance_pb.DeleteInvoiceRequest) (*finance_pb.DeleteInvoiceResponse, error) {
	invoiceID, err := uuid.Parse(req.InvoiceId)
	if err != nil {
		log.Printf("Invalid invoice ID: %v", err)
		return nil, err
	}

	// Perform deletion using the service
	err = h.service.DeleteInvoice(invoiceID)
	if err != nil {
		log.Printf("Error deleting invoice: %v", err)
		return nil, err
	}

	// Return the correct DeleteInvoiceResponse as per the proto definition
	return &finance_pb.DeleteInvoiceResponse{
		Message: "Invoice deleted successfully",
	}, nil
}

func (h *InvoiceHandler) ListInvoices(ctx context.Context, req *finance_pb.ListInvoicesRequest) (*finance_pb.ListInvoicesResponse, error) {
	invoices, err := h.service.ListInvoices(int(req.Page), int(req.PageSize))
	if err != nil {
		log.Printf("Error listing invoices: %v", err)
		return nil, err
	}

	var protoInvoices []*finance_pb.Invoice
	for _, invoice := range invoices {
		protoInvoices = append(protoInvoices, convertModelToProtoInvoice(invoice))
	}

	return &finance_pb.ListInvoicesResponse{
		Invoices: protoInvoices,
	}, nil
}

func convertModelToProtoInvoice(invoice *models.Invoice) *finance_pb.Invoice {
	protoInvoice := &finance_pb.Invoice{
		Id:          invoice.ID.String(),
		Type:        invoice.Type,
		VendorId:    *invoice.VendorID,
		CustomerId:  *invoice.CustomerID,
		TotalAmount: invoice.TotalAmount,
		Cgst:        invoice.CGST,
		Sgst:        invoice.SGST,
		Igst:        invoice.IGST,
		Status:      invoice.Status,
		InvoiceDate: timestamppb.New(invoice.InvoiceDate),
		CreatedAt:   timestamppb.New(invoice.CreatedAt),
		UpdatedAt:   timestamppb.New(invoice.UpdatedAt),
	}

	for _, item := range invoice.Items {
		protoInvoice.Items = append(protoInvoice.Items, &finance_pb.InvoiceItem{
			Id:        item.ID.String(),
			InvoiceId: item.InvoiceID.String(),
			ItemId:    item.ItemID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  int32(item.Quantity),
			Total:     item.Total,
		})
	}

	return protoInvoice
}
