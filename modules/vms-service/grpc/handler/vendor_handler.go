package handler

import (
	"context"
	"log"
	"vms-service/grpc/vms_pb"
	"vms-service/internal/models"
	"vms-service/internal/services"

	"github.com/google/uuid"
)

type VendorHandler struct {
	vms_pb.UnimplementedVendorServiceServer
	service *services.VendorService
}

func NewVendorHandler(service *services.VendorService) *VendorHandler {
	return &VendorHandler{service: service}
}

func (h *VendorHandler) CreateVendor(ctx context.Context, req *vms_pb.CreateVendorRequest) (*vms_pb.VendorResponse, error) {
	vendor := &models.Vendor{
		ID:               uuid.New(),
		Name:             req.Vendor.Name,
		Category:         req.Vendor.Category,
		Service:          req.Vendor.Service,
		Industry:         req.Vendor.Industry,
		GSTIN:            req.Vendor.Gstin,
		Certifications:   req.Vendor.Certifications,
		Licenses:         req.Vendor.Licenses,
		IsCompliant:      req.Vendor.IsCompliant,
		PerformanceScore: float64(req.Vendor.PerformanceScore),
		RiskAssessment:   req.Vendor.RiskAssessment,
	}

	if err := h.service.CreateVendor(vendor); err != nil {
		log.Printf("Failed to create vendor: %v", err)
		return nil, err
	}

	return &vms_pb.VendorResponse{Vendor: req.Vendor}, nil
}

func (h *VendorHandler) GetVendorByID(ctx context.Context, req *vms_pb.GetVendorByIDRequest) (*vms_pb.VendorResponse, error) {
	vendor, err := h.service.GetVendorByID(uuid.MustParse(req.Id))
	if err != nil {
		log.Printf("Failed to get vendor: %v", err)
		return nil, err
	}

	return &vms_pb.VendorResponse{
		Vendor: &vms_pb.Vendor{
			Id:               vendor.ID.String(),
			Name:             vendor.Name,
			Category:         vendor.Category,
			Service:          vendor.Service,
			Industry:         vendor.Industry,
			Gstin:            vendor.GSTIN,
			Certifications:   vendor.Certifications,
			Licenses:         vendor.Licenses,
			IsCompliant:      vendor.IsCompliant,
			PerformanceScore: float32(vendor.PerformanceScore),
			RiskAssessment:   vendor.RiskAssessment,
		},
	}, nil
}

func (h *VendorHandler) UpdateVendor(ctx context.Context, req *vms_pb.UpdateVendorRequest) (*vms_pb.VendorResponse, error) {
	vendor := &models.Vendor{
		ID:               uuid.MustParse(req.Vendor.Id),
		Name:             req.Vendor.Name,
		Category:         req.Vendor.Category,
		Service:          req.Vendor.Service,
		Industry:         req.Vendor.Industry,
		GSTIN:            req.Vendor.Gstin,
		Certifications:   req.Vendor.Certifications,
		Licenses:         req.Vendor.Licenses,
		IsCompliant:      req.Vendor.IsCompliant,
		PerformanceScore: float64(req.Vendor.PerformanceScore),
		RiskAssessment:   req.Vendor.RiskAssessment,
	}

	if err := h.service.UpdateVendor(vendor); err != nil {
		log.Printf("Failed to update vendor: %v", err)
		return nil, err
	}

	return &vms_pb.VendorResponse{Vendor: req.Vendor}, nil
}

func (h *VendorHandler) DeleteVendor(ctx context.Context, req *vms_pb.DeleteVendorRequest) (*vms_pb.DeleteVendorResponse, error) {
	if err := h.service.DeleteVendor(uuid.MustParse(req.Id)); err != nil {
		log.Printf("Failed to delete vendor: %v", err)
		return nil, err
	}

	return &vms_pb.DeleteVendorResponse{Message: "Vendor deleted successfully"}, nil
}

func (h *VendorHandler) SearchVendors(ctx context.Context, req *vms_pb.SearchVendorsRequest) (*vms_pb.SearchVendorsResponse, error) {
	vendors, err := h.service.SearchVendors(req.Query)
	if err != nil {
		log.Printf("Failed to search vendors: %v", err)
		return nil, err
	}

	var vendorList []*vms_pb.Vendor
	for _, vendor := range vendors {
		vendorList = append(vendorList, &vms_pb.Vendor{
			Id:               vendor.ID.String(),
			Name:             vendor.Name,
			Category:         vendor.Category,
			Service:          vendor.Service,
			Industry:         vendor.Industry,
			Gstin:            vendor.GSTIN,
			Certifications:   vendor.Certifications,
			Licenses:         vendor.Licenses,
			IsCompliant:      vendor.IsCompliant,
			PerformanceScore: float32(vendor.PerformanceScore),
			RiskAssessment:   vendor.RiskAssessment,
		})
	}

	return &vms_pb.SearchVendorsResponse{Vendors: vendorList}, nil
}
