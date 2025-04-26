package service

import (
	"context"
	"sort"
	"ums/internal/domain/model"
	"ums/internal/repository"
)

type CourseService interface {
	CreateCourse(ctx context.Context, course *model.Course) error
	GetCourses(ctx context.Context, sortBy string) ([]*model.Course, error)
	UpdateCourse(ctx context.Context, course *model.Course) error
	DeleteCourse(ctx context.Context, id int64) error
	GetCourseByID(ctx context.Context, id int64) (*model.Course, error)
}

type courseService struct {
	repo repository.CourseRepository
}

type SortStrategy interface {
	Sort([]*model.Course)
}

type DateSort struct{}

func (s *DateSort) Sort(courses []*model.Course) {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i].CreatedAt.After(courses[j].CreatedAt)
	})
}

func (s *courseService) GetCourses(ctx context.Context, sortBy string) ([]*model.Course, error) {
	courses, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var strategy SortStrategy
	if sortBy == "date" {
		strategy = &DateSort{}
	}

	if strategy != nil {
		strategy.Sort(courses)
	}

	return courses, nil
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(ctx context.Context, course *model.Course) error {
	return s.repo.Create(ctx, course)
}

func (s *courseService) UpdateCourse(ctx context.Context, course *model.Course) error {
	return s.repo.Update(ctx, course)
}

func (s *courseService) DeleteCourse(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *courseService) GetCourseByID(ctx context.Context, id int64) (*model.Course, error) {
	return s.repo.GetByID(ctx, id)
}
