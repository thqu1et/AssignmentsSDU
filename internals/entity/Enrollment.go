package entity

import "time"

type Enrollment struct {
	Id         uint `gorm:"primaryKey"`
	StudentID  uint
	CourseID   uint
	EnrolledAt time.Time
}
