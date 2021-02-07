package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)


func main(){
	r := gin.Default()
	r.GET ("/unauth/fahrplanservice/v2/standortenearby/:lat/:lng/", standorteNearby)
	r.GET("/unauth/fahrplanservice/v2/abfahrtstabelle/:type/:from", abfahrtsTabelle)
	err := r.RunTLS(":8443", "./fake-backend/server.crt", "./fake-backend/server.key")
	if err != nil {
		logrus.Fatalf("unable to start web server: %v", err)
	}
}

func standorteNearby(context *gin.Context) {
	lat := context.Param("lat")
	lng := context.Param("lng")

	logrus.Infof("requested standorteNearby with %v, %v", lat, lng)

	// Load from static file
	f, err := os.Open("./resources/test/standorte-nearby.json")
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
	}

	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.Writer.WriteHeader(http.StatusOK)
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Writer.WriteHeaderNow()
	_, _ = context.Writer.Write(fileContent)
}

func abfahrtsTabelle(context *gin.Context) {
	paramType := context.Param("type")
	paramFrom := context.Param("from")

	logrus.Infof("requested abfahrtsTabelle with %v, %v: %v", paramType, paramFrom, context.Request.URL.Query())

	// Load from static file
	f, err := os.Open("./resources/test/abfahrt.json")
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
	}

	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.Writer.WriteHeader(http.StatusOK)
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Writer.WriteHeaderNow()
	_, _ = context.Writer.Write(fileContent)
}