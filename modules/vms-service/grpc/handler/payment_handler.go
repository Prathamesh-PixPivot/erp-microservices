package handler

import (
	"context"
	"log"
	"vms-service/grpc/vms_pb"
	"vms-service/internal/models"
	"vms-service/internal/services"

	"github.com/google/uuid"
)

type PaymentHandler struct {
	vms_pb.UnimplementedPaymentServiceServer
	service *services.PaymentService
}

func NewPaymentHandler(service *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) ProcessInvoice(ctx context.Context, req *vms_pb.ProcessInvoiceRequest) (*vms_pb.PaymentResponse, error) {
	payment := &models.Payment{
		ID:              uuid.New(),
		PurchaseOrderID: uuid.MustParse(req.PurchaseOrderId),
		Amount:          float64(req.Amount),
		PaymentTerms:    req.PaymentTerms,
		Status:          "pending",
	}

	if err := h.service.ProcessInvoice(payment); err != nil {
		log.Printf("Failed to process invoice: %v", err)
		return nil, err
	}

	return &vms_pb.PaymentResponse{
		Id:              payment.ID.String(),
		PurchaseOrderId: payment.PurchaseOrderID.String(),
		Amount:          float32(payment.Amount),
		PaymentTerms:    payment.PaymentTerms,
		Status:          payment.Status,
	}, nil
}

func (h *PaymentHandler) UpdatePaymentStatus(ctx context.Context, req *vms_pb.UpdatePaymentStatusRequest) (*vms_pb.PaymentResponse, error) {
	if err := h.service.UpdatePaymentStatus(uuid.MustParse(req.Id), req.Status); err != nil {
		log.Printf("Failed to update payment status: %v", err)
		return nil, err
	}

	payment, err := h.service.GetPaymentByID(uuid.MustParse(req.Id))
	if err != nil {
		log.Printf("Failed to get payment by ID: %v", err)
		return nil, err
	}

	return &vms_pb.PaymentResponse{
		Id:              payment.ID.String(),
		PurchaseOrderId: payment.PurchaseOrderID.String(),
		Amount:          float32(payment.Amount),
		PaymentTerms:    payment.PaymentTerms,
		Status:          payment.Status,
	}, nil
}

func (h *PaymentHandler) GetPaymentByID(ctx context.Context, req *vms_pb.GetPaymentByIDRequest) (*vms_pb.PaymentResponse, error) {
	payment, err := h.service.GetPaymentByID(uuid.MustParse(req.Id))
	if err != nil {
		log.Printf("Failed to get payment by ID: %v", err)
		return nil, err
	}

	return &vms_pb.PaymentResponse{
		Id:              payment.ID.String(),
		PurchaseOrderId: payment.PurchaseOrderID.String(),
		Amount:          float32(payment.Amount),
		PaymentTerms:    payment.PaymentTerms,
		Status:          payment.Status,
	}, nil
}

func (h *PaymentHandler) DeletePayment(ctx context.Context, req *vms_pb.DeletePaymentRequest) (*vms_pb.DeletePaymentResponse, error) {
	if err := h.service.DeletePayment(uuid.MustParse(req.Id)); err != nil {
		log.Printf("Failed to delete payment: %v", err)
		return nil, err
	}

	return &vms_pb.DeletePaymentResponse{Message: "Payment deleted successfully"}, nil
}
