package interfaces

import "github.com/google/uuid"

type SupplierI struct {
	Id           uuid.UUID
	SupplierName string
	SupplierCode string
	IsActive     bool
	CreatedAt    int64
	UpdatedAt    int64
}
