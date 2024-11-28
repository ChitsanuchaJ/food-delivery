package services

import (
	"errors"
	"events"
	"fmt"
	"food-delivery-service/producer"
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

	restaurantResult, err := findRestaurantByID(orderReq.RestaurantId, restaurants.Restaurants)
	if err != nil {
		return nil, errors.New("incorrect restaurant id")
	}

	allMenu, err := ordServ.menuService.GetMenus(orderReq.RestaurantId)
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
	notiMessage = notiMessage + "\n" + msg

	msg = "Menu name:"
	fmt.Println(msg)
	notiMessage = notiMessage + "\n" + msg
	for _, item := range orderReq.Items {
		menu := findMenuByID(item.MenuId, allMenu.Menus)
		if menu != nil {
			price := menu.Price * float64(item.Quantity)
			msg := fmt.Sprintf("%v | Quantity: %v | Price: %v$", menu.Name, item.Quantity, price)
			fmt.Println(msg)
			notiMessage = notiMessage + "\n" + msg
			totalPrice += price
		} else {
			fmt.Println("Menu id:", item.MenuId, "is unavailable")
			return nil, errors.New("incorrect menu id")
		}
	}
	msg = fmt.Sprintf("Order total amount: %v$", totalPrice)
	fmt.Println(msg)
	notiMessage = notiMessage + "\n" + msg
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

func findRestaurantByID(id string, restaurants []Restaurant) (*Restaurant, error) {
	for _, restaurant := range restaurants {
		if id == restaurant.ID {
			return &restaurant, nil
		}
	}

	errorMsg := fmt.Sprintf("Restaurant id %v isn't existing", id)
	return nil, errors.New(errorMsg)
}

func findMenuByID(id string, menus []Menu) *Menu {
	for _, menu := range menus {
		if id == menu.ID {
			return &menu
		}
	}
	return nil
}
