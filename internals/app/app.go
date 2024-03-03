package app

import (
	"github.com/gorilla/mux"
	"github.com/thqu1et/AssignmentsSDU.git/db"
	"github.com/thqu1et/AssignmentsSDU.git/internals/controller"
	"github.com/thqu1et/AssignmentsSDU.git/internals/entity"
	"github.com/thqu1et/AssignmentsSDU.git/internals/repository"
	"log"
	"net/http"
)

func Run() {
	db, err := db.ConnectMySql()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	db.AutoMigrate(&entity.Department{}, &entity.Enrollment{}, &entity.Instructor{}, &entity.Student{}, &entity.Course{})

	router := mux.NewRouter()
	controller.NewRouter(router, repository.NewRepo(db))

	http.ListenAndServe(":8080", router)
}
