package handler

import (
	"net/http"

	"github.com/KIVUOS1999/easyApi/app"
	easyError "github.com/KIVUOS1999/easyApi/errors"
	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/KIVUOS1999/file-uploader-db/constants"
	"github.com/KIVUOS1999/file-uploader-db/models"
	"github.com/KIVUOS1999/file-uploader-db/store"
)

type handlerStruct struct {
	store store.IStore
}

func New(store store.IStore) *handlerStruct {
	return &handlerStruct{
		store: store,
	}
}

func (h *handlerStruct) UploadFile(ctx *app.Context) (interface{}, error) {
	fileDetails := models.FileDetailStructure{}

	err := ctx.Bind(&fileDetails)
	if err != nil {
		return nil, &easyError.CustomError{
			StatusCode: http.StatusBadRequest,
			Response:   err.Error(),
		}
	}

	err = h.store.UploadFile(ctx, &fileDetails)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerStruct) UploadChunksData(ctx *app.Context) (interface{}, error) {
	chunkDetails := models.FileChunkStructure{}

	err := ctx.Bind(&chunkDetails)
	if err != nil {
		return nil, &easyError.CustomError{
			StatusCode: http.StatusBadRequest,
			Response:   err.Error(),
		}
	}

	err = h.store.UploadChunk(ctx, &chunkDetails)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerStruct) GetFileByUser(ctx *app.Context) (interface{}, error) {
	userID := ctx.PathParam(constants.USER_ID)
	log.Debug("user:", userID)

	return h.store.GetFilesByUser(ctx, userID)
}

func (h *handlerStruct) FileDetails(ctx *app.Context) (interface{}, error) {
	fileID := ctx.PathParam(constants.FILE_ID)

	data, err := h.store.GetFileDetails(ctx, fileID)
	if err != nil {
		return nil, err
	}

	log.Infof("data received %s - %+v", fileID, data)

	return data, nil
}

func (h *handlerStruct) GetChunks(ctx *app.Context) (interface{}, error) {
	fileID := ctx.PathParam(constants.FILE_ID)

	chunksArr, err := h.store.GetChunksByOrder(ctx, fileID)
	if err != nil {
		return nil, err
	}

	return chunksArr, nil
}
