package schemas

import (
	"net/http"
)

type Product struct {
	ID       int64   `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

func (p Product) Validate() SchemaResponse {
	if p.Name == "" {
		return SchemaResponse{Error: true, Message: "\"name\" is required", Status: http.StatusBadRequest}
	}
	if len(p.Name) < 3 {
		return SchemaResponse{Error: true, Message: "\"name\" must have at least 3 characters", Status: http.StatusBadRequest}
	}
	if p.Quantity <= 0 {
		return SchemaResponse{Error: true, Message: "\"quantity\" must be greater than 0", Status: http.StatusBadRequest}
	}
	if p.Price <= 0 {
		return SchemaResponse{Error: true, Message: "\"price\" must be greater than 0", Status: http.StatusBadRequest}
	}
	return SchemaResponse{Error: false, Message: "", Status: 0}
}
