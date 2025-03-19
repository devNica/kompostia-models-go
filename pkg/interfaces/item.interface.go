package interfaces

import (
	"github.com/google/uuid"
)

type CatalogItemI struct {
	Id                uuid.UUID
	ItemName          string
	Description       string
	Reference         string
	Sku               string
	SupplierItemName  string
	SupplierItemCode  string
	MultiLangValues   []MultiLangModel
	MultiLangIsActive bool
	BrandId           uuid.UUID
	CategoryId        uuid.UUID
	SupplierId        uuid.UUID
	IsActive          bool
	CreatedAt         int64
	UpdatedAt         int64
}
