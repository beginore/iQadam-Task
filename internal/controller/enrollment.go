package controller

import (
	"net/http"
	"strconv"
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

func (c *EnrollmentController) UnenrollStudent(ctx *gin.Context) {
	var req dto.EnrollmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if err := c.service.UnenrollStudent(ctx, req.StudentID, req.CourseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *EnrollmentController) GetEnrollmentsByStudent(ctx *gin.Context) {
	studentID, _ := strconv.ParseInt(ctx.Param("student_id"), 10, 64)
	enrollments, err := c.service.GetEnrollmentsByStudent(ctx, studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, enrollments)
}

func (c *EnrollmentController) GetEnrollmentsByCourse(ctx *gin.Context) {
	courseID, _ := strconv.ParseInt(ctx.Param("course_id"), 10, 64)
	enrollments, err := c.service.GetEnrollmentsByCourse(ctx, courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, enrollments)
}

func (c *EnrollmentController) GetAllEnrollments(ctx *gin.Context) {
	enrollments, err := c.service.GetAllEnrollments(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, enrollments)
}
