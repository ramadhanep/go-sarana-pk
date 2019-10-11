package models

import "time"

type IdentitasSarana struct {
	ID                  uint              `gorm:"primary_key" json:"id"`
	KlasifikasiSarana   KlasifikasiSarana `gorm:"foreignkey:KlasifikasiSaranaID" json:"klasifikasi_sarana"`
	KlasifikasiSaranaID int               `json:"klasifikasi_sarana_id"`
	Tahun               string            `gorm:"size:191;not null" json:"tahun"`
	NomorUrut           int               `gorm:"not null" json:"nomor_urut"`
	CreatedAt           time.Time         `json:"-"`
	UpdatedAt           time.Time         `json:"-"`
	DeletedAt           *time.Time        `json:"-"`
}
