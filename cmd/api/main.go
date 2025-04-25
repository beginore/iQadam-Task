package main

import (
	"database/sql"
	"log"
	"ums/internal/config"
	"ums/internal/repository"
	"ums/internal/service"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Application struct {
	config            *config.Config
	userRepo          repository.UserRepository
	courseRepo        repository.CourseRepository
	enrollmentRepo    repository.EnrollmentRepository
	authService       service.AuthService
	userService       service.UserService
	courseService     service.CourseService
	enrollmentService service.EnrollmentService
}

func main() {
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	userService := service.NewUserService(userRepo)
	courseService := service.NewCourseService(courseRepo)
	enrollmentService := service.NewEnrollmentService(
		enrollmentRepo,
		courseRepo,
		userRepo,
	)

	app := &Application{
		config:            cfg,
		userRepo:          userRepo,
		courseRepo:        courseRepo,
		enrollmentRepo:    enrollmentRepo,
		authService:       authService,
		userService:       userService,
		courseService:     courseService,
		enrollmentService: enrollmentService,
	}

	if err := app.serve(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
