package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course-file
type Course struct {
	CourseId    string  `json:"courseid"`
	CoursePrice int     `json:"price"`
	CourseName  string  `json:"coursename"`
	Author      *Author `json:"author"`
	// Author Author // will work but it will have copy
}

// custom type
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middlewear, helper-file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

// controllers-file

// server home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api by shivam salame</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)
	
	// loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("No course found with matching id")
	return 
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// if {}
	var course Course
	_ =json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}	

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range 	courses {
			if course.CourseId == params["id"]
				courses = append(courses[:index], courses[index+1]...)
				var course Course
				_ = json.NewDecoder(r.Body).Decode(&course)
				course.CourseId = params.["id"]
				courses = append(courses, course)
				json.NewEncoder(w).Encode(course)
				return
			}
}

func main() {

}
