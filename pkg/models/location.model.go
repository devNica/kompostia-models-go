package models

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type LocationTypeModel struct {
	Id              uuid.UUID              `gorm:"primaryKey;column:id;type:varchar(36);unique"`
	Name            string                 `gorm:"column:name;type:varchar(255);not null;unique"`
	SupplierCode    string                 `gorm:"column:supplier_code;type:varchar(10);not null;unique"`
	MultiLangValues pgtype.JSONB           `gorm:"column:multi_lang_values;type:jsonb;default:'[]'"`
	Rules           pgtype.JSONB           `gorm:"column:rules;type:jsonb;default:'[]'"`
	Locations       []StorageLocationModel `gorm:"foreignKey:location_type_id"`
}

// Establecer nombre personalizado de la tabla
func (LocationTypeModel) TableName() string {
	return "location_type"
}

type StorageLocationModel struct {
	Id             uuid.UUID              `gorm:"primaryKey;column:id;type:uuid;unique;default:gen_random_uuid()"`
	LocatioName    string                 `gorm:"column:location_name;type:varchar(100);not null;unique"`
	Nomeclature    string                 `gorm:"column:nomeclature;type:varchar(10);not null;unique"`
	ParentId       sql.NullString         `gorm:"column:parent_id;type:uuid;null"`
	HasAccounting  bool                   `gorm:"column:has_accounting;type:bool;not null;default:true"`
	IsActive       bool                   `gorm:"column:is_active;type:bool;not null;default:true"`
	CreatedAt      int64                  `gorm:"column:created_at;type:bigint;not null;"`
	UpdatedAt      int64                  `gorm:"column:updated_at;type:bigint;null;default:0"`
	LocationTypeId uuid.UUID              `gorm:"column:location_type_id;type;primaryKey"`
	Items          []ItemHasLocationModel `gorm:"foreignKey:location_id"`
}

// Establecer nombre personalizado de la tabla
func (StorageLocationModel) TableName() string {
	return "storage_location"
}
