package entity

type Instructor struct {
	Id      uint `gorm:"primaryKey"`
	Name    string
	Courses []Course
}
