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

type PaymentDueHandler struct {
	service services.PaymentDueService
	finance_pb.UnimplementedPaymentServiceServer
}

func NewPaymentDueHandler(service services.PaymentDueService) *PaymentDueHandler {
	return &PaymentDueHandler{service: service}
}

func (h *PaymentDueHandler) AddPaymentDue(ctx context.Context, req *finance_pb.AddPaymentDueRequest) (*finance_pb.PaymentDueResponse, error) {
	paymentDue := &models.PaymentDue{
		ID:        uuid.New(),
		InvoiceID: uuid.MustParse(req.InvoiceId),
		AmountDue: req.AmountDue,
		DueDate:   req.DueDate.AsTime(),
		Status:    "unpaid",
	}

	err := h.service.CreatePaymentDue(paymentDue)
	if err != nil {
		log.Printf("Error adding payment due: %v", err)
		return nil, err
	}

	return &finance_pb.PaymentDueResponse{
		PaymentDueId: paymentDue.ID.String(),
		InvoiceId:    paymentDue.InvoiceID.String(),
		AmountDue:    paymentDue.AmountDue,
		DueDate:      timestamppb.New(paymentDue.DueDate),
		Status:       paymentDue.Status,
	}, nil
}

func (h *PaymentDueHandler) GetPaymentDueByID(ctx context.Context, req *finance_pb.GetPaymentDueByIDRequest) (*finance_pb.PaymentDueResponse, error) {
	paymentDueID, err := uuid.Parse(req.PaymentDueId)
	if err != nil {
		log.Printf("Invalid payment due ID: %v", err)
		return nil, err
	}

	paymentDue, err := h.service.GetPaymentDueByID(paymentDueID)
	if err != nil {
		log.Printf("Error fetching payment due: %v", err)
		return nil, err
	}

	return &finance_pb.PaymentDueResponse{
		PaymentDueId: paymentDue.ID.String(),
		InvoiceId:    paymentDue.InvoiceID.String(),
		AmountDue:    paymentDue.AmountDue,
		DueDate:      &timestamppb.Timestamp{Seconds: paymentDue.DueDate.Unix()},
		Status:       paymentDue.Status,
	}, nil
}

func (h *PaymentDueHandler) MarkPaymentAsPaid(ctx context.Context, req *finance_pb.MarkPaymentAsPaidRequest) (*finance_pb.PaymentDueResponse, error) {
	paymentDueID, err := uuid.Parse(req.PaymentDueId)
	if err != nil {
		log.Printf("Invalid payment due ID: %v", err)
		return nil, err
	}

	paymentDue, err := h.service.GetPaymentDueByID(paymentDueID)
	if err != nil {
		log.Printf("Error fetching payment due: %v", err)
		return nil, err
	}

	// Mark the payment as paid
	paymentDue.Status = "paid"

	err = h.service.UpdatePaymentDue(paymentDue)
	if err != nil {
		log.Printf("Error updating payment status: %v", err)
		return nil, err
	}

	return &finance_pb.PaymentDueResponse{
		PaymentDueId: paymentDue.ID.String(),
		InvoiceId:    paymentDue.InvoiceID.String(),
		AmountDue:    paymentDue.AmountDue,
		DueDate:      &timestamppb.Timestamp{Seconds: paymentDue.DueDate.Unix()},
		Status:       paymentDue.Status,
	}, nil
}

func (h *PaymentDueHandler) ListPaymentDues(ctx context.Context, req *finance_pb.ListPaymentDueRequest) (*finance_pb.ListPaymentDueResponse, error) {
	paymentDues, err := h.service.ListAllPaymentDues()
	if err != nil {
		log.Printf("Error listing payment dues: %v", err)
		return nil, err
	}

	var paymentDueResponses []*finance_pb.PaymentDueResponse
	for _, paymentDue := range paymentDues {
		paymentDueResponses = append(paymentDueResponses, &finance_pb.PaymentDueResponse{
			PaymentDueId: paymentDue.ID.String(),
			InvoiceId:    paymentDue.InvoiceID.String(),
			AmountDue:    paymentDue.AmountDue,
			DueDate:      &timestamppb.Timestamp{Seconds: paymentDue.DueDate.Unix()},
			Status:       paymentDue.Status,
		})
	}

	return &finance_pb.ListPaymentDueResponse{Payments: paymentDueResponses}, nil
}
