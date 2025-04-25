package service

import (
	"context"
	"sort"
	"ums/internal/domain/model"
	"ums/internal/repository"
)

type CourseService interface {
	GetCourses(ctx context.Context, strategy string) ([]*model.Course, error)
	CreateCourse(ctx context.Context, course *model.Course) error
	DeleteCourse(ctx context.Context, id int64) error
}

type courseService struct {
	repo repository.CourseRepository
}

type SortStrategy interface {
	Sort([]*model.Course) []*model.Course
}

type defaultSort struct{}
type dateSort struct{}
type enrollmentSort struct{}

func (s *defaultSort) Sort(courses []*model.Course) []*model.Course { return courses }

func (s *dateSort) Sort(courses []*model.Course) []*model.Course {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i].CreatedAt.After(courses[j].CreatedAt)
	})
	return courses
}

func (s *enrollmentSort) Sort(courses []*model.Course) []*model.Course {
	sort.Slice(courses, func(i, j int) bool {
		return len(courses[i].Enrollments) > len(courses[j].Enrollments)
	})
	return courses
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) GetCourses(ctx context.Context, strategy string) ([]*model.Course, error) {
	courses, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var sorter SortStrategy
	switch strategy {
	case "date":
		sorter = &dateSort{}
	case "enrollment":
		sorter = &enrollmentSort{}
	default:
		sorter = &defaultSort{}
	}

	return sorter.Sort(courses), nil
}

func (s *courseService) CreateCourse(ctx context.Context, course *model.Course) error {
	return s.repo.Create(ctx, course)
}

func (s *courseService) DeleteCourse(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
