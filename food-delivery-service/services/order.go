package services

type OrderRequest struct {
	RestaurantId string `json:"restaurant_id"`
	Items        []Item `json:"items"`
}

type Item struct {
	MenuId   string `json:"menu_id"`
	Quantity int    `json:"quantity"`
}

type OrderResponse struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

type OrderService interface {
	PlaceOrder(OrderRequest) (*OrderResponse, error)
}
