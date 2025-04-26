package dto

import (
	"time"
	"ums/internal/domain/model"
)

type CourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	TeacherID   int64  `json:"teacher_id" binding:"required"`
}

type UpdateCourseRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	TeacherID   *int64  `json:"teacher_id,omitempty"`
}

type CourseResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TeacherID   int64     `json:"teacher_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func ToCourseResponse(course *model.Course) CourseResponse {
	return CourseResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		TeacherID:   *course.TeacherID,
		CreatedAt:   course.CreatedAt,
	}
}
