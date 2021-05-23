package client

import (
    "fmt"
    "errors"
    "encoding/json"
    "net/http"
    "go.reizu.org/servemux"
)

import "TCC/model"

var Courses []model.Course

func ReturnAllCourses(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllCourses")
    json.NewEncoder(w).Encode(Courses)
}

func ReturnCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnCourse", )

    course, err := findCourse(servemux.Value(r, "id"))
    if err != nil {
        courseNotFound(w, r, err)
        return
    }

    fmt.Println(fmt.Sprintf("Course found: %v\n", course.Id))

    json.NewEncoder(w).Encode(course)
}

func courseNotFound(w http.ResponseWriter, r *http.Request, err error) {
    fmt.Println("Course not found:", err)
    http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func handlePathError(w http.ResponseWriter, r *http.Request, err error) {
    fmt.Println("Error in returnCourses:", err)
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func findCourse(id string) (model.Course, error) {
    for _, course := range Courses {
        if course.Id == id {
            return course, nil
        }
    }

    return model.Course{}, errors.New("No course found")
}

func updateCourseData(course model.Course, r *http.Request) (model.Course, error) {
    return model.Course{}, nil
}
