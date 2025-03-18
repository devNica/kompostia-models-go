package interfaces

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type LocationTypeI struct {
	Id              uuid.UUID
	Name            string
	SupplierCode    string
	MultiLangValues pgtype.JSONB
	Rules           pgtype.JSONB
}

type StorageLocationI struct {
	Id             uuid.UUID
	LocatioName    string
	Nomeclature    string
	ParentId       uuid.UUID
	HasAccounting  bool
	IsActive       bool
	CreatedAt      int64
	UpdatedAt      int64
	LocationTypeId uuid.UUID
}
