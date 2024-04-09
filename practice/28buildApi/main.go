package main

// model for course-file
type Course struct {
	CourseId string `json:"courseid"`
	CoursePrice int `json:"price"`
	CourseName string `json:"coursename"`
	Author *Author `json:"author"`
	// Author Author // will work but it will have copy 
}

// custom type
type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course

// middlewear, helper-file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && 	c.CourseName == ""
}

func main()  {
	
}