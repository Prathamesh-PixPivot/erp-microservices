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

type LedgerHandler struct {
	service services.LedgerService
	finance_pb.UnimplementedLedgerServiceServer
}

func NewLedgerHandler(service services.LedgerService) *LedgerHandler {
	return &LedgerHandler{service: service}
}

func (h *LedgerHandler) AddLedgerEntry(ctx context.Context, req *finance_pb.AddLedgerEntryRequest) (*finance_pb.LedgerResponse, error) {
	ledgerEntry := &models.LedgerEntry{
		ID:              uuid.New(),
		TransactionID:   req.TransactionId,
		Description:     req.Description,
		Debit:           req.Debit,
		Credit:          req.Credit,
		Balance:         req.Balance,
		TransactionDate: req.TransactionDate.AsTime(),
	}

	err := h.service.CreateLedgerEntry(ledgerEntry)
	if err != nil {
		log.Printf("Error adding ledger entry: %v", err)
		return nil, err
	}

	return &finance_pb.LedgerResponse{
		EntryId:         ledgerEntry.ID.String(),
		TransactionId:   ledgerEntry.TransactionID,
		Description:     ledgerEntry.Description,
		Debit:           ledgerEntry.Debit,
		Credit:          ledgerEntry.Credit,
		Balance:         ledgerEntry.Balance,
		TransactionDate: timestamppb.New(ledgerEntry.TransactionDate),
	}, nil
}

func (h *LedgerHandler) GetLedgerEntryByID(ctx context.Context, req *finance_pb.GetLedgerEntryByIDRequest) (*finance_pb.LedgerResponse, error) {
	ledgerID, err := uuid.Parse(req.EntryId)
	if err != nil {
		log.Printf("Invalid ledger entry ID: %v", err)
		return nil, err
	}

	ledgerEntry, err := h.service.GetLedgerEntryByID(ledgerID)
	if err != nil {
		log.Printf("Error fetching ledger entry: %v", err)
		return nil, err
	}

	return &finance_pb.LedgerResponse{
		EntryId:         ledgerEntry.ID.String(),
		TransactionId:   ledgerEntry.TransactionID,
		Description:     ledgerEntry.Description,
		Debit:           ledgerEntry.Debit,
		Credit:          ledgerEntry.Credit,
		Balance:         ledgerEntry.Balance,
		TransactionDate: timestamppb.New(ledgerEntry.TransactionDate),
	}, nil
}

func (h *LedgerHandler) ListLedgerEntries(ctx context.Context, req *finance_pb.ListLedgerEntriesRequest) (*finance_pb.ListLedgerEntriesResponse, error) {
	entries, err := h.service.ListAllLedgerEntries()
	if err != nil {
		log.Printf("Error listing ledger entries: %v", err)
		return nil, err
	}

	var ledgerResponses []*finance_pb.LedgerResponse
	for _, entry := range entries {
		ledgerResponses = append(ledgerResponses, &finance_pb.LedgerResponse{
			EntryId:         entry.ID.String(),
			TransactionId:   entry.TransactionID,
			Description:     entry.Description,
			Debit:           entry.Debit,
			Credit:          entry.Credit,
			Balance:         entry.Balance,
			TransactionDate: timestamppb.New(entry.TransactionDate),
		})
	}

	return &finance_pb.ListLedgerEntriesResponse{LedgerEntries: ledgerResponses}, nil
}
