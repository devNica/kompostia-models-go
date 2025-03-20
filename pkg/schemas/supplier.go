package schemas

import "github.com/google/uuid"

type SupplierSchema struct {
	Id           uuid.UUID
	SupplierName string
	SupplierCode string
	IsActive     bool
	CreatedAt    int64
	UpdatedAt    int64
}
