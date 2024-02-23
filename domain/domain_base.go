package domain

import (
	"database/sql"
	"time"
)

type Base struct {
	ID          int
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate sql.NullTime
}
