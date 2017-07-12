package model

type BaseSingleResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

type BaseListResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Content []interface{} `json:"content"`
}

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
