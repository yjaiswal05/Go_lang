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

//Model for courses -> goes inside a seperate file

type Course struct {
	CourseId string 	`json:"courseid"`
	CourseName string 	`json:"coursename"`
	CoursePrice int 	`json:"price"`
	Author *Author 		`json:"author"`
}	
type Author struct {
	Fullname string 	`json:"fullname"`
	Website string	 	`json:"website"`
}

//fake database
var courses []Course

//middleware or helper -> goes inside a seperate file
func (c *Course) isEmpty()bool{		//in this we are passing all the values from Course in c and checking is it empty or not
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	courses = append(courses,Course{
		CourseId: "2", CourseName: "Reactjs",CoursePrice: 299,Author: &Author{Fullname: "Yash",Website: "www.google.com"}})
	courses = append(courses,Course{
		CourseId: "1", CourseName: " Nodejs",CoursePrice: 799,Author: &Author{Fullname: "Yash",Website: "www.google.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000",r))
}

//controller -> goes inside a seperate file
func serveHome(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte ("<h1>Welcome to my page </h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)									//grab id from request
	for _,course := range courses{							//loop through the courses and find the matching id and return the response
		if course.CourseId == params["id"] {	
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No code found with given id")
	return	
}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	
	if r.Body == nil {										//what if: body is empty
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {									//what if: body is {}
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	//generate unique id and then convert it in string & then append course into courses
	
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Int())
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return
}	
//func updateOneCourse(w http.ResponseWriter,r *http.Request)
//func deleteOneCourse(w http.ResponseWriter,r *http.Request)

