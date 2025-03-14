package interfaces

import "github.com/google/uuid"

type ItemBrandI struct {
	Id        uuid.UUID
	BrandName string
	CreatedAt uint64
	UpdatedAt uint64
}
