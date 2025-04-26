package dto

type EnrollmentRequest struct {
	StudentID int64 `json:"student_id" binding:"required"`
	CourseID  int64 `json:"course_id" binding:"required"`
}

type EnrollmentResponse struct {
	ID        int64 `json:"id"`
	StudentID int64 `json:"student_id"`
	CourseID  int64 `json:"course_id"`
}
