package models

import (
	"github.com/google/uuid"
)

type ItemBrandModel struct {
	Id        uuid.UUID          `gorm:"primaryKey;column:id;type:varchar(36);unique"`
	BrandName string             `gorm:"column:brand_name;type:varchar(100);not null;unique"`
	IsActive  bool               `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt int64              `gorm:"column:created_at;type:bigint;not null;"`
	UpdatedAt int64              `gorm:"column:updated_at;type:bigint;null;default:0"`
	Items     []CatalogItemModel `gorm:"foreignKey:brand_id"`
}

// Establecer nombre personalizado de la tabla
func (ItemBrandModel) TableName() string {
	return "item_brand"
}
