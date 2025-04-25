package service

import (
	"context"
	"errors"
	"ums/internal/repository"
)

type EnrollmentService interface {
	EnrollStudent(ctx context.Context, studentID, courseID int64) error
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

	return s.enrollmentRepo.Create(ctx, studentID, courseID)
}
