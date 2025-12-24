package config

import (
	"go-employee-analytics/internal/delivery/http"
	"go-employee-analytics/internal/delivery/http/route"
	"go-employee-analytics/internal/repository"
	"go-employee-analytics/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Router   *gin.Engine
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	analyticsRepository := repository.NewAnalyticsRepository(config.DB)

	// setup use cases
	outputDir := config.Config.GetString("OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "output"
	}
	analyticsUseCase := usecase.NewAnalyticsUseCase(config.DB, config.Log, outputDir, analyticsRepository)

	// setup controller
	analyticsController := http.NewAnalyticsController(analyticsUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		Router:              config.Router,
		AnalyticsController: analyticsController,
	}
	routeConfig.Setup()
}
