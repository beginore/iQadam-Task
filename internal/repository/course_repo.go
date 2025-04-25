package repository

import (
	"context"
	"database/sql"
	"ums/internal/domain/model"
)

type CourseRepository interface {
	Create(ctx context.Context, course *model.Course) error
	GetByID(ctx context.Context, id int64) (*model.Course, error)
	Update(ctx context.Context, course *model.Course) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*model.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(ctx context.Context, course *model.Course) error {
	query := `INSERT INTO courses (title, description, teacher_id) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query,
		course.Title,
		course.Description,
		course.TeacherID,
	).Scan(&course.ID, &course.CreatedAt)
}

func (r *courseRepository) GetByID(ctx context.Context, id int64) (*model.Course, error) {
	query := `SELECT id, title, description, teacher_id, created_at 
		FROM courses WHERE id = $1`
	course := &model.Course{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.TeacherID,
		&course.CreatedAt,
	)
	return course, err
}

func (r *courseRepository) Update(ctx context.Context, course *model.Course) error {
	query := `UPDATE courses 
		SET title = $1, description = $2, teacher_id = $3 
		WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query,
		course.Title,
		course.Description,
		course.TeacherID,
		course.ID,
	)
	return err
}

func (r *courseRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM courses WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *courseRepository) GetAll(ctx context.Context) ([]*model.Course, error) {
	query := `SELECT id, title, description, teacher_id, created_at FROM courses`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		var course model.Course
		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.TeacherID,
			&course.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}
