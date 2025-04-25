package dto

type CourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	TeacherID   int64  `json:"teacher_id" binding:"required"`
}

type CourseResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TeacherID   int64  `json:"teacher_id"`
	CreatedAt   string `json:"created_at"`
}
