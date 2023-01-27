package models

import (
	"time"
)

type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
