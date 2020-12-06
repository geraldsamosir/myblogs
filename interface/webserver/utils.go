package webserver

import (
	"net/http"

	"github.com/geraldsamosir/myblogs/domain"
	"github.com/sirupsen/logrus"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrStatusUnauthorized:
		return http.StatusUnauthorized
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
