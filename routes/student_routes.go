package routes

import (
	"Xtern-Matching/handlers"

	"github.com/gorilla/mux"
)

func GetStudentRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/student/reviewer", handlers.GetReviewedStudents).Methods("GET").Name("ReviewedStudents")
	router.HandleFunc("/student/updatestatus", handlers.UpdateStudentsToStatus).Methods("POST").Name("UpdateStudentStatus")
	router.HandleFunc("/student/status", handlers.GetStudentsAtLeastWithStatus).Methods("POST").Name("StudentsWithStatus")
	router.HandleFunc("/student/light", handlers.GetStudentDecisionList).Methods("GET").Name("GetDecisionList")
	router.HandleFunc("/student/resumes", handlers.ExportResumes).Methods("GET").Name("ExportResumes")
	router.HandleFunc("/student/export", handlers.ExportStudents).Methods("GET").Name("ExportStudents")
	router.HandleFunc("/student/{studentKey}", handlers.GetStudent).Methods("GET").Name("GetStudent")
	router.HandleFunc("/student", handlers.GetStudents).Methods("GET").Name("GetStudents")
	router.HandleFunc("/student", handlers.AddStudent).Methods("POST").Name("AddStudent")
	router.HandleFunc("/student/{studentKey}/grade", handlers.SetGrade).Methods("PUT").Name("SetGrade")
	router.HandleFunc("/student/{studentKey}/status", handlers.SetStatus).Methods("PUT").Name("SetStatus")
	return router
}
