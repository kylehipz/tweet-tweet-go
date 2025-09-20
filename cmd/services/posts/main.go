package main

import (
	"github.com/kylehipz/tweet-tweet-go/config"
	"github.com/kylehipz/tweet-tweet-go/pkg/server"
	"github.com/labstack/echo/v4"
)


func main() {
	settings := config.GetSettings()

	apiServer := server.NewApiServer(echo.New(), settings["port"])

	apiServer.RegisterPostRoutes()

	apiServer.Start()
}
