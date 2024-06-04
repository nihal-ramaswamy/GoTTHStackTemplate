package routes

import (
	"github.com/gin-gonic/gin"
	healthcheck_api "github.com/nihal-ramaswamy/gotta_stack/internal/api/healthcheck"
	healthchecktempl_api "github.com/nihal-ramaswamy/gotta_stack/internal/api/healthcheck_templ"
	"github.com/nihal-ramaswamy/gotta_stack/internal/constants"
	"github.com/nihal-ramaswamy/gotta_stack/internal/interfaces"
	"go.uber.org/zap"
)

func NewRoutes(
	server *gin.Engine,
	log *zap.Logger,
) {
	serverGroupHandlers := []interfaces.ServerGroupInterface{
		healthcheck_api.NewHealthCheckGroup(log),
		healthchecktempl_api.NewHealthCheckTemplGroup(log),
	}

	for _, serverGroupHandler := range serverGroupHandlers {
		newGroup(server, serverGroupHandler)
	}
}

func newGroup(server *gin.Engine, groupHandler interfaces.ServerGroupInterface) {
	group := server.Group(groupHandler.Group(), groupHandler.Middlewares()...)
	{
		for _, route := range groupHandler.RouteHandlers() {
			newRoute(group, route)
		}
	}
}

func newRoute(server *gin.RouterGroup, routeHandler interfaces.HandlerInterface) {
	middlewares := routeHandler.Middlewares()
	middlewares = append(middlewares, routeHandler.Handler())
	switch routeHandler.RequestMethod() {
	case constants.GET:
		server.GET(routeHandler.Pattern(), middlewares...)
	case constants.POST:
		server.POST(routeHandler.Pattern(), middlewares...)
	}
}
