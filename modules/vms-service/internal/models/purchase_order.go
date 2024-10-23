package models

import (
    "time"
    "github.com/google/uuid"
)

type PurchaseOrder struct {
    ID           uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
    VendorID     uuid.UUID  `gorm:"type:uuid;index" json:"vendor_id"`
    Vendor       Vendor     `gorm:"foreignKey:VendorID"`
    OrderDetails string     `json:"order_details"`
    Status       string     `json:"status"`
    DeliveryDate time.Time  `json:"delivery_date"`
    ReceivedDate time.Time  `json:"received_date"`
}
