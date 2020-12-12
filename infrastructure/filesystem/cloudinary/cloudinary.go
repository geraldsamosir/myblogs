package cloudinary

import (
	"fmt"
	"mime/multipart"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/komfy/cloudinary"
	"github.com/sirupsen/logrus"
)

type UploadHandler struct {
	cs *cloudinary.Service
}

func NewCloudinary(cs *cloudinary.Service) domain.FileUsecase {
	return &UploadHandler{cs}
}

func (h UploadHandler) UploadMultipleFiles(files []*multipart.FileHeader) ([]string, error) {
	urls := []string{}
	for _, file := range files {
		fmt.Println(file.Filename)
		data, err := file.Open()

		if err == nil {
			upResp, err := h.cs.Upload(file.Filename, data, true)
			if err != nil {
				logrus.Error("sini err upload", err)
				return nil, err
			} else {
				urls = append(urls, upResp.URL)
			}
		}
		data.Close()
	}
	return urls, nil
}
