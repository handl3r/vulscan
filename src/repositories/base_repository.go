package repositories

import "gorm.io/gorm"

type baseRepository struct {
	db *gorm.DB
}

func newBaseRepository(db *gorm.DB) *baseRepository {
	return &baseRepository{db: db}
}
