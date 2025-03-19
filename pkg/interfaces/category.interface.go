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
	Id                string  `json:"id"`
	CategoryName      string  `json:"categoryName"`
	Nomeclature       string  `json:"nomeclature"`
	ParentID          *string `json:"parentId"`
	MultiLangCategory string  `json:"multiLangCategory"`
	MultiLangIsActive bool    `json:"multiLangIsActive"`
	ParentName        *string `json:"parentName"`
}
