package app

import (
	"github.com/aasimsajjad22/bookstore_utils-go/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application ...")
	router.Run(":8082")
}
