package healthcheck_api

import (
	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/gotta_stack/internal/interfaces"
	"go.uber.org/zap"
)

type HealthCheckGroup struct {
	routeHandlers []interfaces.HandlerInterface
	middlewares   []gin.HandlerFunc
}

func (*HealthCheckGroup) Group() string {
	return "/healthcheck"
}

func (h *HealthCheckGroup) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewHealthCheckGroup(log *zap.Logger) *HealthCheckGroup {
	handlers := []interfaces.HandlerInterface{
		NewHealthCheckHandler(log),
	}

	return &HealthCheckGroup{
		routeHandlers: handlers,
		middlewares:   []gin.HandlerFunc{},
	}
}

func (*HealthCheckGroup) AuthRequired() bool {
	return false
}

func (h *HealthCheckGroup) Middlewares() []gin.HandlerFunc {
	return h.middlewares
}
