package repositories

import (
	"gorm.io/gorm"
)

type menu struct {
	ID           string
	Name         string
	Price        float64
	Description  string
	RestaurantId string
}

type MenuRepository interface {
	GetMenus(string) ([]menu, error)
}

func mockMenuData(db *gorm.DB) error {

	var count int64
	db.Model(&menu{}).Count(&count)
	if count > 0 {
		return nil
	}

	products := []menu{}

	products = append(products,
		menu{ID: "1", Name: "Classic Burger", Price: 5.5, Description: "Juicy beef patty with lettuce, tomato, and cheese.", RestaurantId: "1"},
		menu{ID: "2", Name: "Margherita Pizza", Price: 8.50, Description: "Fresh mozzarella, basil, and tomato sauce.", RestaurantId: "1"},
		menu{ID: "3", Name: "Pasta Carbonara", Price: 12.00, Description: "Creamy sauce with pancetta and Parmesan.", RestaurantId: "2"},
		menu{ID: "4", Name: "Caesar Salad", Price: 7.25, Description: "Romaine lettuce with Caesar dressing and croutons.", RestaurantId: "2"},
		menu{ID: "5", Name: "Grilled Salmon", Price: 18.00, Description: "Freshly grilled salmon with lemon butter sauce.", RestaurantId: "3"},
		menu{ID: "6", Name: "Chicken Curry", Price: 10.5, Description: "Spicy curry with tender chicken pieces.", RestaurantId: "3"},
		menu{ID: "7", Name: "Veggie Wrap", Price: 6.75, Description: "Grilled vegetables wrapped in a soft tortilla.", RestaurantId: "4"},
		menu{ID: "8", Name: "Steak Sandwich", Price: 9.50, Description: "Grilled steak with onions and cheese in a baguette.", RestaurantId: "4"},
		menu{ID: "9", Name: "Sushi Platter", Price: 15.00, Description: "Assorted sushi rolls with wasabi and soy sauce.", RestaurantId: "5"},
		menu{ID: "10", Name: "Chocolate Cake", Price: 4.5, Description: "Rich chocolate cake with a creamy ganache.", RestaurantId: "5"},
	)

	return db.Create(&products).Error
}
