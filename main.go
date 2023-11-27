package go_api_mux

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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

func main() {
	router := mux.NewRouter()
	courses = append(courses, Course{CourseId: "2",
		CourseName:  "React",
		CoursePrice: 200,
		Author: &Author{
			Fullname: "karan",
			Website:  "javatpoint.com",
		}})
	router.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":4000", router)

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api build</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getSingleCourse(w http.ResponseWriter, r *http.Request) {
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

func createSingleCourse(w http.ResponseWriter, r *http.Request) {
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
}
