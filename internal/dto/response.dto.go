package dto

type ResponseSuccess struct {
	Status  string `json:"status" example:"Success"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
	Message string `json:"message" example:"Data retrieved successfully"`
}

type ResponseError struct {
	Status  string `json:"status" example:"Error"`
	Message string `json:"message" example:"Failed get data"`
	Error   string `json:"errors,omitempty" example:"failed get data"`
}

type LoginResponse struct {
	ResponseSuccess
}

type RegisterResponse struct {
	ResponseSuccess
}