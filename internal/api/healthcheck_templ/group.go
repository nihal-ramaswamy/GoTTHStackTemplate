package healthchecktempl_api

import (
	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/gotta_stack/internal/interfaces"
	htmx_middleware "github.com/nihal-ramaswamy/gotta_stack/internal/middleware/htmx"
	"go.uber.org/zap"
)

type HealthCheckTemplGroup struct {
	routeHandlers []interfaces.HandlerInterface
	middlewares   []gin.HandlerFunc
}

func (*HealthCheckTemplGroup) Group() string {
	return "/templ"
}

func (h *HealthCheckTemplGroup) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewHealthCheckTemplGroup(log *zap.Logger) *HealthCheckTemplGroup {
	handlers := []interfaces.HandlerInterface{
		NewHealthCheckTemplHandler(log),
	}

	return &HealthCheckTemplGroup{
		routeHandlers: handlers,
		middlewares: []gin.HandlerFunc{
			htmx_middleware.TextHtmlMiddleware(),
		},
	}
}

func (*HealthCheckTemplGroup) AuthRequired() bool {
	return false
}

func (h *HealthCheckTemplGroup) Middlewares() []gin.HandlerFunc {
	return h.middlewares
}
