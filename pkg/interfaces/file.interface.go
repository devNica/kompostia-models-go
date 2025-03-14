package interfaces

import "github.com/google/uuid"

type FileI struct {
	Id       uuid.UUID
	FileType string
	FileSize uint64
	Binary   []byte
}
