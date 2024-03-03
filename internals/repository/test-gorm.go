package repository

import (
	"fmt"
	"github.com/thqu1et/AssignmentsSDU.git/internals/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type dbRun struct {
	db *gorm.DB
}
type IdString struct {
	id string
}

func TestGetStudentCountByDepartmentID(t *testing.T) {
	var dsn = "root:password@tcp(localhost:3306)/assigment2?charset=utf8mb4&parseTime=True&loc=Local"

	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}
	// Migrate the schema
	db.AutoMigrate(&entity.Student{})

	// Test the function
	tests := []struct {
		name      string
		dbRun     dbRun
		IdString  IdString
		wantCount int64
		wantErr   bool
	}{
		{
			name: "Test for not existing student",
			dbRun: dbRun{
				db: db,
			},
			IdString: IdString{
				id: "4",
			},
			wantCount: 0,
			wantErr:   false,
		},
		{
			name: "Test for success",
			dbRun: dbRun{
				db: db,
			},
			IdString: IdString{
				id: "1",
			},
			wantCount: 2,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.dbRun.db,
			}
			gotCount, err := r.CountStudentsInDepartment(tt.IdString.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountStudentsInDepartment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("CountStudentsInDepartment() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
