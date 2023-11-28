package api

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/emicklei/go-restful"
	"github.com/fillipehmeireles/order-service/core/domain/order/ports"
	handlers "github.com/fillipehmeireles/order-service/pkg/handlers/order"
	"github.com/fillipehmeireles/order-service/pkg/handlers/order/dto"
)

const USER_ID_PARAM = "user_id"
const ORDER_ID_PARAM = "id"

type OrderHandler struct {
	orderUseCase ports.OrderUseCase
}

func NewOrderHandler(orderUseCase ports.OrderUseCase, ws *restful.WebService) *OrderHandler {
	orderHandler := &OrderHandler{
		orderUseCase: orderUseCase,
	}

	ws.Route(ws.POST("/orders").To(orderHandler.Create))
	ws.Route(ws.GET("/orders").To(orderHandler.GetAll))
	ws.Route(ws.GET(fmt.Sprintf("orders/{%s}", ORDER_ID_PARAM)).To(orderHandler.GetOne))
	ws.Route(ws.GET(fmt.Sprintf("orders/user/{%s}", USER_ID_PARAM)).To(orderHandler.GetByUser))
	ws.Route(ws.DELETE(fmt.Sprintf("orders/{%s}", USER_ID_PARAM)).To(orderHandler.Delete))

	return orderHandler
}

func (oH *OrderHandler) Create(req *restful.Request, resp *restful.Response) {
	var newOrder dto.CreateOrderRequestDto
	if err := req.ReadEntity(&newOrder); err != nil {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	if err := oH.orderUseCase.Create(newOrder); err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	resp.WriteAsJson(handlers.SuccessResponse{Data: handlers.OKOrderCreated})
}

func (oH *OrderHandler) GetAll(req *restful.Request, resp *restful.Response) {
	orders, err := oH.orderUseCase.GetAll()
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: orders})
}

func (oH *OrderHandler) GetOne(req *restful.Request, resp *restful.Response) {
	orderID := req.PathParameter(ORDER_ID_PARAM)
	if orderID == "" {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: handlers.ErrNoOrderIDProvided.Error()}))
		return
	}

	id, err := strconv.Atoi(orderID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	order, err := oH.orderUseCase.GetOne(id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: order})
}

func (oH *OrderHandler) GetByUser(req *restful.Request, resp *restful.Response) {
	userID := req.PathParameter(USER_ID_PARAM)
	if userID == "" {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: handlers.ErrNoUserIDProvided.Error()}))
		return
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	orders, err := oH.orderUseCase.GetByUser(id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: orders})
}

func (oH *OrderHandler) Delete(req *restful.Request, resp *restful.Response) {
	orderID := req.PathParameter(USER_ID_PARAM)
	if orderID == "" {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: handlers.ErrNoUserIDProvided.Error()}))
		return
	}

	id, err := strconv.Atoi(orderID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	if err := oH.orderUseCase.Delete(id); err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: handlers.OKOrderCreated})
}
