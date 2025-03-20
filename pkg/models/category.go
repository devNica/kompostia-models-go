package models

import (
	"database/sql"

	schema "github.com/devnica/kompostia-models-go/pkg/schemas"
	"github.com/google/uuid"
)

type CategoryModel struct {
	Id                uuid.UUID                `gorm:"primaryKey;column:id;type:uuid;unique;default:gen_random_uuid()"`
	CategoryName      string                   `gorm:"column:category_name;type:varchar(100);not null;unique"`
	Nomeclature       string                   `gorm:"column:nomeclature;type:varchar(10);not null;unique"`
	ParentId          sql.NullString           `gorm:"column:parent_id;type:uuid;null"`
	MultiLangValues   []schema.MultiLangSchema `gorm:"column:multi_lang_values;type:jsonb;default:'[]'"`
	MultiLangIsActive bool                     `gorm:"column:multi_lang_is_active;type:bool;not null;default:false"`
	IsActive          bool                     `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt         int64                    `gorm:"column:created_at;type:bigint;not null;"`
	UpdatedAt         int64                    `gorm:"column:updated_at;type:bigint;null;default:0"`
	Items             []CatalogItemModel       `gorm:"foreignKey:category_id"`
}

// Establecer nombre personalizado de la tabla
func (CategoryModel) TableName() string {
	return "category"
}
