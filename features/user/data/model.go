package data

import (
	"go-clean-arch/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	HP string
	Address string
	Role string
	Password string
}

func ToCore(data User) user.Core {
	return user.Core{
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
		HP: data.HP,
		Address: data.Address,
		Role: data.Role,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model: gorm.Model{ID: data.ID},
		Name: data.Name,
		Email: data.Email,
		HP: data.HP,
		Address: data.Address,
		Role: data.Role,
		Password: data.Password,
	}
}