package models

import "github.com/google/uuid"

type SupplierModel struct {
	Id           uuid.UUID          `gorm:"primaryKey;column:id;type:varchar(36);unique"`
	SupplierName string             `gorm:"column:supplier_name;type:varchar(255);not null;unique"`
	SupplierCode string             `gorm:"column:supplier_code;type:varchar(10);not null;unique"`
	IsActive     bool               `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt    int64              `gorm:"column:created_at;type:bigint;not null;"`
	UpdatedAt    int64              `gorm:"column:updated_at;type:bigint;null;default:0"`
	Items        []CatalogItemModel `gorm:"foreignKey:supplier_id"`
}

// Establecer nombre personalizado de la tabla
func (SupplierModel) TableName() string {
	return "supplier"
}
