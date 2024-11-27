package repositories

import (
	"gorm.io/gorm"
)

type rider struct {
	ID   string
	Name string
}

type RiderRepository interface {
	GetRiders() ([]rider, error)
}

func mockRiderData(db *gorm.DB) error {

	var count int64
	db.Model(&rider{}).Count(&count)
	if count > 0 {
		return nil
	}

	products := []rider{}

	products = append(products,
		rider{ID: "1", Name: "Rider 1"},
		rider{ID: "2", Name: "Rider 2"},
		rider{ID: "3", Name: "Rider 3"},
	)

	return db.Create(&products).Error
}
