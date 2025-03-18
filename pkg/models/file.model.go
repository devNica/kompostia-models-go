package models

import "github.com/google/uuid"

type FileModel struct {
	Id       uuid.UUID          `gorm:"primaryKey;column:id;type:varchar(36);unique"`
	FileType string             `gorm:"column:filetype;type:varchar(10);not null"`
	FileSize int32              `gorm:"column:filesize;type:bigint;not null"`
	Binary   []byte             `gorm:"column:binary;type:bytea;not null"`
	Items    []ItemHasFileModel `gorm:"foreignKey:file_id"`
}

// Establecer nombre personalizado de la tabla
func (FileModel) TableName() string {
	return "file"
}
