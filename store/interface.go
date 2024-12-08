package store

import (
	"github.com/KIVUOS1999/easyApi/app"
	"github.com/KIVUOS1999/file-uploader-db/models"
)

type IStore interface {
	UploadFile(ctx *app.Context, fileDetails *models.FileDetailStructure) error
	UploadChunk(ctx *app.Context, chunkDetails *models.FileChunkStructure) error

	GetFilesByUser(ctx *app.Context, userID string) ([]models.FileDetailStructure, error)
	GetFileDetails(ctx *app.Context, fileID string) (*models.FileDetailStructure, error)
	GetChunksByOrder(ctx *app.Context, fileID string) ([]models.FileChunkStructure, error)

	RemoveFile(ctx *app.Context, fileID string) error
}
