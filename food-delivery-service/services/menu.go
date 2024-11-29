package services

type MenuWrapper struct {
	RestaurantId string `json:"restaurant_id"`
	Menus        []Menu `json:"menu"`
}

type Menu struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type MenuService interface {
	GetMenu(string) (*MenuWrapper, error)
}
