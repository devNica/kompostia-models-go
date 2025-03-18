package interfaces

import "github.com/google/uuid"

type ItemBrandI struct {
	Id        uuid.UUID
	BrandName string
	IsActive  bool
	CreatedAt int64
	UpdatedAt int64
}
