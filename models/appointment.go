package model

import (
	"errors"
	"html"
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	// gorm.Model
	ID          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"not null" json:"title" validate:"required,min=2,max=100"`
	Description string         `gorm:"text" json:"description"`
	Status      int8           `gorm:"default:1" json:"status" validate:"min=0,max=4"`
	CreatedBy   int            `json:"-"`
	Creator     User           `gorm:"foreignKey:ID;references:CreatedBy" json:"creator"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Comments    []Comment      `gorm:"foreignKey:AppointmentID" json:"comments"`
}

func GetAllAppointment(db *gorm.DB, page int) (*[]Appointment, error) {
	appointments := []Appointment{}
	offset := PAGE_SIZE * page

	res := db.Preload("Creator").Order("updated_at desc").Offset(offset).Limit(PAGE_SIZE).Find(&appointments)
	if res.Error != nil {
		return nil, res.Error
	}
	return &appointments, nil
}

func GetOneAppointment(db *gorm.DB, id int) (*Appointment, error) {
	appointment := Appointment{}

	res := db.First(&appointment, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &appointment, nil
}

func GetOneAppointmentWithComments(db *gorm.DB, id int) (*Appointment, error) {
	appointment := Appointment{}

	res := db.Preload("Comments").Preload("Users").First(&appointment, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &appointment, nil
}

//
// DB OPS
//

func (appointment *Appointment) Create(db *gorm.DB) *gorm.DB {
	appointment.Prepare()
	return db.Create(appointment)
}

func (appointment *Appointment) Prepare() {
	appointment.Title = html.EscapeString(appointment.Title)
	appointment.Description = html.EscapeString(appointment.Description)
}

func (appointment *Appointment) Update(db *gorm.DB) (*gorm.DB, error) {
	appointment.Prepare()
	res := db.Debug().Updates(appointment)
	if res.Error != nil {
		return nil, res.Error
	}
	return res, nil
}

func (appointment *Appointment) SoftDelete(db *gorm.DB) (*gorm.DB, error) {
	db.DB()
	res := db.Delete(appointment)
	if res.Error != nil {
		return nil, res.Error
	}
	return res, nil
}
