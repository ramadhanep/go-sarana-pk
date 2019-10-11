package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"saranapk/app/helpers"
	"saranapk/app/models"
	"saranapk/config/database"
	"strconv"

	"github.com/valyala/fasthttp"
)

func IndexKlasifikasiSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	klasifikasiSaranas := []models.KlasifikasiSarana{}
	data := make(map[string]interface{})

	db.Preload("KodefikasiSarana.Sarana").Find(&klasifikasiSaranas)

	data["data"] = klasifikasiSaranas

	helpers.JSON_(ctx, data)
}
func ShowByIdKlasifikasiSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	_id := ctx.UserValue("id")
	klasifikasiSarana := []models.KlasifikasiSarana{}
	data := make(map[string]interface{})

	db.Preload("KodefikasiSarana.Sarana").Find(&klasifikasiSarana, _id)

	data["data"] = klasifikasiSarana

	helpers.JSON_(ctx, data)
}
func PostKlasifikasiSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB

	postValues := ctx.PostBody()
	data := make(map[string]interface{})

	klasifikasiSarana := models.KlasifikasiSarana{}
	if err := json.Unmarshal(postValues, &klasifikasiSarana); err != nil {
		log.Println("Error UnMarshal", err)

		data["message"] = "Invalid JSON Field"
		data["error"] = err.Error()
		helpers.JSON_(ctx, data)
	}

	err := db.Create(&klasifikasiSarana).Error

	if err != nil {
		data["message"] = "Failed to save data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
func UpdateKlasifikasiSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	postValues := ctx.PostBody()

	data := make(map[string]interface{})

	klasifikasiSarana := models.KlasifikasiSarana{}
	if err := json.Unmarshal(postValues, &klasifikasiSarana); err != nil {
		data["message"] = "Invalid JSON"
		helpers.JSON_(ctx, data)
	}

	klasifikasiSarana_db := models.KlasifikasiSarana{}
	if err := db.Where("id = ?", klasifikasiSarana.ID).First(&klasifikasiSarana_db).Error; err != nil {
		fmt.Println("")
		data["message"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}
	klasifikasiSarana_db.KodefikasiSaranaID = klasifikasiSarana.KodefikasiSaranaID
	klasifikasiSarana_db.Kode = klasifikasiSarana.Kode
	klasifikasiSarana_db.Nama = klasifikasiSarana.Nama
	klasifikasiSarana_db.SeriTipe = klasifikasiSarana.SeriTipe

	err := db.Save(&klasifikasiSarana).Error

	if err != nil {
		data["message"] = "Failed to updating data"
	} else {
		data["message"] = "success"
	}

	helpers.JSON_(ctx, data)
}
func DestroyKlasifikasiSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	idValue := fmt.Sprintf("%v", ctx.UserValue("id"))
	klasifikasiSarana_id, err := strconv.Atoi(idValue)

	if err != nil {
		fmt.Println("Error Convert ID %v", err)
	}

	klasifikasiSarana := models.KlasifikasiSarana{}
	data := make(map[string]interface{})
	if db.Find(&klasifikasiSarana, klasifikasiSarana_id).RecordNotFound() {
		data["error"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", klasifikasiSarana_id).Delete(&klasifikasiSarana).Error; err != nil {
		data["message"] = "Failed to deleting data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
