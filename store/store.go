package store

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/KIVUOS1999/easyApi/app"
	easyError "github.com/KIVUOS1999/easyApi/errors"
	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/KIVUOS1999/file-uploader-db/models"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type storeStruct struct {
	db *sql.DB
}

func New(cred models.DBCredentials) (IStore, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cred.User, cred.Database, cred.Pass, cred.Host, cred.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Error in opening connection", err.Error())

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error("Error ping db", err.Error())

		return nil, err
	}

	log.Info("DB connection succes")

	return &storeStruct{db: db}, nil
}

func (s *storeStruct) UploadFile(ctx *app.Context, fileDetails *models.FileDetailStructure) error {
	time := time.Now().UTC().Unix()

	_, err := s.db.Exec(
		INSERT_TO_FILE_DETAILS,
		fileDetails.ID,
		fileDetails.Meta.Name,
		fileDetails.Meta.Size,
		fileDetails.TotalChunks,
		time,
	)

	if err != nil {
		log.Errorf("upload_file exec: %v - %+v", fileDetails.ID, err.Error())
		return err
	}

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
	rows, err := s.db.Query(FETCH_FILE_PER_USER)
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

const (
	INSERT_TO_FILE_DETAILS = `
		insert into file_details 
		(file_id,file_name,file_size,total_chunks,created_at)
		values 
		($1, $2, $3, $4, $5)
	`

	INSERT_TO_CHUNK_DETAILS = `
		insert into chunk_details
		(chunk_id,file_id,check_sum,chunk_order,created_at)
		values
		($1,$2,$3,$4,$5)
	`

	FETCH_FILE_DETAILS_BY_FILE_ID = `
		select * from file_deails
		where
		file_id=$1
	`

	FETCH_CHUNK_BY_ORDER = `
		select chunk_id,check_sum,chunk_order from chunk_details
		where file_id=$1
		order by chunk_order asc
	`

	// currently there is no user but in future there will be
	FETCH_FILE_PER_USER = `
	select file_id, file_name, file_size, created_at
	from file_details`
)
