package schemas

import "github.com/google/uuid"

type ItemBrandSchema struct {
	Id        uuid.UUID
	BrandName string
	IsActive  bool
	CreatedAt int64
	UpdatedAt int64
}
