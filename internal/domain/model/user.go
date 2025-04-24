package model

type Student struct {
	ID      int64
	Name    string
	Surname string
	Courses []Course
}
