package entity

type Student struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	DepartmentID uint
	Department   Department
	Enrollments  []Enrollment
}
