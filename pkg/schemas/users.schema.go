package schemas

import "net/http"

type User struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (u User) Validate() SchemaResponse {
	if u.Password == "" {
		return SchemaResponse{Error: true, Message: "\"password\" is required", Status: http.StatusBadRequest}
	}
	if len(u.Password) < 6 {
		return SchemaResponse{Error: true, Message: "\"password\" must have at least 6 characters", Status: http.StatusBadRequest}
	}
	if u.Email == "" {
		return SchemaResponse{Error: true, Message: "\"email\" is required", Status: http.StatusBadRequest}
	}
	return SchemaResponse{Error: false, Message: "", Status: 0}
}
