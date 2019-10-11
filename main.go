package main

import (
	"saranapk/app/models"
	"saranapk/config/database"
	"saranapk/config/routes"
)

func main() {
	database.DB.Debug().AutoMigrate(
		&models.Sarana{},
		&models.KodefikasiSarana{},
		&models.KlasifikasiSarana{},
		&models.IdentitasSarana{},
	)
	routes.Setup()
}
