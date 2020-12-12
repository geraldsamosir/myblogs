package domain

import (
	"mime/multipart"
)

type FileSystem struct {
	Filename []*multipart.FileHeader `json:"files" validate:"required"`
}

type FileUsecase interface {
	UploadMultipleFiles(files []*multipart.FileHeader) ([]string, error)
}
