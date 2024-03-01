package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

func CreateStudent(db *gorm.DB, student *Student) {
	db.Create(student)
}

func GetCourseByID(db *gorm.DB, id uint) (Course, error) {
	var course Course
	if err := db.First(&course, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Course with ID %d not found", id)
			return course, err
		}
		return course, err
	}
	return course, nil
}

func UpdateInstructorName(db *gorm.DB, id uint, newName string) error {
	var instructor Instructor
	if err := db.First(&instructor, id).Error; err != nil {
		return err
	}
	instructor.Name = newName
	return db.Save(&instructor).Error
}

func DeleteDepartment(db *gorm.DB, id uint) error {
	//var count int64
	//db.Model(&Student{}).Where("department_id = ?", id).Count(&count)
	//if count > 0 {
	//	return fmt.Errorf("cannot delete department with associated students")
	//}

	if err := db.Delete(&Department{}, id).Error; err != nil {
		return err
	}
	return nil
}

func AddNewColumnToStudent(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}

func GetStudentsByDepartment(db *gorm.DB, departmentID uint) ([]Student, error) {
	var students []Student
	if err := db.Where("department_id = ?", departmentID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetCoursesByInstructor(db *gorm.DB, instructorID uint) ([]Course, error) {
	var courses []Course
	if err := db.Where("instructor_id = ?", instructorID).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func GetEnrollmentsForStudent(db *gorm.DB, studentID uint) ([]Enrollment, error) {
	var enrollments []Enrollment
	if err := db.Where("student_id = ?", studentID).Find(&enrollments).Error; err != nil {
		return nil, err
	}
	return enrollments, nil
}

func EnrollStudentInCourse(db *gorm.DB, studentID, courseID uint) error {
	var student Student
	if err := db.First(&student, studentID).Error; err != nil {
		return fmt.Errorf("cannot enroll student: %w", err)
	}

	var course Course
	if err := db.First(&course, courseID).Error; err != nil {
		return fmt.Errorf("cannot enroll student: %w", err)
	}

	enrollment := Enrollment{
		StudentID:  studentID,
		CourseID:   courseID,
		EnrolledAt: time.Now(),
	}
	if err := db.Create(&enrollment).Error; err != nil {
		return fmt.Errorf("cannot enroll student: %w", err)
	}

	return nil
}

func main() {
	db := ConnectDB()
	//defer db.Close()

	newStudent := &Student{Name: "Askar", DepartmentID: 1}
	CreateStudent(db, newStudent)

	department1 := &Department{Name: "IT"}
	department2 := &Department{Name: "Marketing"}
	InsertDepartment(db, department1)
	InsertDepartment(db, department2)

	instructor1 := &Instructor{Name: "Dr. Smith"}
	instructor2 := &Instructor{Name: "Prof. Johnson"}
	InsertInstructor(db, instructor1)
	InsertInstructor(db, instructor2)

	course1 := &Course{Name: "Web Development", InstructorID: instructor1.Id}
	course2 := &Course{Name: "Data Science", InstructorID: instructor2.Id}
	InsertCourse(db, course1)
	InsertCourse(db, course2)

	student1 := &Student{Name: "John", DepartmentID: department1.Id}
	student2 := &Student{Name: "Alice", DepartmentID: department2.Id}
	InsertStudent(db, student1)
	InsertStudent(db, student2)

	enrollment1 := &Enrollment{StudentID: student1.Id, CourseID: course1.Id, EnrolledAt: time.Now()}
	enrollment2 := &Enrollment{StudentID: student2.Id, CourseID: course2.Id, EnrolledAt: time.Now()}
	InsertEnrollment(db, enrollment1)
	InsertEnrollment(db, enrollment2)

	course, err := GetCourseByID(db, 1)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error retrieving course: %v", err)
	} else if err == gorm.ErrRecordNotFound {
		log.Printf("Course with ID 1 not found")
	} else {
		fmt.Println(course)
	}

	err = UpdateInstructorName(db, 1, "Abylkhaiyrov")
	if err != nil {
		log.Fatal(err)
	}

	err = DeleteDepartment(db, 22)
	if err != nil {
		log.Printf("Error deleting department: %v", err)
	}

	students, err := GetStudentsByDepartment(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(students)

	courses, err := GetCoursesByInstructor(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(courses)

	enrollments, err := GetEnrollmentsForStudent(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(enrollments)

	// Example transaction
	err = EnrollStudentInCourse(db, 1, 1)
	if err != nil {
		log.Printf("Error inserting enrollment: %v", err)
	}
}

// Insert a new student into the database
func InsertStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

// Insert a new course into the database
func InsertCourse(db *gorm.DB, course *Course) error {
	return db.Create(course).Error
}

// Insert a new department into the database
func InsertDepartment(db *gorm.DB, department *Department) error {
	return db.Create(department).Error
}

// Insert a new enrollment into the database
func InsertEnrollment(db *gorm.DB, enrollment *Enrollment) error {
	return db.Create(enrollment).Error
}

// Insert a new instructor into the database
func InsertInstructor(db *gorm.DB, instructor *Instructor) error {
	return db.Create(instructor).Error
}
