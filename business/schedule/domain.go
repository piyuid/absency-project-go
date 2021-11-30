package schedule

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint64    `json:"id"`
	DateSchedule time.Time `json:"DateSchedule"`
	InSchedule   time.Time `json:"InSchedule"`
	OutSchedule  time.Time `json:"OutSchedule"`
	Category     string    `json:"Category"`
	CategoryDesc string    `json:"CategoryDesc"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
