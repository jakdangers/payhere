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
	ID          int        `json:"id"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate"`
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
