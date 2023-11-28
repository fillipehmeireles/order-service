package handlers

import "errors"

var (
	ErrNoUserIDProvided  = errors.New("no user id provided")
	ErrNoOrderIDProvided = errors.New("no user id provided")
)

const (
	OKOrderCreated = "Order created successfully"
	OKOrderDeleted = "Order deleted successfully"
)

type FailResponse struct {
	ErrorReason string `json:"error_reason"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
