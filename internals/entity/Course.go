package entity

type Course struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	InstructorID uint
	Instructor   Instructor `gorm:"foreignKey:InstructorID"`
	Enrollments  []Enrollment
}
