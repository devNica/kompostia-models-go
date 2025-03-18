package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type CategoryModel struct {
	Id                uuid.UUID          `gorm:"primaryKey;column:id;type:varchar(36);unique"`
	CategoryName      string             `gorm:"column:category_name;type:varchar(100);not null;unique"`
	Nomeclature       string             `gorm:"column:item_name;type:varchar(10);not null;unique"`
	ParentId          uuid.UUID          `gorm:"column:parent_id;type:varchar(36);null"`
	MultiLangValues   pgtype.JSONB       `gorm:"column:multi_lang_values;type:jsonb;default:'[]'"`
	MultiLangIsActive bool               `gorm:"column:multi_lang_is_active;type:bool;not null;default:false"`
	IsActive          bool               `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt         int64              `gorm:"column:created_at;type:bigint;not null;"`
	UpdatedAt         int64              `gorm:"column:updated_at;type:bigint;null;default:0"`
	Items             []CatalogItemModel `gorm:"foreignKey:category_id"`
}

// Establecer nombre personalizado de la tabla
func (CategoryModel) TableName() string {
	return "category"
}
