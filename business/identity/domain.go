package identity

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Birthday      time.Time `json:"birthday"`
	Status        string    `json:"status"`
	Gender        string    `json:"gender"`
	Npwp          string    `json:"npwp"`
	Ktp           uint16    `json:"ktp"`
	Bpjs          string    `json:"bpjs"`
	HpEmployee    string    `json:"HpEmployee`
	NameEmergency string    `json:"NameEmergency"`
	HpEmergency   string    `json:"HpEmergency`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
