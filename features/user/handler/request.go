package handler

import "go-clean-arch/features/user"

type RegisterRequest struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	HP string `json:"hp" form:"hp"`
	Address string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email string `json:"email"  form:"email"`
	Password string `json:"password" form:"password"`
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