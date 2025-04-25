package repository

import (
	"context"
	"database/sql"
	"ums/internal/domain/model"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, studentID, courseID int64) error
	Delete(ctx context.Context, studentID, courseID int64) error
	GetByStudent(ctx context.Context, studentID int64) ([]*model.Enrollment, error)
	GetByCourse(ctx context.Context, courseID int64) ([]*model.Enrollment, error)
}

type enrollmentRepository struct {
	db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

func (r *enrollmentRepository) Create(ctx context.Context, studentID, courseID int64) error {
	query := `INSERT INTO enrollment (student_id, course_id) 
		VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, studentID, courseID)
	return err
}

func (r *enrollmentRepository) Delete(ctx context.Context, studentID, courseID int64) error {
	query := `DELETE FROM enrollment 
		WHERE student_id = $1 AND course_id = $2`
	_, err := r.db.ExecContext(ctx, query, studentID, courseID)
	return err
}

func (r *enrollmentRepository) GetByStudent(ctx context.Context, studentID int64) ([]*model.Enrollment, error) {
	query := `SELECT id, course_id FROM enrollment WHERE student_id = $1`
	rows, err := r.db.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []*model.Enrollment
	for rows.Next() {
		var e model.Enrollment
		e.StudentID = studentID
		err := rows.Scan(&e.ID, &e.CourseID)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, &e)
	}
	return enrollments, nil
}

func (r *enrollmentRepository) GetByCourse(ctx context.Context, courseID int64) ([]*model.Enrollment, error) {
	query := `SELECT id, student_id FROM enrollment WHERE course_id = $1`
	rows, err := r.db.QueryContext(ctx, query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []*model.Enrollment
	for rows.Next() {
		var e model.Enrollment
		e.CourseID = courseID
		err := rows.Scan(&e.ID, &e.StudentID)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, &e)
	}
	return enrollments, nil
}
