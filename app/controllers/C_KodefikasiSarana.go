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

func IndexKodefikasiSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	kodefikasiSaranas := []models.KodefikasiSarana{}
	data := make(map[string]interface{})

	db.Preload("Sarana").Find(&kodefikasiSaranas)

	data["data"] = kodefikasiSaranas

	helpers.JSON_(ctx, data)
}
func ShowByIdKodefikasiSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	_id := ctx.UserValue("id")
	kodefikasiSarana := []models.KodefikasiSarana{}
	data := make(map[string]interface{})

	db.Preload("Sarana").Find(&kodefikasiSarana, _id)

	data["data"] = kodefikasiSarana

	helpers.JSON_(ctx, data)
}
func PostKodefikasiSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB

	postValues := ctx.PostBody()
	data := make(map[string]interface{})

	kodefikasiSarana := models.KodefikasiSarana{}
	if err := json.Unmarshal(postValues, &kodefikasiSarana); err != nil {
		log.Println("Error UnMarshal", err)

		data["message"] = "Invalid JSON Field"
		data["error"] = err.Error()
		helpers.JSON_(ctx, data)
	}

	err := db.Create(&kodefikasiSarana).Error

	if err != nil {
		data["message"] = "Failed to save data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
func UpdateKodefikasiSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	postValues := ctx.PostBody()

	data := make(map[string]interface{})

	kodefikasiSarana := models.KodefikasiSarana{}
	if err := json.Unmarshal(postValues, &kodefikasiSarana); err != nil {
		data["message"] = "Invalid JSON"
		helpers.JSON_(ctx, data)
	}

	kodefikasiSarana_db := models.KodefikasiSarana{}
	if err := db.Where("id = ?", kodefikasiSarana.ID).First(&kodefikasiSarana_db).Error; err != nil {
		fmt.Println("")
		data["message"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}
	kodefikasiSarana_db.SaranaID = kodefikasiSarana.SaranaID
	kodefikasiSarana_db.Kode = kodefikasiSarana.Kode
	kodefikasiSarana_db.Nama = kodefikasiSarana.Nama
	kodefikasiSarana_db.JumlahGandar = kodefikasiSarana.JumlahGandar
	kodefikasiSarana_db.JumlahBogie = kodefikasiSarana.JumlahBogie
	kodefikasiSarana_db.Kelas = kodefikasiSarana.Kelas

	err := db.Save(&kodefikasiSarana).Error

	if err != nil {
		data["message"] = "Failed to updating data"
	} else {
		data["message"] = "success"
	}

	helpers.JSON_(ctx, data)
}
func DestroyKodefikasiSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	idValue := fmt.Sprintf("%v", ctx.UserValue("id"))
	kodefikasiSarana_id, err := strconv.Atoi(idValue)

	if err != nil {
		fmt.Println("Error Convert ID %v", err)
	}

	kodefikasiSarana := models.KodefikasiSarana{}
	data := make(map[string]interface{})
	if db.Find(&kodefikasiSarana, kodefikasiSarana_id).RecordNotFound() {
		data["error"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", kodefikasiSarana_id).Delete(&kodefikasiSarana).Error; err != nil {
		data["message"] = "Failed to deleting data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
