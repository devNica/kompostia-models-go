package schemas

import (
	"github.com/google/uuid"
)

type CatalogItemSchema struct {
	Id                uuid.UUID
	ItemName          string
	Description       string
	Reference         string
	Sku               string
	SupplierItemName  string
	SupplierItemCode  string
	MultiLangValues   []MultiLangSchema
	MultiLangIsActive bool
	BrandId           uuid.UUID
	CategoryId        uuid.UUID
	SupplierId        uuid.UUID
	IsActive          bool
	CreatedAt         int64
	UpdatedAt         int64
}
