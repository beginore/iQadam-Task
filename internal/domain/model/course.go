package model

import "time"

type Course struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	TeacherID   *int64       `json:"teacher_id,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	Enrollments []Enrollment `json:"enrollments,omitempty"`
}
