package services

type ServiceResponse struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}
