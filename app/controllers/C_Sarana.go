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

func IndexSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	saranas := []models.Sarana{}
	data := make(map[string]interface{})

	db.Find(&saranas)

	data["data"] = saranas

	helpers.JSON_(ctx, data)
}
func ShowByIdSarana(ctx *fasthttp.RequestCtx) {

	db := database.DB

	_id := ctx.UserValue("id")
	sarana := []models.Sarana{}
	data := make(map[string]interface{})

	db.Find(&sarana, _id)

	data["data"] = sarana

	helpers.JSON_(ctx, data)
}
func PostSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB

	postValues := ctx.PostBody()
	data := make(map[string]interface{})

	sarana := models.Sarana{}
	if err := json.Unmarshal(postValues, &sarana); err != nil {
		log.Println("Error UnMarshal", err)

		data["message"] = "Invalid JSON Field"
		data["error"] = err.Error()
		helpers.JSON_(ctx, data)
	}

	err := db.Create(&sarana).Error

	if err != nil {
		data["message"] = "Failed to save data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
func UpdateSarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	postValues := ctx.PostBody()

	data := make(map[string]interface{})

	sarana := models.Sarana{}
	if err := json.Unmarshal(postValues, &sarana); err != nil {
		data["message"] = "Invalid JSON"
		helpers.JSON_(ctx, data)
	}

	sarana_db := models.Sarana{}
	if err := db.Where("id = ?", sarana.ID).First(&sarana_db).Error; err != nil {
		fmt.Println("")
		data["message"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}
	sarana_db.Nama = sarana.Nama

	err := db.Save(&sarana).Error

	if err != nil {
		data["message"] = "Failed to updating data"
	} else {
		data["message"] = "success"
	}

	helpers.JSON_(ctx, data)
}
func DestroySarana(ctx *fasthttp.RequestCtx) {
	db := database.DB
	idValue := fmt.Sprintf("%v", ctx.UserValue("id"))
	sarana_id, err := strconv.Atoi(idValue)

	if err != nil {
		fmt.Println("Error Convert ID %v", err)
	}

	sarana := models.Sarana{}
	data := make(map[string]interface{})
	if db.Find(&sarana, sarana_id).RecordNotFound() {
		data["error"] = "Data Not Found"
		helpers.JSON_(ctx, data)
	}

	if err := db.Where("id = ? AND deleted_at IS NULL", sarana_id).Delete(&sarana).Error; err != nil {
		data["message"] = "Failed to deleting data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON_(ctx, data)
}
