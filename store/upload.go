package store

import (
	"time"

	"github.com/KIVUOS1999/easyApi/app"
	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/KIVUOS1999/file-uploader-db/models"
)

func (s *storeStruct) UploadFile(ctx *app.Context, fileDetails *models.FileDetailStructure) error {
	time := time.Now().UTC().Unix()

	_, err := s.db.Exec(
		INSERT_TO_FILE_DETAILS,
		fileDetails.ID,
		fileDetails.Meta.Name,
		fileDetails.Meta.Size,
		fileDetails.TotalChunks,
		fileDetails.UserID,
		time,
	)

	if err != nil {
		log.Errorf("upload_file exec: %v - %+v", fileDetails.ID, err.Error())
		return err
	}

	log.Info("upload file success", fileDetails.ID)

	return nil
}

func (s *storeStruct) UploadChunk(ctx *app.Context, chunkDetails *models.FileChunkStructure) error {
	time := time.Now().UTC().Unix()

	_, err := s.db.Exec(
		INSERT_TO_CHUNK_DETAILS,
		chunkDetails.ID,
		chunkDetails.FileID,
		chunkDetails.CheckSum,
		chunkDetails.Order,
		time,
	)

	if err != nil {
		log.Errorf("upload_chunk exec: %+v - %+v", chunkDetails.ID, err.Error())
		return err
	}

	return nil
}
