package handler

import (
	"context"

	"amaa/internal/models"
	"amaa/internal/services"
	pblicense "amaa/internal/transport/grpc/proto"
	pbcommon "amaa/internal/transport/grpc/proto/common"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LicenseHandler struct {
	pblicense.UnimplementedLicenseServiceServer
	licenseService services.LicenseService
	logger         *zap.Logger
}

func NewLicenseHandler(service services.LicenseService, logger *zap.Logger) *LicenseHandler {
	return &LicenseHandler{
		licenseService: service,
		logger:         logger,
	}
}

func (h *LicenseHandler) RegisterLicense(ctx context.Context, req *pblicense.RegisterLicenseRequest) (*pblicense.LicenseResponse, error) {
	h.logger.Info("RegisterLicense request", zap.String("asset_id", req.AssetId))
	license := &models.License{
		// ID:             // Generate an ID if not provided
		ID:             uuid.New().String(),
		AssetID:         req.AssetId,
		LicenseKey:      req.LicenseKey,
		Vendor:          req.Vendor,
		ContractDetails: req.ContractDetails,
	}
	license.ExpiryDate = req.ExpiryDate.AsTime()
	created, err := h.licenseService.RegisterLicense(license)
	if err != nil {
		h.logger.Error("RegisterLicense error", zap.Error(err))
		return nil, err
	}
	return &pblicense.LicenseResponse{
		LicenseId:       created.ID,
		AssetId:         created.AssetID,
		LicenseKey:      created.LicenseKey,
		ExpiryDate:      timestamppb.New(created.ExpiryDate),
		Vendor:          created.Vendor,
		ContractDetails: created.ContractDetails,
	}, nil
}

func (h *LicenseHandler) GetLicense(ctx context.Context, req *pblicense.GetLicenseRequest) (*pblicense.LicenseResponse, error) {
	h.logger.Info("GetLicense request", zap.String("license_id", req.LicenseId))
	license, err := h.licenseService.GetLicense(req.LicenseId)
	if err != nil {
		h.logger.Error("GetLicense error", zap.Error(err))
		return nil, err
	}
	return &pblicense.LicenseResponse{
		LicenseId:       license.ID,
		AssetId:         license.AssetID,
		LicenseKey:      license.LicenseKey,
		ExpiryDate:      timestamppb.New(license.ExpiryDate),
		Vendor:          license.Vendor,
		ContractDetails: license.ContractDetails,
	}, nil
}

func (h *LicenseHandler) UpdateLicense(ctx context.Context, req *pblicense.UpdateLicenseRequest) (*pblicense.LicenseResponse, error) {
	h.logger.Info("UpdateLicense request", zap.String("license_id", req.LicenseId))
	license := &models.License{
		ID:              req.LicenseId,
		LicenseKey:      req.LicenseKey,
		Vendor:          req.Vendor,
		ContractDetails: req.ContractDetails,
	}
	license.ExpiryDate = req.ExpiryDate.AsTime()
	updated, err := h.licenseService.UpdateLicense(license)
	if err != nil {
		h.logger.Error("UpdateLicense error", zap.Error(err))
		return nil, err
	}
	return &pblicense.LicenseResponse{
		LicenseId:       updated.ID,
		AssetId:         updated.AssetID,
		LicenseKey:      updated.LicenseKey,
		ExpiryDate:      timestamppb.New(updated.ExpiryDate),
		Vendor:          updated.Vendor,
		ContractDetails: updated.ContractDetails,
	}, nil
}

func (h *LicenseHandler) DeleteLicense(ctx context.Context, req *pblicense.DeleteLicenseRequest) (*pbcommon.GenericResponse, error) {
	h.logger.Info("DeleteLicense request", zap.String("license_id", req.LicenseId))
	err := h.licenseService.DeleteLicense(req.LicenseId)
	if err != nil {
		h.logger.Error("DeleteLicense error", zap.Error(err))
		return nil, err
	}
	return &pbcommon.GenericResponse{Message: "License deleted successfully"}, nil
}
