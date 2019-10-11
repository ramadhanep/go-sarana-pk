package helpers

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func JSON_(ctx *fasthttp.RequestCtx, data map[string]interface{}) {
	ctx.Response.Header.Set("content-type", "application/json")

	res, err := json.Marshal(data)

	if err != nil {
		log.Println("Error Convert to JSON")
		data["error"] = err
	}

	ctx.Write(res)
	ctx.SetStatusCode(fasthttp.StatusOK)
}
