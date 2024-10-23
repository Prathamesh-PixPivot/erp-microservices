package services

import (
    "vms-service/internal/models"
    "vms-service/internal/repository"
    "github.com/google/uuid"
    "time"
)

type PurchaseOrderService struct {
    repo *repository.PurchaseOrderRepository
}

func NewPurchaseOrderService(repo *repository.PurchaseOrderRepository) *PurchaseOrderService {
    return &PurchaseOrderService{repo: repo}
}

func (service *PurchaseOrderService) CreatePurchaseOrder(po *models.PurchaseOrder) error {
    return service.repo.CreatePurchaseOrder(po)
}

func (service *PurchaseOrderService) GetPurchaseOrderByID(id uuid.UUID) (*models.PurchaseOrder, error) {
    return service.repo.GetPurchaseOrderByID(id)
}

func (service *PurchaseOrderService) UpdatePurchaseOrder(po *models.PurchaseOrder) error {
    return service.repo.UpdatePurchaseOrder(po)
}

func (service *PurchaseOrderService) DeletePurchaseOrder(id uuid.UUID) error {
    return service.repo.DeletePurchaseOrder(id)
}

func (service *PurchaseOrderService) TrackOrderStatus(id uuid.UUID) (*models.PurchaseOrder, error) {
    return service.repo.TrackOrderStatus(id)
}

func (service *PurchaseOrderService) ReceiveGoods(id uuid.UUID, receivedDate time.Time) error {
    return service.repo.ReceiveGoods(id, receivedDate)
}
