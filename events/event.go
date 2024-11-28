package events

const (
	TOPIC_ORDER_CREATE   = "order_created"
	TOPIC_ORDER_ACCEPT   = "order_accepted"
	TOPIC_ORDER_PICK_UP  = "order_picked_up"
	TOPIC_ORDER_DELIVERY = "order_delivered"

	GROUP_ORDER_CREATE   = "order_create_group"
	GROUP_ORDER_ACCEPT   = "order_accept_group"
	GROUP_ORDER_PICK_UP  = "order_pick_up_group"
	GROUP_ORDER_DELIVERY = "order_delivery_group"

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
	GetTopicName() string
	// GetConsumerGroupName() string
}

type OptionalField struct {
	Message string
}

//////////////////////////////////////////////////////

type OrderCreate struct {
	OrderId      string
	RestaurantId string
	Items        []Item
	OptField     OptionalField
}

type Item struct {
	MenuId   string
	Quantity int
}

func (obj OrderCreate) GetTopicName() string {
	return TOPIC_ORDER_CREATE
}

// func (obj OrderCreate) GetConsumerGroupName() string {
// 	return GROUP_ORDER_CREATED
// }

//////////////////////////////////////////////////////

type OrderAccept struct {
	OrderId  string
	OptField OptionalField
}

func (obj OrderAccept) GetTopicName() string {
	return TOPIC_ORDER_ACCEPT
}

// func (obj OrderAccept) GetConsumerGroupName() string {
// 	return GROUP_ORDER_ACCEPTED
// }

//////////////////////////////////////////////////////

type OrderPickUp struct {
	OrderId  string
	RiderId  string
	OptField OptionalField
}

func (obj OrderPickUp) GetTopicName() string {
	return TOPIC_ORDER_PICK_UP
}

// func (obj OrderPickUp) GetConsumerGroupName() string {
// 	return GROUP_ORDER_PICKED_UP
// }

//////////////////////////////////////////////////////

type OrderDelivery struct {
	OrderId  string
	RiderId  string
	OptField OptionalField
}

func (obj OrderDelivery) GetTopicName() string {
	return TOPIC_ORDER_DELIVERY
}

// func (obj OrderDelivery) GetConsumerGroupName() string {
// 	return GROUP_ORDER_DELIVERED
// }
