package main

import (
	"github.com/learning-drops-api/app/controllers"
	"github.com/learning-drops-api/config"
	"github.com/learning-drops-api/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StartWebServer()
}
