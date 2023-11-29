package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/fillipehmeireles/order-service/pkg/handlers/order/dto"
)

func GenerateTokenPairWithOrderData(order dto.CreateOrderRequestDto) string {
	t := time.Now().Format(time.RFC850)
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%f%d%d%s", order.UserID, order.Amount, order.Direction, order.OrderType, t)))
}
