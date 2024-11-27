package services

import (
	"errors"
	"events"
	"fmt"
	"food-delivery-services/producer"
	"strconv"
	"time"

	"math/rand"
)

type orderService struct {
	eventProducer     producer.EventProducer
	restaurantService RestaurantService
	menuService       MenuService
}

func NewOrderService(
	eventProducer producer.EventProducer,
	restaurantService RestaurantService,
	menuService MenuService,
) OrderService {
	return orderService{eventProducer, restaurantService, menuService}
}

func (ordServ orderService) PlaceOrder(orderReq OrderRequest) (orderResp *OrderResponse, err error) {

	// Logic for validate request
	// Is existing restaurantId ?
	// slice menu empty ?
	// Is existing menu in restaurant ?

	restaurants, err := ordServ.restaurantService.GetRestaurants()
	if err != nil {
		return nil, err
	}

	isExistRestaurant := findRestaurantByID(orderReq.RestaurantId, restaurants.Restaurants)
	if !isExistRestaurant {
		return nil, errors.New("incorrect restaurant id")
	}

	allMenu, err := ordServ.menuService.GetMenus(orderReq.RestaurantId)
	if err != nil {
		return nil, err
	}

	totalPrice := 0.0
	fmt.Println("--------------------------------------------------------")
	fmt.Println("Menu name:")
	for _, item := range orderReq.Items {
		menu := findMenuByID(item.MenuId, allMenu.Menus)
		if menu != nil {
			price := menu.Price * float64(item.Quantity)
			fmt.Println(menu.Name, "| Quantity:", item.Quantity, "| Price:", price, "$")
			totalPrice += price
		} else {
			fmt.Println("Menu id:", item.MenuId, "is unavailable")
			return nil, errors.New("incorrect menu id")
		}
	}
	fmt.Println("Order total amount:", totalPrice, "$")
	fmt.Println("--------------------------------------------------------")

	items := []events.Item{}
	for _, item := range orderReq.Items {
		items = append(items, events.Item{
			MenuId:   item.MenuId,
			Quantity: item.Quantity,
		})
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	orderId := strconv.Itoa(random.Intn(1000))
	event := events.OrderCreate{
		OrderId:      orderId,
		RestaurantId: orderReq.RestaurantId,
		Items:        items,
	}

	// Payment mock
	err = Payment(orderId, totalPrice)
	if err != nil {
		fmt.Println("payment failed")
		return nil, errors.New("payment failed")
	}

	fmt.Println("Publish event:", event.GetTopicName())
	err = ordServ.eventProducer.Produce(event)
	if err != nil {
		return nil, err
	}

	orderResp = &OrderResponse{
		OrderId: orderId,
		Status:  events.ORDER_STATUS_CREATED,
	}

	return orderResp, nil
}

func findRestaurantByID(id string, restaurants []Restaurant) bool {
	for _, restaurant := range restaurants {
		if id == restaurant.ID {
			return true
		}
	}

	return false
}

func findMenuByID(id string, menus []Menu) *Menu {
	for _, menu := range menus {
		if id == menu.ID {
			return &menu
		}
	}
	return nil
}
