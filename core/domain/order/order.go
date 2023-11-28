package order

type OrderType int

const (
	OrderTypeMarket OrderType = 1
	OrderTypeLimit  OrderType = 2
)

type Direction int

const (
	DirectionBuy  Direction = 1
	DirectionSell Direction = 2
)

type Order struct {
	ID        int
	UserID    int
	Pair      string
	Amount    float32
	Direction Direction
	OrderType OrderType
}
type Orders []Order

func NewOrder(id int, userID int, pair string, amount float32, direction Direction, orderType OrderType) *Order {

	return &Order{
		ID:        id,
		UserID:    userID,
		Pair:      pair,
		Amount:    amount,
		Direction: direction,
		OrderType: orderType,
	}
}
