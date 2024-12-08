package store

import (
	"github.com/KIVUOS1999/easyApi/app"
	"github.com/KIVUOS1999/easyLogs/pkg/log"
)

func (s *storeStruct) RemoveFile(ctx *app.Context, fileID string) error {
	log.Info("Delete file:", fileID)

	tx, err := s.db.Begin()
	if err != nil {
		log.Error(fileID, err.Error())
		return err
	}

	defer func() {
		if err != nil {
			log.Warn("Rollback", fileID)
			tx.Rollback()
		}
	}()

	res, err := tx.Exec(DELETE_CHUNK_BY_FILE_ID, fileID)
	if err != nil {
		log.Error(fileID, err.Error())
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Error("Error getting rows affected", err.Error())
	}

	log.Debug("rows in chunk changed:", rowsAffected)

	res, err = tx.Exec(DELETE_FILE_BY_FILE_ID, fileID)
	if err != nil {
		log.Error(fileID, err.Error())
		return err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		log.Error("Error getting rows affected", err.Error())
	}

	log.Debug("rows in file changed:", rowsAffected)

	err = tx.Commit()
	if err != nil {
		log.Error(fileID, err.Error())
		return err
	}

	log.Info("Delete success", fileID)
	return nil
}
