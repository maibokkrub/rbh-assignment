package model

import (
	"html"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	AppointmentID int    `gorm:"indexed" json:"-"`
	UserID        int    `gorm:"foreignKey:ID" json:"-"`
	Body          string `gorm:"text" json:"body"`
	Creator       User   `gorm:"foreignKey:UserID" json:"creator"`
}

//
// DB OPS
//

func (comment *Comment) Create(db *gorm.DB) *gorm.DB {
	comment.Prepare()
	return db.Create(comment)
}

func (appointment *Comment) Prepare() {
	appointment.Body = html.EscapeString(appointment.Body)
}

func GetAllComment(db *gorm.DB, appointmentId int) (*[]Comment, error) {
	comments := []Comment{}

	res := db.Where("appointment_id == ?", appointmentId).Preload("Creator").Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return &comments, nil
}
