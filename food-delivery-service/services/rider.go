package services

type RiderWrapper struct {
	Riders []Rider `json:"rider"`
}

type Rider struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//////////////////////////////////////////////////////

type PickUpOrderRequest struct {
	OrderID string `json:"order_id"`
	RiderID string `json:"rider_id"`
}

type PickUpOrderResponse struct {
	Status string `json:"status"`
}

//////////////////////////////////////////////////////

type DeliverOrderRequest struct {
	OrderID string `json:"order_id"`
	RiderID string `json:"rider_id"`
}

type DeliverOrderResponse struct {
	Status string `json:"status"`
}

type RiderService interface {
	GetRiders() (*RiderWrapper, error)
	PickUpOrder(PickUpOrderRequest) (*PickUpOrderResponse, error)
	DeliverOrder(DeliverOrderRequest) (*DeliverOrderResponse, error)
}
