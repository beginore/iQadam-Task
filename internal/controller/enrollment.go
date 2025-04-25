package controller

import (
	"net/http"
	"ums/internal/dto"
	"ums/internal/service"

	"github.com/gin-gonic/gin"
)

type EnrollmentController struct {
	service service.EnrollmentService
}

func NewEnrollmentController(service service.EnrollmentService) *EnrollmentController {
	return &EnrollmentController{service: service}
}

func (c *EnrollmentController) EnrollStudent(ctx *gin.Context) {
	var req dto.EnrollmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if err := c.service.EnrollStudent(ctx, req.StudentID, req.CourseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
