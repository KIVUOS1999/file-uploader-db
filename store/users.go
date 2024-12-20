package store

import (
	"time"

	"github.com/KIVUOS1999/easyApi/app"
	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/KIVUOS1999/file-uploader-db/models"
)

func (s *storeStruct) AddUser(ctx *app.Context, userData *models.Users) error {
	time := time.Now().UTC().Unix()

	_, err := s.db.Exec(
		INSERT_TO_USER,
		userData.ID,
		userData.Name,
		userData.Email,
		userData.Picture,
		time,
	)

	if err != nil {
		log.Errorf("upload_file exec: %v - %+v", userData.ID, err.Error())
		return err
	}

	return nil
}

func (s *storeStruct) GetUser(ctx *app.Context, userID string) (*models.Users, error) {
	user := models.Users{}

	err := s.db.QueryRow(FETCH_USER_BY_ID, userID).Scan(
		&user.Name,
		&user.Email,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &user, nil
}
