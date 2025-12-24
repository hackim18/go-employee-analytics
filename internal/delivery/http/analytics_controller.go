package http

import (
	"encoding/json"
	"go-employee-analytics/internal/constants"
	"go-employee-analytics/internal/model"
	"go-employee-analytics/internal/usecase"
	"go-employee-analytics/internal/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AnalyticsController struct {
	Log      *logrus.Logger
	UseCase  *usecase.AnalyticsUseCase
	Validate *validator.Validate
}

func NewAnalyticsController(useCase *usecase.AnalyticsUseCase, logger *logrus.Logger, validate *validator.Validate) *AnalyticsController {
	return &AnalyticsController{
		Log:      logger,
		UseCase:  useCase,
		Validate: validate,
	}
}

func (c *AnalyticsController) Health(ctx *gin.Context) {
	payload := map[string]string{"status": "ok"}
	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsHealthSuccess, payload)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q2(ctx *gin.Context) {
	results, err := c.UseCase.ActiveSmithEmployees(ctx.Request.Context())
	if err != nil {
		c.Log.Warnf("Failed to fetch Q2 data : %+v", err)
		utils.HandleHTTPError(ctx, err)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ2Success, results)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q3(ctx *gin.Context) {
	results, err := c.UseCase.EmployeesNoReviews(ctx.Request.Context())
	if err != nil {
		c.Log.Warnf("Failed to fetch Q3 data : %+v", err)
		utils.HandleHTTPError(ctx, err)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ3Success, results)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q4(ctx *gin.Context) {
	result, err := c.UseCase.HireDateDiff(ctx.Request.Context())
	if err != nil {
		c.Log.Warnf("Failed to fetch Q4 data : %+v", err)
		utils.HandleHTTPError(ctx, err)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ4Success, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q5(ctx *gin.Context) {
	results, err := c.UseCase.SalaryProjections(ctx.Request.Context())
	if err != nil {
		c.Log.Warnf("Failed to fetch Q5 data : %+v", err)
		utils.HandleHTTPError(ctx, err)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ5Success, results)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q6Save(ctx *gin.Context) {
	files, err := c.UseCase.SaveOutputs(ctx.Request.Context())
	if err != nil {
		c.Log.Warnf("Failed to save output files : %+v", err)
		utils.HandleHTTPError(ctx, err)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ6Success, files)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q7Read(ctx *gin.Context) {
	request := new(model.FileQueryRequest)
	if err := ctx.ShouldBindQuery(request); err != nil {
		c.Log.Warnf("Failed to parse request query : %+v", err)
		utils.HandleHTTPError(ctx, utils.Error(constants.FailedInputFormat, http.StatusBadRequest, err))
		return
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Validation failed : %+v", err)
		message := utils.TranslateValidationError(c.Validate, err)
		utils.HandleHTTPError(ctx, utils.Error(message, http.StatusBadRequest, err))
		return
	}

	data, err := c.UseCase.ReadJSONFile(request.Name)
	if err != nil {
		if os.IsNotExist(err) {
			c.Log.Warnf("Requested file not found: %s", request.Name)
			utils.HandleHTTPError(ctx, model.ErrNotFound)
			return
		}
		c.Log.Warnf("Failed to read output file : %+v", err)
		utils.HandleHTTPError(ctx, err)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ7Success, json.RawMessage(data))
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q8(ctx *gin.Context) {
	request := new(model.CityQueryRequest)
	if err := ctx.ShouldBindQuery(request); err != nil {
		c.Log.Warnf("Failed to parse request query : %+v", err)
		utils.HandleHTTPError(ctx, utils.Error(constants.FailedInputFormat, http.StatusBadRequest, err))
		return
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Validation failed : %+v", err)
		message := utils.TranslateValidationError(c.Validate, err)
		utils.HandleHTTPError(ctx, utils.Error(message, http.StatusBadRequest, err))
		return
	}

	result := c.UseCase.CityLookup(request.City)
	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ8Success, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q9SortedUnique(ctx *gin.Context) {
	results := c.UseCase.SortedUniqueNumbers()
	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ9ASuccess, results)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q9Duplicates(ctx *gin.Context) {
	results := c.UseCase.DuplicateCounts()
	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ9BSuccess, results)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q9Remove(ctx *gin.Context) {
	request := new(model.RemoveNumbersRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		utils.HandleHTTPError(ctx, utils.Error(constants.FailedDataFromBody, http.StatusBadRequest, err))
		return
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Validation failed : %+v", err)
		message := utils.TranslateValidationError(c.Validate, err)
		utils.HandleHTTPError(ctx, utils.Error(message, http.StatusBadRequest, err))
		return
	}

	results := c.UseCase.RemoveNumbers(request.Remove)
	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ9CSuccess, results)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q9Add(ctx *gin.Context) {
	request := new(model.AddNumbersRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		utils.HandleHTTPError(ctx, utils.Error(constants.FailedDataFromBody, http.StatusBadRequest, err))
		return
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Validation failed : %+v", err)
		message := utils.TranslateValidationError(c.Validate, err)
		utils.HandleHTTPError(ctx, utils.Error(message, http.StatusBadRequest, err))
		return
	}

	results := c.UseCase.AddWithCap(request.Add)
	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ9DSuccess, results)
	ctx.JSON(http.StatusOK, res)
}

func (c *AnalyticsController) Q10(ctx *gin.Context) {
	report := c.UseCase.GenerateRandomReport()
	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AnalyticsQ10Success, report)
	ctx.JSON(http.StatusOK, res)
}
