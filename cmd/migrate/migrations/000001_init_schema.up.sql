CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       full_name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       role VARCHAR(20) NOT NULL CHECK (role IN ('ADMIN', 'TEACHER', 'STUDENT')),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE courses (
                         id SERIAL PRIMARY KEY,
                         title VARCHAR(100) NOT NULL,
                         description TEXT,
                         teacher_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE enrollment (
                            id SERIAL PRIMARY KEY,
                            student_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                            course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                            CONSTRAINT unique_enrollment UNIQUE (student_id, course_id)
);

CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_courses_title ON courses (title);
CREATE INDEX idx_enrollment_student ON enrollment (student_id);
CREATE INDEX idx_enrollment_course ON enrollment (course_id);