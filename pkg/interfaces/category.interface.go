package interfaces

import (
	"github.com/google/uuid"
)

type CategoryI struct {
	Id                uuid.UUID
	CategoryName      string
	Nomeclature       string
	ParentId          *uuid.UUID
	MultiLangValues   []MultiLangModel
	MultiLangIsActive bool
	IsActive          bool
	CreatedAt         int64
	UpdatedAt         int64
}
