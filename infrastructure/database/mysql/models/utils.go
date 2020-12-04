package models

import (
	"context"

	"gorm.io/gorm"
)

func Paginate(ctx context.Context, skip int64, limmit int64, filter interface{}) func(db *gorm.DB) *gorm.DB {
	if skip <= 0 {
		skip = 0
	} else {
		skip = (skip - 1) * limmit
	}

	if limmit <= 0 {
		limmit = 10
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skip)).Limit(int(limmit)).Where(filter)
	}
}
