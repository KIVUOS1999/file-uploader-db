package store

import (
	"database/sql"
	"net/http"

	"github.com/KIVUOS1999/easyApi/app"
	easyError "github.com/KIVUOS1999/easyApi/errors"
	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/KIVUOS1999/file-uploader-db/models"
	"github.com/google/uuid"
)

func (s *storeStruct) GetFileDetails(ctx *app.Context, fileID string) (*models.FileDetailStructure, error) {
	fileName := ""
	fileSize := 0
	fileDetails := models.FileDetailStructure{}

	err := s.db.QueryRow(FETCH_FILE_DETAILS_BY_FILE_ID, fileID).Scan(
		&fileDetails.ID,
		&fileName,
		&fileSize,
		&fileDetails.TotalChunks,
		&fileDetails.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn("Entry not found for file_id:", fileID)
			return nil, &easyError.CustomError{
				StatusCode: http.StatusNotFound,
				Response:   "could not find entry for: " + fileID,
			}
		}

		log.Errorf("file_details queryrow: %+v - %+v", fileID, err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *storeStruct) GetFilesByUser(ctx *app.Context, userID string) ([]models.FileDetailStructure, error) {
	rows, err := s.db.Query(FETCH_FILE_PER_USER, userID)
	if err != nil {
		log.Error("query", err.Error())
		return nil, err
	}

	defer rows.Close()

	res := []models.FileDetailStructure{}
	for rows.Next() {
		fileDetail := models.FileDetailStructure{}

		if err := rows.Scan(&fileDetail.ID, &fileDetail.Meta.Name, &fileDetail.Meta.Size, &fileDetail.CreatedAt); err != nil {
			log.Error("scan:", err.Error())
			return nil, err
		}

		res = append(res, fileDetail)
	}

	return res, nil
}

func (s *storeStruct) GetChunksByOrder(ctx *app.Context, fileID string) ([]models.FileChunkStructure, error) {
	log.Debugf("Entry - %+v", fileID)
	rows, err := s.db.Query(FETCH_CHUNK_BY_ORDER, fileID)
	if err != nil {
		log.Error(fileID, err.Error())
		return nil, err
	}

	defer rows.Close()

	res := []models.FileChunkStructure{}
	for rows.Next() {
		chunkID := ""
		checkSum := ""
		order := 0

		if err := rows.Scan(&chunkID, &checkSum, &order); err != nil {
			log.Error(fileID, err.Error())
			return nil, err
		}

		chunkUUID, _ := uuid.Parse(chunkID)
		fileUUID, _ := uuid.Parse(fileID)

		chunk := models.FileChunkStructure{
			ID:       chunkUUID,
			FileID:   fileUUID,
			CheckSum: checkSum,
			Order:    order,
		}

		res = append(res, chunk)
	}

	log.Debugf("Exit - %+v", res)

	return res, nil
}
