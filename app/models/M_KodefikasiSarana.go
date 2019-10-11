package models

import "time"

type KodefikasiSarana struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	Sarana       Sarana     `gorm:"foreignkey:SaranaID" json:"sarana"`
	SaranaID     int        `json:"sarana_id"`
	Kode         string     `gorm:"size:5;not null" json:"kode"`
	Nama         string     `gorm:"size:191;not null" json:"nama"`
	JumlahGandar int        `json:"jumlah_gandar"`
	JumlahBogie  int        `json:"jumlah_bogie"`
	Kelas        string     `json:"kelas"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `json:"-"`
}
