package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type CatalogItemModel struct {
	Id                uuid.UUID              `gorm:"primaryKey;column:id;type:uuid;unique;default:gen_random_uuid()"`
	ItemName          string                 `gorm:"column:item_name;type:varchar(255);not null;unique"`
	Description       string                 `gorm:"column:description;type:varchar(255);not null"`
	Reference         string                 `gorm:"column:reference;type:varchar(255);not null;unique"`
	Sku               string                 `gorm:"column:sku;type:varchar(20);not null;unique"`
	SupplierItemName  string                 `gorm:"column:supplier_item_name;type:varchar(255);not null;unique"`
	SupplierItemCode  string                 `gorm:"column:supplier_item_code;type:varchar(50);not null;unique"`
	MultiLangValues   pgtype.JSONB           `gorm:"column:multi_lang_values;type:jsonb;default:'[]'"`
	MultiLangIsActive bool                   `gorm:"column:multi_lang_is_active;type:bool;not null;default:false"`
	BrandId           uuid.UUID              `gorm:"column:brand_id;type;primaryKey"`
	CategoryId        uuid.UUID              `gorm:"column:category_id;type;primaryKey"`
	SupplierId        uuid.UUID              `gorm:"column:supplier_id;type;primaryKey"`
	IsActive          bool                   `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt         int64                  `gorm:"column:created_at;type:bigint;not null;"`
	UpdatedAt         int64                  `gorm:"column:updated_at;type:bigint;null;default:0"`
	Locations         []ItemHasLocationModel `gorm:"foreignKey:item_id"`
	Files             []ItemHasFileModel     `gorm:"foreignKey:item_id"`
}

// Establecer nombre personalizado de la tabla
func (CatalogItemModel) TableName() string {
	return "catalog_item"
}

type ItemHasLocationModel struct {
	ItemId     uuid.UUID `gorm:"column:item_id;type;primaryKey"`
	LocationId uuid.UUID `gorm:"column:location_id;type;primaryKey"`
}

// Establecer nombre personalizado de la tabla
func (ItemHasLocationModel) TableName() string {
	return "item_has_location"
}

type ItemHasFileModel struct {
	ItemId     uuid.UUID `gorm:"column:item_id;type;primaryKey"`
	LocationId uuid.UUID `gorm:"column:file_id;type;primaryKey"`
}

// Establecer nombre personalizado de la tabla
func (ItemHasFileModel) TableName() string {
	return "item_has_file"
}
