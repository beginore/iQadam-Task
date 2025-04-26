package main

import (
	"github.com/gin-gonic/gin"
	"ums/internal/controller"
	"ums/internal/middleware"
)

func (app *Application) routes() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.ErrorHandler())

	// Public routes
	authCtrl := controller.NewAuthController(app.authService)
	r.POST("/auth/register", authCtrl.Register)
	r.POST("/auth/login", authCtrl.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(app.config.JWTSecret))
	{
		// Users
		userCtrl := controller.NewUserController(app.userService)
		auth.GET("/users/:id", userCtrl.GetUser)
		auth.PUT("/users/:id", userCtrl.UpdateUser)
		auth.DELETE("/users/:id", userCtrl.DeleteUser)
		auth.GET("/users", userCtrl.ListUsers)

		// Courses
		courseCtrl := controller.NewCourseController(app.courseService)
		auth.GET("/courses", courseCtrl.GetCourses)
		auth.POST("/courses", courseCtrl.CreateCourse).Use(middleware.RoleMiddleware("ADMIN", "TEACHER"))
		auth.PUT("/courses/:id", courseCtrl.UpdateCourse).Use(middleware.RoleMiddleware("ADMIN", "TEACHER"))
		auth.DELETE("/courses/:id", courseCtrl.DeleteCourse).Use(middleware.RoleMiddleware("ADMIN"))

		enrollCtrl := controller.NewEnrollmentController(app.enrollmentService)
		auth.POST("/enroll", enrollCtrl.EnrollStudent).Use(middleware.RoleMiddleware("STUDENT", "TEACHER"))
		auth.DELETE("/enroll", enrollCtrl.UnenrollStudent).Use(middleware.RoleMiddleware("ADMIN", "TEACHER"))
		auth.GET("/students/:student_id/enrollments", enrollCtrl.GetEnrollmentsByStudent)
		auth.GET("/courses/:course_id/enrollments", enrollCtrl.GetEnrollmentsByCourse)
		auth.GET("/enrollments", enrollCtrl.GetAllEnrollments).Use(middleware.RoleMiddleware("ADMIN"))
	}

	return r
}
