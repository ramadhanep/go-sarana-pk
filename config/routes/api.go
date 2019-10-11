package routes

import (
	"log"
	"saranapk/app/controllers"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Setup() {
	router := fasthttprouter.New()

	// SARANA
	router.GET("/sarana", controllers.IndexSarana)
	router.GET("/sarana/:id", controllers.ShowByIdSarana)
	router.POST("/sarana/", controllers.PostSarana)
	router.PUT("/sarana", controllers.UpdateSarana)
	router.DELETE("/sarana/:id", controllers.DestroySarana)

	// KODEFIKASI SARANA
	router.GET("/kodefikasiSarana", controllers.IndexKodefikasiSarana)
	router.GET("/kodefikasiSarana/:id", controllers.ShowByIdKodefikasiSarana)
	router.POST("/kodefikasiSarana/", controllers.PostKodefikasiSarana)
	router.PUT("/kodefikasiSarana", controllers.UpdateKodefikasiSarana)
	router.DELETE("/kodefikasiSarana/:id", controllers.DestroyKodefikasiSarana)

	// KLASIFIKASI SARANA
	router.GET("/klasifikasiSarana", controllers.IndexKlasifikasiSarana)
	router.GET("/klasifikasiSarana/:id", controllers.ShowByIdKlasifikasiSarana)
	router.POST("/klasifikasiSarana/", controllers.PostKlasifikasiSarana)
	router.PUT("/klasifikasiSarana", controllers.UpdateKlasifikasiSarana)
	router.DELETE("/klasifikasiSarana/:id", controllers.DestroyKlasifikasiSarana)

	// IDENTITAS SARANA
	router.GET("/identitasSarana", controllers.IndexIdentitasSarana)
	router.GET("/identitasSarana/:id", controllers.ShowByIdIdentitasSarana)
	router.POST("/identitasSarana/", controllers.PostIdentitasSarana)
	router.PUT("/identitasSarana", controllers.UpdateIdentitasSarana)
	router.DELETE("/identitasSarana/:id", controllers.DestroyIdentitasSarana)

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8081"},
		AllowedHeaders:   []string{"x-something-client", "content-type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: false,
		AllowMaxAge:      5600,
		Debug:            true,
	})

	// =======================================================================
	// RUNNING SERVER
	listenAddr := ":8000"
	if err := fasthttp.ListenAndServe(listenAddr, withCors.CorsMiddleware(router.Handler)); err != nil {
		log.Fatal("Error in Listen And Serve : %v", err)
	}
	// END RUNNING SERVER
	// =======================================================================
}
