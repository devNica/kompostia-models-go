package interfaces

import (
	"github.com/google/uuid"
)

type LocationTypeSchema struct {
	Id              uuid.UUID
	Name            string
	SupplierCode    string
	MultiLangValues []MultiLangSchema
	Rules           []SuggestedLocations
}

type StorageLocationSchema struct {
	Id             uuid.UUID
	LocatioName    string
	Nomeclature    string
	ParentId       *uuid.UUID
	HasAccounting  bool
	IsActive       bool
	CreatedAt      int64
	UpdatedAt      int64
	LocationTypeId uuid.UUID
}
