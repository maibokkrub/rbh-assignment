package dto

import (
	model "github.com/maibokkrub/simple-backend/models"
	"gopkg.in/go-playground/validator.v9"
)

type CreateUserDTO struct {
	DisplayName string `validate:"required,min=2,max=50" json:"displayName"`
	Password    string `validate:"required"`
	Email       string `validate:"required,email"`
	AvatarURL   string `json:"avatar"`
}

type LoginUserDTO struct {
	Email    string `validate:"required,email"`
	Password string
}
type FakeLoginDTO struct {
	ID int `json:"userID"`
}

func (d *CreateUserDTO) ToModel() (*model.User, error) {
	v := validator.New()

	if err := v.Struct(d); err != nil {
		// todo: clean up error message leak
		return nil, err
	}

	// todo: implement bcrypt hashing for password
	// password := d.Password

	return &model.User{
		DisplayName: d.DisplayName,
		Email:       d.Email,
		// Password:    password,
		AvatarURL: d.AvatarURL,
	}, nil
}

func (d *LoginUserDTO) ToModel() (*model.User, error) {
	v := validator.New()

	if err := v.Struct(d); err != nil {
		// todo: clean up error message leak
		return nil, err
	}

	// todo: implement bcrypt hashing for password
	// password := d.Password

	return &model.User{
		Email: d.Email,
		// Password: password,
	}, nil
}
