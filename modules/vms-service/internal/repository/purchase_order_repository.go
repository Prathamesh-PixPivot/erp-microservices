package repository

import (
    "vms-service/internal/models"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "time"
)

type PurchaseOrderRepository struct {
    db *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) *PurchaseOrderRepository {
    return &PurchaseOrderRepository{db: db}
}

func (repo *PurchaseOrderRepository) CreatePurchaseOrder(po *models.PurchaseOrder) error {
    po.ID = uuid.New()
    return repo.db.Create(po).Error
}

func (repo *PurchaseOrderRepository) GetPurchaseOrderByID(id uuid.UUID) (*models.PurchaseOrder, error) {
    var po models.PurchaseOrder
    if err := repo.db.First(&po, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &po, nil
}

func (repo *PurchaseOrderRepository) UpdatePurchaseOrder(po *models.PurchaseOrder) error {
    return repo.db.Save(po).Error
}

func (repo *PurchaseOrderRepository) DeletePurchaseOrder(id uuid.UUID) error {
    return repo.db.Delete(&models.PurchaseOrder{}, "id = ?", id).Error
}

func (repo *PurchaseOrderRepository) TrackOrderStatus(id uuid.UUID) (*models.PurchaseOrder, error) {
    var po models.PurchaseOrder
    return &po, repo.db.First(&po, "id = ?", id).Error
}

func (repo *PurchaseOrderRepository) ReceiveGoods(id uuid.UUID, receivedDate time.Time) error {
    return repo.db.Model(&models.PurchaseOrder{}).Where("id = ?", id).Updates(map[string]interface{}{
        "status":       "received",
        "received_date": receivedDate,
    }).Error
}
