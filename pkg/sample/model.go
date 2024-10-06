package sample

import (
	"time"
)

type Sample struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
