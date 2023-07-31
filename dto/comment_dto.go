package dto

import (
	model "github.com/maibokkrub/simple-backend/models"
	"gopkg.in/go-playground/validator.v9"
)

type CreateCommentDTO struct {
	AppointmentId uint   `json:"appointmentId" validate:"required"`
	Comment       string `json:"comment" validate:"required"`
	CreatedBy     uint   `json:"creator" validate:"required"`
}

func (d *CreateCommentDTO) ToModel() (*model.Comment, error) {
	v := validator.New()

	if err := v.Struct(d); err != nil {
		// todo: clean up error message leak
		return nil, err
	}

	return &model.Comment{
		AppointmentID: d.AppointmentId,
		Body:          d.Comment,
		UserID:        d.CreatedBy,
	}, nil
}
