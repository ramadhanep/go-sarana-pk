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

func IndexIdentitasSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	identitasSaranas := []models.IdentitasSarana{}
	data := make(map[string]interface{})

	db.Preload("KlasifikasiSarana.KodefikasiSarana.Sarana").Find(&identitasSaranas)

	data["data"] = identitasSaranas

	helpers.JSON_(ctx, data)
}
func ShowByIdIdentitasSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	_id := ctx.UserValue("id")
	identitasSarana := []models.IdentitasSarana{}
	data := make(map[string]interface{})

	db.Preload("KlasifikasiSarana.KodefikasiSarana.Sarana").Find(&identitasSarana, _id)

	data["data"] = identitasSarana

	helpers.JSON_(ctx, data)
}
func PostIdentitasSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB

	postValues := ctx.PostBody()
	data := make(map[string]interface{})

	identitasSarana := models.IdentitasSarana{}
	if err := json.Unmarshal(postValues, &identitasSarana); err != nil {
		log.Println("Error UnMarshal", err)

		data["message"] = "Invalid JSON Field"
		data["error"] = err.Error()
		helpers.JSON_(ctx, data)
	}

	err := db.Create(&identitasSarana).Error

	if err != nil {
		data["message"] = "Failed to save data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
func UpdateIdentitasSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	postValues := ctx.PostBody()

	data := make(map[string]interface{})

	identitasSarana := models.IdentitasSarana{}
	if err := json.Unmarshal(postValues, &identitasSarana); err != nil {
		data["message"] = "Invalid JSON"
		helpers.JSON_(ctx, data)
	}

	identitasSarana_db := models.IdentitasSarana{}
	if err := db.Where("id = ?", identitasSarana.ID).First(&identitasSarana_db).Error; err != nil {
		fmt.Println("")
		data["message"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}
	identitasSarana_db.KlasifikasiSaranaID = identitasSarana.KlasifikasiSaranaID
	identitasSarana_db.Tahun = identitasSarana.Tahun
	identitasSarana_db.NomorUrut = identitasSarana.NomorUrut

	err := db.Save(&identitasSarana).Error

	if err != nil {
		data["message"] = "Failed to updating data"
	} else {
		data["message"] = "success"
	}

	helpers.JSON_(ctx, data)
}
func DestroyIdentitasSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	idValue := fmt.Sprintf("%v", ctx.UserValue("id"))
	identitasSarana_id, err := strconv.Atoi(idValue)

	if err != nil {
		fmt.Println("Error Convert ID %v", err)
	}

	identitasSarana := models.IdentitasSarana{}
	data := make(map[string]interface{})
	if db.Find(&identitasSarana, identitasSarana_id).RecordNotFound() {
		data["error"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", identitasSarana_id).Delete(&identitasSarana).Error; err != nil {
		data["message"] = "Failed to deleting data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
