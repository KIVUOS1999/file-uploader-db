package main

import (
	"github.com/KIVUOS1999/easyApi/app"
	"github.com/KIVUOS1999/file-uploader-db/handler"
	"github.com/KIVUOS1999/file-uploader-db/models"
	"github.com/KIVUOS1999/file-uploader-db/store"
)

func main() {
	app := app.New()

	dbCredentials := models.DBCredentials{
		Host:     app.Configs.Get("DB_HOST"),
		Port:     app.Configs.Get("DB_PORT"),
		Database: app.Configs.Get("DB_NAME"),
		User:     app.Configs.Get("DB_USER"),
		Pass:     app.Configs.Get("DB_PASSWORD"),
	}

	s, err := store.New(dbCredentials)
	if err != nil {
		return
	}

	h := handler.New(s)

	app.Post("/upload_file", h.UploadFile)
	app.Post("/upload_chunks", h.UploadChunksData)

	app.Get("/files/{user-id}", h.GetFileByUser)
	app.Get("/chunks/{file-id}", h.GetChunks)

	app.Delete("/file/{file-id}", h.DeleteFile)

	app.Start()
}
