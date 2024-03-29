package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseID string `json:"courseID"`
	CourseName string `json:"courseName"`
	CoursePrice float64 `json:"coursePrice"`
	// CoursePrice float64 `json:"-"`
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
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to build RESTful API with Golang!")

	// init router
	r := mux.NewRouter()
	
	// seeding
	courses = append(courses, Course{CourseID: "1", CourseName: "Golang", CoursePrice: 100.00, Student: &Student{FullName: "Abhi Singh", Email: "Abhi@gmail.com"}})
	courses = append(courses, Course{CourseID: "2", CourseName: "Js", CoursePrice: 199.00, Student: &Student{FullName: "Abhi Singh", Email: "Abhi@gmail.com"}})
	courses = append(courses, Course{CourseID: "3", CourseName: "DB", CoursePrice: 199.00, Student: &Student{FullName: "Abhi Singh", Email: "Abhi@gmail.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	r.HandleFunc("/course", createCourse).Methods("POST")

	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")

	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
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

// create a course
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	// what if the request body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// generate a unique id, string
	// append the course to the courses

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// update a course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	// get the course id from the url
	params := mux.Vars(r);

	// loop through the courses to find, add and remove the course with the id

	for index, course := range courses {
		if course.CourseID == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

// delete a course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	// get the course id from the url
	params := mux.Vars(r);

	// loop through the courses to find and remove the course with the id
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the id")
	return
}
