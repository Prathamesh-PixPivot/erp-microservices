package grpc

import (
	"context"
	"hrms/internal/dto"
	proto "hrms/internal/transport/grpc/proto"

	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateEmployeeDocument handles the creation of an employee document
func (h *HrmsHandler) CreateEmployeeDocument(ctx context.Context, req *proto.CreateEmployeeDocumentRequest) (*proto.EmployeeDocumentResponse, error) {
	var expiryDate *time.Time
	if req.ExpiryDate != nil {
		expiry := req.ExpiryDate.AsTime()
		expiryDate = &expiry
	}

	doc, err := h.HrmsUsecase.CreateEmployeeDocument(ctx, dto.CreateEmployeeDocumentRequest{
		EmployeeID:   uint(req.EmployeeId),
		DocumentName: req.DocumentName,
		DocumentURL:  req.DocumentUrl,
		ExpiryDate:   expiryDate,
	})
	if err != nil {
		return nil, err
	}

	return &proto.EmployeeDocumentResponse{Document: mapToProtoDocument(doc)}, nil
}

// GetEmployeeDocumentByID retrieves a document by its ID
func (h *HrmsHandler) GetEmployeeDocumentByID(ctx context.Context, req *proto.GetEmployeeDocumentByIDRequest) (*proto.EmployeeDocumentResponse, error) {
	doc, err := h.HrmsUsecase.GetEmployeeDocumentByID(ctx, uint(req.DocumentId))
	if err != nil {
		return nil, err
	}
	return &proto.EmployeeDocumentResponse{Document: mapToProtoDocument(doc)}, nil
}

// GetDocumentsByEmployee retrieves all documents for a specific employee
func (h *HrmsHandler) GetDocumentsByEmployee(ctx context.Context, req *proto.GetDocumentsByEmployeeRequest) (*proto.ListEmployeeDocumentsResponse, error) {
	docs, err := h.HrmsUsecase.GetDocumentsByEmployee(ctx, uint(req.EmployeeId))
	if err != nil {
		return nil, err
	}

	protoDocs := make([]*proto.EmployeeDocument, len(docs))
	for i, doc := range docs {
		protoDocs[i] = mapToProtoDocument(&doc)
	}

	return &proto.ListEmployeeDocumentsResponse{Documents: protoDocs}, nil
}

// GetExpiredDocuments retrieves all expired documents
func (h *HrmsHandler) GetExpiredDocuments(ctx context.Context, _ *emptypb.Empty) (*proto.ListEmployeeDocumentsResponse, error) {
	docs, err := h.HrmsUsecase.GetExpiredDocuments(ctx)
	if err != nil {
		return nil, err
	}

	protoDocs := make([]*proto.EmployeeDocument, len(docs))
	for i, doc := range docs {
		protoDocs[i] = mapToProtoDocument(&doc)
	}

	return &proto.ListEmployeeDocumentsResponse{Documents: protoDocs}, nil
}

// UpdateEmployeeDocument updates an employee document
func (h *HrmsHandler) UpdateEmployeeDocument(ctx context.Context, req *proto.UpdateEmployeeDocumentRequest) (*emptypb.Empty, error) {
	var expiryDate *time.Time
	if req.ExpiryDate != nil {
		expiry := req.ExpiryDate.AsTime()
		expiryDate = &expiry
	}

	err := h.HrmsUsecase.UpdateEmployeeDocument(ctx, uint(req.Id), dto.UpdateEmployeeDocumentRequest{
		DocumentName: req.DocumentName,
		DocumentURL:  req.DocumentUrl,
		ExpiryDate:   expiryDate,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// DeleteEmployeeDocument deletes an employee document
func (h *HrmsHandler) DeleteEmployeeDocument(ctx context.Context, req *proto.DeleteEmployeeDocumentRequest) (*emptypb.Empty, error) {
	if err := h.HrmsUsecase.DeleteEmployeeDocument(ctx, uint(req.DocumentId)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// mapToProtoDocument converts an EmployeeDocumentDTO to a proto EmployeeDocument
func mapToProtoDocument(d *dto.EmployeeDocumentDTO) *proto.EmployeeDocument {
	var expiryDate *timestamppb.Timestamp
	if d.ExpiryDate != nil {
		expiryDate = timestamppb.New(*d.ExpiryDate)
	}

	return &proto.EmployeeDocument{
		Id:           uint64(d.ID),
		EmployeeId:   uint64(d.EmployeeID),
		DocumentName: d.DocumentName,
		DocumentUrl:  d.DocumentURL,
		ExpiryDate:   expiryDate,
	}
}