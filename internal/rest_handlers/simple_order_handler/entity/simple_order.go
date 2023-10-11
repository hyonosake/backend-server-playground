package entity

type SimpleOrderRequest struct {
	ProductId int64 `json:"productId"`
}

type Error struct {
	Message string `json:"message"`
}

type Result struct {
	Message string `json:"message"`
	OrderId int64  `json:"orderId"`
}

type SimpleOrderResponse struct {
	Error  *Error  `json:"error,omitempty"`
	Result *Result `json:"result,omitempty"`
}
