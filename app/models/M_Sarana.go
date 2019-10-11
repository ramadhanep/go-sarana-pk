package models

import "time"

type Sarana struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Nama      string     `gorm:"size:191;not null" json:"nama"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
