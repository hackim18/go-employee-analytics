package route

import "github.com/gin-gonic/gin"

func (c *RouteConfig) RegisterAnalyticsRoutes(rg *gin.RouterGroup) {
	analytics := rg.Group("")

	analytics.GET("/health", c.AnalyticsController.Health)
	analytics.GET("/q2", c.AnalyticsController.Q2)
	analytics.GET("/q3", c.AnalyticsController.Q3)
	analytics.GET("/q4", c.AnalyticsController.Q4)
	analytics.GET("/q5", c.AnalyticsController.Q5)
	analytics.POST("/q6/save", c.AnalyticsController.Q6Save)
	analytics.GET("/q7/file", c.AnalyticsController.Q7Read)
	analytics.GET("/q8", c.AnalyticsController.Q8)
	analytics.GET("/q9/sorted-unique", c.AnalyticsController.Q9SortedUnique)
	analytics.GET("/q9/duplicates", c.AnalyticsController.Q9Duplicates)
	analytics.POST("/q9/remove", c.AnalyticsController.Q9Remove)
	analytics.POST("/q9/add", c.AnalyticsController.Q9Add)
	analytics.GET("/q10", c.AnalyticsController.Q10)
}
