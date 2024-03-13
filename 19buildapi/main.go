package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseID string `json:"courseID"`
	CourseName string `json:"courseName"`
	CoursePrice float64 `json:"coursePrice"`
	Student *Student `json:"student"`
}

// Model for students - file
type Student struct {
	FullName string `json:"fullName"`
	Email string `json:"email"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to build RESTful API with Golang!")
}

// Controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// get the course id from the url
	params := mux.Vars(r);

	// loop through the courses to find the course with the id
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the id")
	return
}