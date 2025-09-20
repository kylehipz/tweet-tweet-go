package server

import (
	"fmt"

	"github.com/kylehipz/tweet-tweet-go/pkg/routes"
	"github.com/labstack/echo/v4"
)


type ApiServer struct {
	e *echo.Echo
	port string
}

func (a *ApiServer) Start() {
	a.e.Start(fmt.Sprintf(":%s", a.port))
}

func (a *ApiServer) RegisterUserRoutes() {
	routes.RegisterUserRoutes(a.e)
}

func (a *ApiServer) RegisterPostRoutes() {
	routes.RegisterUserRoutes(a.e)
}

func (a *ApiServer) RegisterFollowRoutes() {
	routes.RegisterUserRoutes(a.e)
}

func (a *ApiServer) RegisterTimelineRoutes() {
	routes.RegisterUserRoutes(a.e)
}

func (a *ApiServer) RegisterNotificationRoutes() {
	routes.RegisterUserRoutes(a.e)
}

func (a *ApiServer) RegisterSearchRoutes() {
	routes.RegisterUserRoutes(a.e)
}

func NewApiServer(e *echo.Echo, port string) *ApiServer {
	return &ApiServer{e, port}
}

