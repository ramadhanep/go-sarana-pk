package models

import "time"

type KlasifikasiSarana struct {
	ID                 uint             `gorm:"primary_key" json:"id"`
	KodefikasiSarana   KodefikasiSarana `gorm:"foreignkey:KodefikasiSaranaID" json:"kodefikasi_sarana"`
	KodefikasiSaranaID int              `json:"kodefikasi_sarana_id"`
	Kode               string           `gorm:"size:5;not null" json:"kode"`
	Nama               string           `gorm:"size:191;not null" json:"nama"`
	SeriTipe           int              `json:"seri_tipe"`
	CreatedAt          time.Time        `json:"-"`
	UpdatedAt          time.Time        `json:"-"`
	DeletedAt          *time.Time       `json:"-"`
}
