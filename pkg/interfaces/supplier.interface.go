package interfaces

import "github.com/google/uuid"

type SupplierI struct {
	Id           uuid.UUID
	SupplierName string
	SupplierCode string
	IsActive     bool
	CreatedAt    uint64
	UpdatedAt    uint64
}
