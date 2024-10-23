package handler

import (
	"context"
	"finance-service/grpc/finance_pb"
	"finance-service/internal/models"
	"finance-service/internal/services"
	"log"

	"github.com/google/uuid"
)

type CreditDebitNoteHandler struct {
	service services.CreditDebitNoteService
	finance_pb.UnimplementedCreditDebitNoteServiceServer
}

func NewCreditDebitNoteHandler(service services.CreditDebitNoteService) *CreditDebitNoteHandler {
	return &CreditDebitNoteHandler{service: service}
}

func (h *CreditDebitNoteHandler) CreateCreditDebitNote(ctx context.Context, req *finance_pb.CreateCreditDebitNoteRequest) (*finance_pb.CreditDebitNoteResponse, error) {
	note := &models.CreditDebitNote{
		ID:        uuid.New(),
		InvoiceID: uuid.MustParse(req.InvoiceId),
		Type:      req.Type, // "credit" or "debit"
		Amount:    req.Amount,
		Reason:    req.Reason,
	}

	err := h.service.CreateCreditDebitNote(note)
	if err != nil {
		log.Printf("Error creating credit/debit note: %v", err)
		return nil, err
	}

	return &finance_pb.CreditDebitNoteResponse{
		NoteId:    note.ID.String(),
		InvoiceId: note.InvoiceID.String(),
		Type:      note.Type,
		Amount:    note.Amount,
		Reason:    note.Reason,
	}, nil
}

func (h *CreditDebitNoteHandler) GetCreditDebitNoteByID(ctx context.Context, req *finance_pb.GetCreditDebitNoteByIDRequest) (*finance_pb.CreditDebitNoteResponse, error) {
	noteID, err := uuid.Parse(req.NoteId)
	if err != nil {
		log.Printf("Invalid note ID: %v", err)
		return nil, err
	}

	note, err := h.service.GetCreditDebitNoteByID(noteID)
	if err != nil {
		log.Printf("Error fetching credit/debit note: %v", err)
		return nil, err
	}

	return &finance_pb.CreditDebitNoteResponse{
		NoteId:    note.ID.String(),
		InvoiceId: note.InvoiceID.String(),
		Type:      note.Type,
		Amount:    note.Amount,
		Reason:    note.Reason,
	}, nil
}

func (h *CreditDebitNoteHandler) ListCreditDebitNotesByInvoiceID(ctx context.Context, req *finance_pb.ListCreditDebitNotesByInvoiceIDRequest) (*finance_pb.ListCreditDebitNotesResponse, error) {
	invoiceID, err := uuid.Parse(req.InvoiceId)
	if err != nil {
		log.Printf("Invalid invoice ID: %v", err)
		return nil, err
	}

	notes, err := h.service.ListCreditDebitNotesByInvoiceID(invoiceID)
	if err != nil {
		log.Printf("Error listing credit/debit notes: %v", err)
		return nil, err
	}

	var noteResponses []*finance_pb.CreditDebitNoteResponse
	for _, note := range notes {
		noteResponses = append(noteResponses, &finance_pb.CreditDebitNoteResponse{
			NoteId:    note.ID.String(),
			InvoiceId: note.InvoiceID.String(),
			Type:      note.Type,
			Amount:    note.Amount,
			Reason:    note.Reason,
		})
	}

	var noteMessages []*finance_pb.CreditDebitNote
	for _, note := range notes {
		noteMessages = append(noteMessages, &finance_pb.CreditDebitNote{
			NoteId:    note.ID.String(),
			InvoiceId: note.InvoiceID.String(),
			Type:      note.Type,
			Amount:    note.Amount,
			Reason:    note.Reason,
		})
	}

	return &finance_pb.ListCreditDebitNotesResponse{Notes: noteMessages}, nil
}

func (h *CreditDebitNoteHandler) ListAllCreditDebitNotes(ctx context.Context, req *finance_pb.ListCreditDebitNotesRequest) (*finance_pb.ListCreditDebitNotesResponse, error) {
	notes, err := h.service.ListAllCreditDebitNotes()
	if err != nil {
		log.Printf("Error listing credit/debit notes: %v", err)
		return nil, err
	}

	var noteResponses []*finance_pb.CreditDebitNote
	for _, note := range notes {
		noteResponses = append(noteResponses, &finance_pb.CreditDebitNote{
			NoteId:    note.ID.String(),
			InvoiceId: note.InvoiceID.String(),
			Type:      note.Type,
			Amount:    note.Amount,
			Reason:    note.Reason,
		})
	}

	return &finance_pb.ListCreditDebitNotesResponse{Notes: noteResponses}, nil
}
