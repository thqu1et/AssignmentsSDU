package entity

type Department struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Students []Student
}
