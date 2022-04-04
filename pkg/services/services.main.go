package services

type ServiceResponse struct {
	Status  int                    `json:"status"`
	Payload map[string]interface{} `json:"payload"`
}
