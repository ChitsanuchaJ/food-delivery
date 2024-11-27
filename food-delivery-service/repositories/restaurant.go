package repositories

import (
	"gorm.io/gorm"
)

type restaurant struct {
	ID   string
	Name string
}

type RestaurantRepository interface {
	GetRestaurants() ([]restaurant, error)
}

func mockRestaurantData(db *gorm.DB) error {

	var count int64
	db.Model(&restaurant{}).Count(&count)
	if count > 0 {
		return nil
	}

	products := []restaurant{}

	products = append(products,
		restaurant{ID: "1", Name: "Pizza world"},
		restaurant{ID: "2", Name: "Gag Donald"},
		restaurant{ID: "3", Name: "Larb Mai"},
		restaurant{ID: "4", Name: "Mai kin gor D"},
		restaurant{ID: "5", Name: "Mai tong dak"},
	)

	return db.Create(&products).Error
}
