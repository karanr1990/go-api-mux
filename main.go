package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/karanr1990/go-api-mux/router"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	//c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api build</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get single course")
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == param["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No Such Record")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get single course")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Plz send some data")
	}
	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Data inside json")
		return
	}
	//generate unique id string
	//append course intocourses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from request
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete date one course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from request
	params := mux.Vars(r)

	//remove index data
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			return
		}
	}

}

func main() {
	//router := mux.NewRouter()
	//courses = append(courses, Course{CourseId: "100",
	//	CourseName:  "React",
	//	CoursePrice: 200,
	//	Author: &Author{
	//		Fullname: "karan",
	//		Website:  "javatpoint.com",
	//	}})
	//courses = append(courses, Course{CourseId: "200",
	//	CourseName:  "JAVA",
	//	CoursePrice: 200,
	//	Author: &Author{
	//		Fullname: "ARJUN",
	//		Website:  "javatpoint.com",
	//	}})
	//router.HandleFunc("/", serveHome).Methods("GET")
	//router.HandleFunc("/courses", getAllCourses).Methods("GET")
	//router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	//router.HandleFunc("/course", createOneCourse).Methods("POST")
	//router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	//router.HandleFunc("/courses", deleteOneCourse).Methods("DELETE")
	//
	//http.ListenAndServe(":4000", router)

	fmt.Println("MongoDb API")
	r := router.Router()

	fmt.Println("server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
