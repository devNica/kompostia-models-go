package interfaces

import "github.com/google/uuid"

type FileSchema struct {
	Id       uuid.UUID
	FileType string
	FileSize int32
	Binary   []byte
}
