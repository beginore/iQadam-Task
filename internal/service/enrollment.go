// service/enrollment_service.go
package service

import (
	"context"
	"errors"
	"ums/internal/domain/model"
	"ums/internal/repository"
)

type EnrollmentService interface {
	EnrollStudent(ctx context.Context, studentID, courseID int64) error
	UnenrollStudent(ctx context.Context, studentID, courseID int64) error
	GetEnrollmentsByStudent(ctx context.Context, studentID int64) ([]*model.Enrollment, error)
	GetEnrollmentsByCourse(ctx context.Context, courseID int64) ([]*model.Enrollment, error)
	GetAllEnrollments(ctx context.Context) ([]*model.Enrollment, error)
}

type enrollmentService struct {
	enrollmentRepo repository.EnrollmentRepository
	courseRepo     repository.CourseRepository
	userRepo       repository.UserRepository
}

func NewEnrollmentService(
	enrollmentRepo repository.EnrollmentRepository,
	courseRepo repository.CourseRepository,
	userRepo repository.UserRepository,
) EnrollmentService {
	return &enrollmentService{
		enrollmentRepo: enrollmentRepo,
		courseRepo:     courseRepo,
		userRepo:       userRepo,
	}
}

func (s *enrollmentService) EnrollStudent(ctx context.Context, studentID, courseID int64) error {
	if _, err := s.courseRepo.GetByID(ctx, courseID); err != nil {
		return errors.New("course not found")
	}

	user, err := s.userRepo.GetByID(ctx, studentID)
	if err != nil || user.Role != "STUDENT" {
		return errors.New("invalid student")
	}

	if exists, _ := s.enrollmentExists(ctx, studentID, courseID); exists {
		return errors.New("student already enrolled")
	}

	return s.enrollmentRepo.Create(ctx, studentID, courseID)
}

func (s *enrollmentService) UnenrollStudent(ctx context.Context, studentID, courseID int64) error {
	exists, err := s.enrollmentExists(ctx, studentID, courseID)
	if err != nil || !exists {
		return errors.New("enrollment not found")
	}
	return s.enrollmentRepo.Delete(ctx, studentID, courseID)
}

func (s *enrollmentService) enrollmentExists(ctx context.Context, studentID, courseID int64) (bool, error) {
	enrollments, err := s.enrollmentRepo.GetByStudent(ctx, studentID)
	if err != nil {
		return false, err
	}
	for _, e := range enrollments {
		if e.CourseID == courseID {
			return true, nil
		}
	}
	return false, nil
}

func (s *enrollmentService) GetEnrollmentsByStudent(ctx context.Context, studentID int64) ([]*model.Enrollment, error) {
	return s.enrollmentRepo.GetByStudent(ctx, studentID)
}

func (s *enrollmentService) GetEnrollmentsByCourse(ctx context.Context, courseID int64) ([]*model.Enrollment, error) {
	return s.enrollmentRepo.GetByCourse(ctx, courseID)
}

func (s *enrollmentService) GetAllEnrollments(ctx context.Context) ([]*model.Enrollment, error) {
	return s.enrollmentRepo.GetAll(ctx)
}
