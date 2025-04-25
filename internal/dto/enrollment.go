package dto

type EnrollmentRequest struct {
	StudentID int64 `json:"student_id" binding:"required"`
	CourseID  int64 `json:"course_id" binding:"required"`
}
