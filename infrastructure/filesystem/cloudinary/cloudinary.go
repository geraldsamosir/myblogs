package cloudinary

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
	"sync"

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
	errs := []string{}
	urlSuccess := make(chan string)
	urlErrors := make(chan interface{})
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go h.paralelUploads(file, urlSuccess, urlErrors, &wg)

	}

	go func() {
		wg.Wait()
		close(urlErrors)
		close(urlSuccess)
	}()
	for url := range urlSuccess {
		urls = append(urls, url)
	}
	for err := range urlErrors {
		errs = append(errs, fmt.Sprintf("%v", err))
	}
	if len(errs) > 0 {
		return nil, errors.New(strings.Join(errs, " "))

	}
	return urls, nil
}

func (h UploadHandler) paralelUploads(file *multipart.FileHeader, success chan string, errors chan interface{}, wg *sync.WaitGroup) {
	defer (*wg).Done()
	logrus.Info(file.Filename)
	data, err := file.Open()
	if err == nil {
		upResp, err := h.cs.Upload(file.Filename, data, true)
		if err != nil {
			logrus.Error("sini err upload", err)
			errors <- err
		} else {
			success <- upResp.URL
		}
	}
	data.Close()
}
