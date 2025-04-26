package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ums/internal/domain/model"
	"ums/internal/dto"
	"ums/internal/service"
)

type CourseController struct {
	service service.CourseService
}

func NewCourseController(service service.CourseService) *CourseController {
	return &CourseController{service: service}
}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var req dto.CourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := model.Course{
		Title:       req.Title,
		Description: req.Description,
		TeacherID:   &req.TeacherID,
	}

	if err := c.service.CreateCourse(ctx, &course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, course)
}

func (c *CourseController) GetCourses(ctx *gin.Context) {
	strategy := ctx.DefaultQuery("sort", "default")
	courses, err := c.service.GetCourses(ctx, strategy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	var req dto.UpdateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем существующий курс
	existing, err := c.service.GetCourseByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	}

	// Обновляем только переданные поля
	if req.Title != nil {
		existing.Title = *req.Title
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.TeacherID != nil {
		existing.TeacherID = req.TeacherID
	}

	if err := c.service.UpdateCourse(ctx, existing); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToCourseResponse(existing))
}

func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := c.service.DeleteCourse(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
