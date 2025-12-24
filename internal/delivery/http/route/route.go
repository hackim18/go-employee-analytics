package route

import (
	"go-employee-analytics/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Router              *gin.Engine
	AnalyticsController *http.AnalyticsController
}

func (c *RouteConfig) Setup() {
	root := c.Router.Group("")

	c.RegisterAnalyticsRoutes(root)
	c.RegisterCommonRoutes(c.Router)
}
