package handler

import (
	"context"
	"finance-service/grpc/finance_pb"
	"finance-service/internal/models"
	"finance-service/internal/services"
	"finance-service/utils"
	"log"

	"time"

	"github.com/google/uuid"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type InvoiceHandler struct {
	service services.InvoiceService
	finance_pb.UnimplementedInvoiceServiceServer
}

func NewInvoiceHandler(service services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) CreateInvoice(ctx context.Context, req *finance_pb.CreateInvoiceRequest) (*finance_pb.InvoiceResponse, error) {
	// Check if request is nil
	if req == nil || req.Invoice == nil {
		log.Printf("Error: Received nil request or nil invoice")
	}
	// date distribution
	// Convert timestamps using utils function
	var InvoiceDate, DueDate, DeliveryDate, ChallanDate, AgainstInvoiceDate time.Time

	InvoiceDate, _ = utils.ConvertStringToTime(req.Invoice.InvoiceDate)
	DueDate, _ = utils.ConvertStringToTime(req.Invoice.DueDate)
	DeliveryDate, _ = utils.ConvertStringToTime(req.Invoice.DeliveryDate)
	ChallanDate, _ = utils.ConvertStringToTime(req.Invoice.ChallanDate)
	AgainstInvoiceDate, _ = utils.ConvertStringToTime(req.Invoice.AgainstInvoiceDate)

	log.Printf("Invoice Dates:\n InvoiceDate: %v\n DueDate: %v\n DeliveryDate: %v\n ChallanDate: %v\n AgainstInvoiceDate: %v",
		InvoiceDate,
		DueDate,
		DeliveryDate,
		ChallanDate,
		AgainstInvoiceDate,
	)

	// Initialize invoice with required fields
	invoice := &models.Invoice{
		ID:                   uuid.New(),
		InvoiceNumber:        req.Invoice.InvoiceNumber,
		InvoiceDate:          InvoiceDate,
		Type:                 req.Invoice.Type,
		VendorId:             req.Invoice.VendorId,
		CustomerId:           req.Invoice.CustomerId,
		OrganizationId:       req.Invoice.OrganizationId,
		DueDate:              DueDate,
		DeliveryDate:         DeliveryDate,
		PoNumber:             req.Invoice.PoNumber,
		EwayNumber:           req.Invoice.EwayNumber,
		Status:               req.Invoice.Status,
		PaymentType:          req.Invoice.PaymentType,
		ChequeNumber:         req.Invoice.ChequeNumber,
		ChallanNumber:        req.Invoice.ChallanNumber,
		ChallanDate:          ChallanDate,
		ReverseCharge:        req.Invoice.ReverseCharge,
		LrNumber:             req.Invoice.LrNumber,
		TransporterName:      req.Invoice.TransporterName,
		TransporterId:        req.Invoice.TransporterId,
		VehicleNumber:        req.Invoice.VehicleNumber,
		AgainstInvoiceNumber: req.Invoice.AgainstInvoiceNumber,
		AgainstInvoiceDate:   AgainstInvoiceDate,
		TotalAmount:          req.Invoice.TotalAmount,
		GstRate:              req.Invoice.GstRate,
		CGST:                 req.Invoice.Cgst,
		SGST:                 req.Invoice.Sgst,
		IGST:                 req.Invoice.Igst,
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
			ID:          uuid.New(),
			InvoiceID:   invoice.ID,
			Hsn:         int(item.Hsn),
			Description: item.Description,
			Name:        item.Name,
			Price:       item.Price,
			Quantity:    int(item.Quantity),
			Total:       total,
		})
	}

	// Set calculated total amount if it's supposed to be done here
	if invoice.TotalAmount == 0 {
		invoice.TotalAmount = calculatedTotalAmount
	}

	// // Calculate taxes if not provided
	if invoice.CGST == 0 && invoice.SGST == 0 && invoice.IGST == 0 {
		invoice.CGST, invoice.SGST, invoice.IGST = h.service.CalculateTaxes(invoice.TotalAmount, invoice.Type)
	}

	// Generate invoice number if it's not provided
	if invoice.InvoiceNumber == "" {
		var err error
		invoice.InvoiceNumber, err = h.service.GenerateInvoiceNumber(invoice.OrganizationId)
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
	log.Printf("reached before save")

	// Convert and return the response
	return &finance_pb.InvoiceResponse{
		Invoice: convertModelToProtoInvoice(invoice),
	}, nil
}

