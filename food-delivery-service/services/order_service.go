package services

import (
	"errors"
	"events"
	"fmt"
	"food-delivery-service/producer"
	"food-delivery-service/repositories"
	"strconv"
	"time"

	"math/rand"
)

type orderService struct {
	eventProducer  producer.EventProducer
	restaurantRepo repositories.RestaurantRepository
	menuRepo       repositories.MenuRepository
}

func NewOrderService(
	eventProducer producer.EventProducer,
	restaurantRepo repositories.RestaurantRepository,
	menuRepo repositories.MenuRepository,
) OrderService {
	return orderService{eventProducer, restaurantRepo, menuRepo}
}

func (s orderService) PlaceOrder(orderReq OrderRequest) (orderResp *OrderResponse, err error) {

	fmt.Println("PlaceOrder called")

	// Logic for validate request
	// Is existing restaurantId ?
	// slice menu empty ?
	// Is existing menu in restaurant ?

	restaurantResult, err := s.restaurantRepo.GetRestaurantByID(orderReq.RestaurantId)
	if err != nil {
		return nil, errors.New("incorrect restaurant id")
	}

	allMenu, err := s.menuRepo.GetMenu(orderReq.RestaurantId)
	if err != nil {
		return nil, err
	}

	var notiMessage string

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	orderId := strconv.Itoa(random.Intn(1000))

	notiMessage = fmt.Sprintf("Order id: %v \"create\"", orderId)

	totalPrice := 0.0
	fmt.Println("--------------------------------------------------------")
	msg := fmt.Sprintf("Restaurant name: %v", restaurantResult.Name)
	fmt.Println(msg)
	notiMessage += "\n" + msg

	msg = "Menu name:"
	fmt.Println(msg)
	notiMessage += "\n" + msg
	for _, item := range orderReq.Items {
		menu := findMenuByID(item.MenuId, allMenu)
		if menu != nil {
			price := menu.Price * float64(item.Quantity)
			msg := fmt.Sprintf("%v | Quantity: %v | Price: %v$", menu.Name, item.Quantity, price)
			fmt.Println(msg)
			notiMessage += "\n" + msg
			totalPrice += price
		} else {
			fmt.Println("Menu id:", item.MenuId, "is unavailable")
			return nil, errors.New("incorrect menu id")
		}
	}
	msg = fmt.Sprintf("Order total amount: %v$", totalPrice)
	fmt.Println(msg)
	notiMessage += "\n" + msg
	fmt.Println("--------------------------------------------------------")

	items := []events.Item{}
	for _, item := range orderReq.Items {
		items = append(items, events.Item{
			MenuId:   item.MenuId,
			Quantity: item.Quantity,
		})
	}

	event := events.OrderCreate{
		OrderId:      orderId,
		RestaurantId: orderReq.RestaurantId,
		Items:        items,
		OptField:     events.OptionalField{Message: notiMessage},
	}

	// Payment mock
	err = Payment(orderId, totalPrice)
	if err != nil {
		fmt.Println("payment failed")
		return nil, errors.New("payment failed")
	}

	fmt.Println("Publish event:", event.GetTopicName())
	err = s.eventProducer.Produce(event)
	if err != nil {
		return nil, err
	}

	orderResp = &OrderResponse{
		OrderId: orderId,
		Status:  events.ORDER_STATUS_CREATED,
	}

	return orderResp, nil
}

func findMenuByID(id string, menus []repositories.Menu) *repositories.Menu {
	for _, menu := range menus {
		if id == menu.ID {
			return &menu
		}
	}
	return nil
}
