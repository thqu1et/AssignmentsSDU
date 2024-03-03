package controller

import (
	"github.com/gorilla/mux"
	"github.com/thqu1et/AssignmentsSDU.git/internals/repository"
	"net/http"
)

type router struct {
	mux  *mux.Router
	repo *repository.Repo
}

func NewRouter(mux *mux.Router, repo *repository.Repo) {
	r := &router{mux: mux, repo: repo}

	studentsRouter := mux.PathPrefix("/student").Subrouter()
	courseRouter := mux.PathPrefix("/course").Subrouter()
	departmentRouter := mux.PathPrefix("/department").Subrouter()
	instructorRouter := mux.PathPrefix("/instructor").Subrouter()
	enrollmentRouter := mux.PathPrefix("/enrollment").Subrouter()

	NewStudentRouter(studentsRouter, repo)
	NewCourseRouter(courseRouter, repo)
	NewDepartmentRouter(departmentRouter, repo)
	NewInstructorRouter(instructorRouter, repo)
	NewEnrollmentRouter(enrollmentRouter, repo)

	mux.HandleFunc("/", r.getByID)

}

func (r router) getByID(w http.ResponseWriter, req *http.Request) {

}
