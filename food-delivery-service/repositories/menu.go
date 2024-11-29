package repositories

import (
	"gorm.io/gorm"
)

type Menu struct {
	ID           string
	Name         string
	Price        float64
	Description  string
	RestaurantId string
}

type MenuRepository interface {
	GetMenu(string) ([]Menu, error)
}

func mockMenuData(db *gorm.DB) error {

	var count int64
	db.Model(&Menu{}).Count(&count)
	if count > 0 {
		return nil
	}

	products := []Menu{}

	products = append(products,
		Menu{ID: "1", Name: "Classic Burger", Price: 5.5, Description: "Juicy beef patty with lettuce, tomato, and cheese.", RestaurantId: "1"},
		Menu{ID: "2", Name: "Margherita Pizza", Price: 8.50, Description: "Fresh mozzarella, basil, and tomato sauce.", RestaurantId: "1"},
		Menu{ID: "3", Name: "Pasta Carbonara", Price: 12.00, Description: "Creamy sauce with pancetta and Parmesan.", RestaurantId: "2"},
		Menu{ID: "4", Name: "Caesar Salad", Price: 7.25, Description: "Romaine lettuce with Caesar dressing and croutons.", RestaurantId: "2"},
		Menu{ID: "5", Name: "Grilled Salmon", Price: 18.00, Description: "Freshly grilled salmon with lemon butter sauce.", RestaurantId: "3"},
		Menu{ID: "6", Name: "Chicken Curry", Price: 10.5, Description: "Spicy curry with tender chicken pieces.", RestaurantId: "3"},
		Menu{ID: "7", Name: "Veggie Wrap", Price: 6.75, Description: "Grilled vegetables wrapped in a soft tortilla.", RestaurantId: "4"},
		Menu{ID: "8", Name: "Steak Sandwich", Price: 9.50, Description: "Grilled steak with onions and cheese in a baguette.", RestaurantId: "4"},
		Menu{ID: "9", Name: "Sushi Platter", Price: 15.00, Description: "Assorted sushi rolls with wasabi and soy sauce.", RestaurantId: "5"},
		Menu{ID: "10", Name: "Chocolate Cake", Price: 4.5, Description: "Rich chocolate cake with a creamy ganache.", RestaurantId: "5"},
	)

	return db.Create(&products).Error
}
