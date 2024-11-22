package services

type RestaurantWrapper struct {
	Restaurants []Restaurant `json:"restaurant"`
}

type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RestaurantService interface {
	GetRestaurants() (*RestaurantWrapper, error)
}
