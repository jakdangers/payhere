package domain

import (
	"time"
)

type Base struct {
	ID          int
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate time.Time
}
