package domain

import (
	"net/http"
	"time"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PayhereResponse struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type BaseDTO struct {
	ID         int       `json:"id" validate:"required" example:"1"`
	CreateDate time.Time `json:"createDate" validate:"required" example:"2024-02-28T15:04:05Z"`
	UpdateDate time.Time `json:"updateDate" validate:"required" example:"2024-02-28T15:04:05Z"`
}

func PayhereResponseFrom(code int, data any) (int, PayhereResponse) {
	return code, PayhereResponse{
		Meta: Meta{
			Code:    code,
			Message: http.StatusText(code),
		},
		Data: data,
	}
}
