package attendauto

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         uint64    `json:"id"`
	DateAttend string    `json:"DateAttend"`
	ClockIn    time.Time `json:"ClockIn"`
	ClockOut   time.Time `json:"ClockOut"`
	TypeInOut  string    `json:"TypeInOut"`
	InPrint    bool      `json:"InPrint"`
	OutPrint   bool      `json:"OutPrint"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
