package webserver

import (
	"errors"
	"net/http"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo"
)

type FileHandler struct {
	fileSistemUsecase domain.FileUsecase
	validation        helper.ValidationRequest
}

func NewFilesystemHandler(e *echo.Group, fileUsecase domain.FileUsecase, valreq helper.ValidationRequest) {
	handler := &FileHandler{
		fileSistemUsecase: fileUsecase,
		validation:        valreq,
	}
	e.POST("/files-upload", handler.Uploads)

}
func (fh *FileHandler) Uploads(c echo.Context) error {
	var filesys domain.FileSystem
	form, err := c.MultipartForm()
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, errors.New("your form is empty"), c)
	}
	files := form.File["Filename"]
	filesys = domain.FileSystem{
		Filename: files,
	}
	if newErr := fh.validation.ValidateHandling(filesys); newErr != nil {
		return helper.Response(http.StatusBadRequest, nil, newErr, c)
	}
	urls, err := fh.fileSistemUsecase.UploadMultipleFiles(filesys.Filename)
	if err != nil {
		return helper.Response(http.StatusBadRequest, nil, err, c)
	}

	return helper.Response(http.StatusCreated, urls, nil, c)

}
