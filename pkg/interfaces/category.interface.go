package interfaces

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

// Estructura recomendada para registro y actualizacion de una categoria
type FlatCategoryI struct {
	Id                uuid.UUID
	CategoryName      string
	Nomeclature       string
	ParentId          *uuid.UUID
	MultiLangValues   pgtype.JSONB
	MultiLangIsActive bool
	IsActive          bool
	CreatedAt         int64
	UpdatedAt         int64
}

// Estructura recomendada para listar todas las categorias
type ReflexCategoryI struct {
	Id                string
	CategoryName      string
	Nomeclature       string
	ParentID          *string
	MultiLangCategory []MultiLangModel
	MultiLangIsActive bool
	ParentName        *string
}
