package dto

import (
	"errors"

	model "github.com/maibokkrub/simple-backend/models"
	"gopkg.in/go-playground/validator.v9"
)

type CreateAppointmentDTO struct {
	Title       string `json:"title" validate:"required,min=2,max=100"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"createdBy"`
}

type UpdateAppointmentDTO struct {
	ID          uint   `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=2,max=100"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
}

type SetAppointmentStatusDTO struct {
	ID     uint
	Status uint8 `json:"status" validate:"min=1,max=3"`
}

func (d *CreateAppointmentDTO) ToModel() (*model.Appointment, error) {
	v := validator.New()

	if err := v.Struct(d); err != nil {
		// todo: clean up error message leak
		return nil, err
	}

	return &model.Appointment{
		Title:       d.Title,
		Description: d.Description,
		CreatedBy:   d.CreatedBy,
	}, nil
}

func (d *UpdateAppointmentDTO) ToModel(oldData *model.Appointment) (*model.Appointment, error) {
	v := validator.New()

	if err := v.Struct(d); err != nil {
		// todo: clean up error message leak
		return nil, err
	}

	if d.Status < 0 || d.Status > 4 {
		return nil, errors.New("Invalid input")
	}

	if d.Title != "" {
		oldData.Title = d.Title
	}
	if d.Description != "" {
		oldData.Description = d.Description
	}
	if d.Status > 0 {
		oldData.Status = d.Status
	}

	return oldData, nil
}

func (d *SetAppointmentStatusDTO) ToModel(oldData *model.Appointment) (*model.Appointment, error) {
	v := validator.New()

	if err := v.Struct(d); err != nil {
		// todo: clean up error message leak
		return nil, err
	}

	oldData.Status = d.Status

	return oldData, nil
}
