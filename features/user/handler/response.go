package handler

import (
	"go-clean-arch/features/user"
	"net/http"
	"strings"
)

type UserResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	HP string `json:"hp"`
	Address string `json:"address"`
}

func ToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
		HP: data.HP,
		Address: data.Address,
	}
}

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = ToResponse(data[0].(user.Core))
	} else {
		resp["data"] = ToResponse(data[0].(user.Core))
		resp["token"] = data[1].(string)
	}

	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "match") {
		code = http.StatusUnauthorized
	} else {
		code = http.StatusNotFound
	} 

	return code, resp
}