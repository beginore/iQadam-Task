# University Management System (UMS)

![Go](https://img.shields.io/badge/Go-1.21%2B-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15%2B-brightgreen)

REST API University Management System with JWT-authentication and RBAC (ADMIN/TEACHER/STUDENT).

---

## 🚀 Quick start

### Requirements
- Go 1.21+
- PostgreSQL 15+
-  `migrate` ([установка](https://github.com/golang-migrate/migrate))

### Project setup

1. **Database**:
   ```bash
   # Create .env file
   echo 'DATABASE_URL="postgres://<user>:<password>@localhost:<5432>/ums?sslmode=disable"' > .env
   # Build migration util
   go build -o migrate cmd/migrate/main.go
   # Execute migration (up|down)
   ./migrate up
   ```
2. **Start server**:
   ```bash
   go run cmd/api/main.go
   ```

### 📂 Project structure
```bash
├── cmd/
│   ├── api/          # main server
│   └── migrate/      # migration tool
├── internal/
│   ├── config/       # config
│   ├── controller/   # HTTP handlers
│   ├── repository/   # database operations
│   └── dto           # converts structs to DTO
│   └── service       # use-case layer
│   └── middleware    # HTTP middlewares
│   └── utils         # JWT, password tools   
└── migrations/       # SQL-migrations
```
## End-points
### Authentication

### 1. Register User
  Endpoint: POST `/auth/register`  
  Description: Register a new user in the system.  
  Roles Allowed: Public access  
  Request Body:
```json
{
  "username": "teacher",
  "password": "Te@cherPass123",
  "full_name": "John Doe",
  "email": "teacher@university.edu",
  "role": "TEACHER"
}
```
Response (201 Created):
```json
{
  "id": 2,
  "username": "teacher",
  "email": "teacher@university.edu",
  "role": "TEACHER",
  "created_at": "2023-10-20T15:04:05Z"
}
```
### 2. Login
   Endpoint: POST `/auth/login`
   Description: Authenticate user and return JWT token.
   Roles Allowed: Public access
   Request Body:

```json
{
  "username": "teacher",
  "password": "Te@cherPass123"
}
```
Response (200 OK):
```json
{
   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
   "user": {
      "id": 2,
      "username": "teacher",
      "role": "TEACHER"
   }
}
```
### Users Management
### 3. Get All Users
Endpoint: GET `/users`
Description: Retrieve list of all users (Admin only).
Roles Allowed: ADMIN
Response (200 OK):
```json
[
   {
      "id": 1,
      "username": "admin",
      "role": "ADMIN",
      "created_at": "2023-10-20T14:30:00Z"
   }
]
```
### 4. Get User by ID
Endpoint: GET `/users/{id}`
Description: Get specific user details.
Roles Allowed: ADMIN
Response (200 OK):
```json
{
   "id": 2,
   "username": "teacher",
   "role": "TEACHER",
   "courses": []
}
```
### Courses Management
### 5. Create Course
Endpoint: POST `/courses`
Description: Create new course.
Roles Allowed: ADMIN, TEACHER
Request Body:
```json
{
   "title": "Advanced Mathematics",
   "description": "Master calculus and linear algebra",
   "teacher_id": 2
}
```
Response (201 Created):
```json
{
   "id": 1,
   "title": "Advanced Mathematics",
   "teacher_id": 2,
   "created_at": "2023-10-20T16:00:00Z"
}
```
### 6. Get Courses
Endpoint: GET `/courses`
Description: Get list of courses with optional sorting.
Roles Allowed: All authenticated users
Query Params:

?sort=date (default|date|enrollment)
Response (200 OK):
```json
[
   {
      "id": 1,
      "title": "Advanced Mathematics",
      "students_count": 15
   }
]
```
### Enrollment Management
### 7. Enroll Student
Endpoint: POST `/enroll`
Description: Enroll student to course.
Roles Allowed: ADMIN, TEACHER
Request Body:
```json
{
   "student_id": 3, //must be userID of user with role [STUDENT]
   "course_id": 1
}
```
Response (201 Created):
```json
{
  "message": "Student enrolled successfully"
}
```