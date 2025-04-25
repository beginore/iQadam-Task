package controller

import (
	"net/http"
	"strconv"
	"ums/internal/domain/model"
	"ums/internal/dto"
	"ums/internal/service"

	"github.com/gin-gonic/gin"
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

func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := c.service.DeleteCourse(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
