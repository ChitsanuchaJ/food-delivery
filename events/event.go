package events

const (
	TOPIC_ORDER_CREATED   = "order_created"
	TOPIC_ORDER_ACCEPTED  = "order_accepted"
	TOPIC_ORDER_PICKED_UP = "order_picked_up"
	TOPIC_ORDER_DELIVERED = "order_delivered"

	ORDER_STATUS_CREATED   = "created"
	ORDER_STATUS_ACCEPTED  = "accepted"
	ORDER_STATUS_PICKED_UP = "picked_up"
	ORDER_STATUS_DELIVERED = "delivered"
	ORDER_STATUS_SENT      = "sent"

	RECIPIENT_CUSTOMER   = "customer"
	RECIPIENT_RESTAURANT = "restaurant"
	RECIPIENT_RIDER      = "rider"
)

type Event interface {
	GetName() string
}

//////////////////////////////////////////////////////

type OrderCreate struct {
	OrderId      string
	RestaurantId string
	Items        []Item
}

type Item struct {
	MenuId   string
	Quantity int
}

func (obj OrderCreate) GetName() string {
	return TOPIC_ORDER_CREATED
}

//////////////////////////////////////////////////////

type OrderAccept struct {
	OrderId string
}

func (obj OrderAccept) GetName() string {
	return TOPIC_ORDER_ACCEPTED
}

//////////////////////////////////////////////////////

type OrderPickUp struct {
	OrderId string
	RiderId string
}

func (obj OrderPickUp) GetName() string {
	return TOPIC_ORDER_PICKED_UP
}

//////////////////////////////////////////////////////

type OrderDelivery struct {
	OrderId string
	RiderId string
}

func (obj OrderDelivery) GetName() string {
	return TOPIC_ORDER_DELIVERED
}
