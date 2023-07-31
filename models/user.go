package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string        `gorm:"not null" json:"username"`
	Password     string        `gorm:"not null" json:"-"`
	Email        string        `json:"email" validate:"required,email"`
	AvatarURL    string        `json:"avatar"`
	Appointments []Appointment `gorm:"foreignKey:created_by" json:"-"`
	Comments     []Comment     `gorm:"foreignKey:UserID" json:"-"`
}

func GetAllUsers(db *gorm.DB, page int) (*[]User, error) {
	users := []User{}

	res := db.Select("username", "email").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return &users, nil
}

//
// DB OPS
//
func (user *User) Create(db *gorm.DB) error {
	if tx := db.Create(user); tx.Error != nil {
		return tx.Error
	}

	return nil
}
