package main

import (
	"os"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"gogolook_task/routes"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	log.SetReportCaller(true)

}

func main() {

	container := routes.BuildContainer()

	if err := container.Invoke(func(router *gin.Engine) {
		log.WithFields(log.Fields{}).Info("start serving http request")

		manners.ListenAndServe(":8080", router)
	}); err != nil {
		log.WithFields(log.Fields{}).Errorf("container.Invoke failed: %s", err)
		panic(err)
	}
}
