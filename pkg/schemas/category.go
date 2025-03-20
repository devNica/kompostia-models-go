package interfaces

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

// Estructura recomendada para registro y actualizacion de una categoria
type FlatCategorySchema struct {
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
type CategoryWithParentRefSchema struct {
	Id                string          `json:"id" gorm:"column:id"`
	CategoryName      string          `json:"categoryName" gorm:"column:categoryName"`
	Nomeclature       string          `json:"nomeclature" gorm:"column:nomeclature"`
	ParentId          *string         `json:"parentId" gorm:"column:parentId"` // Puede ser null
	MultiLangCategory json.RawMessage `json:"multiLangCategory" gorm:"column:multiLangCategory"`
	MultiLangIsActive bool            `json:"multiLangIsActive" gorm:"column:multiLangIsActive"`
	ParentName        *string         `json:"parentName" gorm:"column:parentName"` // Puede ser null
}

// Estructura recomendada para consulta jerarquica de las categorias
type HierarchicalCategoryRelationshipsSchema struct {
	CategoryId        string          `json:"categoryId" gorm:"column:categoryId"`
	CategoryName      string          `json:"categoryName" gorm:"column:categoryName"`
	Nomeclature       string          `json:"nomeclature" gorm:"column:nomeclature"`
	ParentId          *string         `json:"parentId" gorm:"column:parentId"` // Puede ser null
	MultiLangCategory json.RawMessage `json:"multiLangCategory" gorm:"column:multiLangCategory"`
	MultiLangIsActive bool            `json:"multiLangIsActive" gorm:"column:multiLangIsActive"`
}