func (h *InvoiceHandler) GetInvoiceByID(ctx context.Context, req *finance_pb.GetInvoiceByIDRequest) (*finance_pb.InvoiceResponse, error) {
	invoice, err := h.service.GetInvoiceByID(uuid.MustParse(req.Id))
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
	if req.Invoice.InvoiceDate != "" {
		InvoiceDate, _ := utils.ConvertStringToTime(req.Invoice.InvoiceDate)
		invoice.InvoiceDate = InvoiceDate
	}

	// Update items if provided and recalculate totals and taxes
	if len(req.Invoice.Items) > 0 {
		invoice.Items = []models.InvoiceItem{} // Reset the items

		// Loop through and add new items
		for _, item := range req.Invoice.Items {
			invoice.Items = append(invoice.Items, models.InvoiceItem{
				ID:          uuid.New(),
				InvoiceID:   invoice.ID,
				Name:        item.Name,
				Hsn:         int(item.Hsn),
				Description: item.Description,
				Price:       item.Price,
				Quantity:    int(item.Quantity),
				Total:       item.Price * float64(item.Quantity), // Calculate the total for the item
			})
		}

		// Recalculate the total and taxes
		invoice.TotalAmount = h.service.CalculateTotalAmount(invoice.Items)
		invoice.CGST, invoice.SGST, invoice.IGST = h.service.CalculateTaxes(invoice.TotalAmount, invoice.Type)
	}

	// Generate invoice number if it's not provided
	if invoice.InvoiceNumber == "" {
		log.Println("Organization ID:", invoice.OrganizationId)
		invoice.InvoiceNumber, err = h.service.GenerateInvoiceNumber(invoice.OrganizationId)
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
	invoiceID, err := uuid.Parse(req.Id)
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
		Id:                   invoice.ID.String(),
		InvoiceNumber:        invoice.InvoiceNumber,
		InvoiceDate:          utils.ConvertTimeToString(invoice.InvoiceDate),
		Type:                 invoice.Type,
		VendorId:             invoice.VendorId,
		CustomerId:           invoice.CustomerId,
		OrganizationId:       invoice.OrganizationId,
		DueDate:              utils.ConvertTimeToString(invoice.DueDate),
		DeliveryDate:         utils.ConvertTimeToString(invoice.DeliveryDate),
		PoNumber:             invoice.PoNumber,
		EwayNumber:           invoice.EwayNumber,
		Status:               invoice.Status,
		PaymentType:          invoice.PaymentType,
		ChequeNumber:         invoice.ChequeNumber,
		ChallanNumber:        invoice.ChallanNumber,
		ChallanDate:          utils.ConvertTimeToString(invoice.ChallanDate),
		ReverseCharge:        invoice.ReverseCharge,
		LrNumber:             invoice.LrNumber,
		TransporterName:      invoice.TransporterName,
		TransporterId:        invoice.TransporterId,
		VehicleNumber:        invoice.VehicleNumber,
		AgainstInvoiceNumber: invoice.AgainstInvoiceNumber,
		AgainstInvoiceDate:   utils.ConvertTimeToString(invoice.AgainstInvoiceDate),
		TotalAmount:          invoice.TotalAmount,
		GstRate:              invoice.GstRate,
		Cgst:                 invoice.CGST,
		Sgst:                 invoice.SGST,
		Igst:                 invoice.IGST,
		CreatedAt:           utils.ConvertTimeToString(invoice.CreatedAt),
		UpdatedAt:            utils.ConvertTimeToString(invoice.UpdatedAt),
	}

	// Convert invoice items
	for _, item := range invoice.Items {
		protoInvoice.Items = append(protoInvoice.Items, &finance_pb.InvoiceItem{
			Id:          item.ID.String(),
			InvoiceId:   item.InvoiceID.String(),
			Name:        item.Name,
			Description: item.Description,
			Hsn:         int32(item.Hsn),
			Quantity:    int32(item.Quantity),
			Price:       item.Price,
			Total:       item.Total,
		})
	}

	return protoInvoice
}
