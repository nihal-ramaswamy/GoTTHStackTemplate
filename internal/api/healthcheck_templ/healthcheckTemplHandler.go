package healthchecktempl_api

import (
	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/gotta_stack/internal/constants"
	healthcheck_templ "github.com/nihal-ramaswamy/gotta_stack/web/templates/healthcheck"
	layout_templ "github.com/nihal-ramaswamy/gotta_stack/web/templates/layout"
	"go.uber.org/zap"
)

type HealthCheckTemplHandler struct {
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewHealthCheckTemplHandler(log *zap.Logger) *HealthCheckTemplHandler {
	return &HealthCheckTemplHandler{
		middlewares: []gin.HandlerFunc{},
	}
}

func (*HealthCheckTemplHandler) Pattern() string {
	return "/healthcheck"
}

func (*HealthCheckTemplHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := healthcheck_templ.HealthCheck()
		layout_templ.Layout(c, "My App").Render(ctx, ctx.Writer)
	}
}

func (*HealthCheckTemplHandler) RequestMethod() string {
	return constants.GET
}

func (h *HealthCheckTemplHandler) Middlewares() []gin.HandlerFunc {
	return h.middlewares
}
