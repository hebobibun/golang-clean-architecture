package data

import (
	"go-clean-arch/features/book/data"
	"go-clean-arch/features/user"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name string
	Email string
	HP string
	Address string
	Password string
	Books []data.Books `gorm:"foreignkey:UserID"`
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
		HP: data.HP,
		Address: data.Address,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) Users {
	return Users{
		Model: gorm.Model{ID: data.ID},
		Name: data.Name,
		Email: data.Email,
		HP: data.HP,
		Address: data.Address,
		Password: data.Password,
	}
}