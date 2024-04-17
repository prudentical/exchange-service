package persistence

import "gorm.io/gorm"

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * size).Limit(size)
	}
}
