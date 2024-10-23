package handler

import (
	"context"
	"log"
	"vms-service/grpc/vms_pb"
	"vms-service/internal/models"
	"vms-service/internal/services"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PurchaseOrderHandler struct {
	vms_pb.UnimplementedPurchaseOrderServiceServer
	service *services.PurchaseOrderService
}

func NewPurchaseOrderHandler(service *services.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{service: service}
}

func (h *PurchaseOrderHandler) CreatePurchaseOrder(ctx context.Context, req *vms_pb.CreatePurchaseOrderRequest) (*vms_pb.PurchaseOrderResponse, error) {
	po := &models.PurchaseOrder{
		ID:           uuid.New(),
		VendorID:     uuid.MustParse(req.VendorId),
		OrderDetails: req.OrderDetails,
		Status:       "pending",
		DeliveryDate: req.DeliveryDate.AsTime(),
	}

	if err := h.service.CreatePurchaseOrder(po); err != nil {
		log.Printf("Failed to create purchase order: %v", err)
		return nil, err
	}

	return &vms_pb.PurchaseOrderResponse{
		Id:           po.ID.String(),
		VendorId:     po.VendorID.String(),
		OrderDetails: po.OrderDetails,
		Status:       po.Status,
		DeliveryDate: timestamppb.New(po.DeliveryDate).String(),
	}, nil
}

func (h *PurchaseOrderHandler) GetPurchaseOrderByID(ctx context.Context, req *vms_pb.GetPurchaseOrderByIDRequest) (*vms_pb.PurchaseOrderResponse, error) {
	po, err := h.service.GetPurchaseOrderByID(uuid.MustParse(req.Id))
	if err != nil {
		log.Printf("Failed to get purchase order: %v", err)
		return nil, err
	}

	return &vms_pb.PurchaseOrderResponse{
		Id:           po.ID.String(),
		VendorId:     po.VendorID.String(),
		OrderDetails: po.OrderDetails,
		Status:       po.Status,
		DeliveryDate: timestamppb.New(po.DeliveryDate).String(),
		ReceivedDate: timestamppb.New(po.ReceivedDate).String(),
	}, nil
}

func (h *PurchaseOrderHandler) UpdatePurchaseOrder(ctx context.Context, req *vms_pb.UpdatePurchaseOrderRequest) (*vms_pb.PurchaseOrderResponse, error) {
	po := &models.PurchaseOrder{
		ID:           uuid.MustParse(req.PurchaseOrder.Id),
		VendorID:     uuid.MustParse(req.PurchaseOrder.VendorId),
		OrderDetails: req.PurchaseOrder.OrderDetails,
		Status:       req.PurchaseOrder.Status,
		DeliveryDate: req.PurchaseOrder.DeliveryDate.AsTime(),
		ReceivedDate: req.PurchaseOrder.ReceivedDate.AsTime(),
	}

	if err := h.service.UpdatePurchaseOrder(po); err != nil {
		log.Printf("Failed to update purchase order: %v", err)
		return nil, err
	}

	return &vms_pb.PurchaseOrderResponse{
		Id:           po.ID.String(),
		VendorId:     po.VendorID.String(),
		OrderDetails: po.OrderDetails,
		Status:       po.Status,
		DeliveryDate: timestamppb.New(po.DeliveryDate).String(),
		ReceivedDate: timestamppb.New(po.ReceivedDate).String(),
	}, nil
}

func (h *PurchaseOrderHandler) TrackOrderStatus(ctx context.Context, req *vms_pb.TrackOrderStatusRequest) (*vms_pb.PurchaseOrderResponse, error) {
	po, err := h.service.TrackOrderStatus(uuid.MustParse(req.Id))
	if err != nil {
		log.Printf("Failed to track order status: %v", err)
		return nil, err
	}

	return &vms_pb.PurchaseOrderResponse{
		Id:           po.ID.String(),
		VendorId:     po.VendorID.String(),
		OrderDetails: po.OrderDetails,
		Status:       po.Status,
		DeliveryDate: timestamppb.New(po.DeliveryDate).String(),
		ReceivedDate: timestamppb.New(po.ReceivedDate).String(),
	}, nil
}

func (h *PurchaseOrderHandler) ReceiveGoods(ctx context.Context, req *vms_pb.ReceiveGoodsRequest) (*vms_pb.ReceiveGoodsResponse, error) {
	if err := h.service.ReceiveGoods(uuid.MustParse(req.Id), req.ReceivedDate.AsTime()); err != nil {
		log.Printf("Failed to receive goods: %v", err)
		return nil, err
	}

	return &vms_pb.ReceiveGoodsResponse{Message: "Goods received successfully"}, nil
}
