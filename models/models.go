package models

import "github.com/google/uuid"

type DBCredentials struct {
	Host     string
	Port     string
	Database string
	User     string
	Pass     string
}

type FileMetaData struct {
	Name string `json:"name"`
	Size uint64 `json:"file_size"`
}

type FileDetailStructure struct {
	Meta FileMetaData `json:"meta_data"`

	ID          uuid.UUID `json:"file_id"`
	TotalChunks int       `json:"total_chunks"`
	UserID      string    `json:"user_id"`
	CreatedAt   int64     `json:"created_at"`
}

type FileChunkStructure struct {
	ID       uuid.UUID `json:"chunk_id"`
	FileID   uuid.UUID `json:"file_id"`
	CheckSum string    `json:"check_sum"`
	Order    int       `json:"order"`
}

type Users struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Picture     string `json:"picture"`
	AllotedSize uint64 `json:"alloted_size"`
}
