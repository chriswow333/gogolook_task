package main

import (
	"os"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"

	"gogolook_task/routes"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	logrus.SetReportCaller(true)

}

func main() {

	container := routes.BuildContainer()

	if err := container.Invoke(func(router *gin.Engine) {
		logrus.WithFields(logrus.Fields{}).Info("start serving http request")

		manners.ListenAndServe(":8080", router)
	}); err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("container.Invoke failed: %s", err)
		panic(err)
	}
}
