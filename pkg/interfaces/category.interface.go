package interfaces

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type CategoryI struct {
	Id                uuid.UUID
	CategoryName      string
	Nomeclature       string
	ParentId          uuid.UUID
	MultiLangValues   pgtype.JSONB
	MultiLangIsActive bool
	IsActive          bool
	CreatedAt         int64
	UpdatedAt         int64
}
