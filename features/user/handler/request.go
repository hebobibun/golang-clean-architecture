package handler

import "go-clean-arch/features/user"

type RegisterRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	HP string `json:"hp"`
	Address string `json:"address"`
	Role string `json:"role"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.HP = cnv.HP
		res.Address = cnv.Address
		res.Role = cnv.Role
		res.Password = cnv.Password
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}