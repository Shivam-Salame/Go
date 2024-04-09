package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	return c.CourseId == "" && c.CourseName == ""
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

func main() {

}
