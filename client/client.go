package client

import (
    "fmt"
    "errors"
    "io/ioutil"
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
    fmt.Println("Endpoint Hit: returnCourse")

    course, err := findCourse(servemux.Value(r, "id"))
    if err != nil {
        courseNotFound(w, r, err)
        return
    }

    fmt.Println(fmt.Sprintf("Course found: %v\n", course.Id))

    json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: UpdateCourse")

    course, err := findCourse(servemux.Value(r, "id"))
    if err != nil {
        courseNotFound(w, r, err)
        return
    }

    newCourse := generateCourseFromBody(course, r)

    for index, course := range Courses {
        if course.Id == newCourse.Id {
            Courses[index] = newCourse
        }
    }

    json.NewEncoder(w).Encode(newCourse)

    fmt.Println("Course Updated:", newCourse.Id)
}

func courseNotFound(w http.ResponseWriter, r *http.Request, err error) {
    fmt.Println("Course not found:", err)
    http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func findCourse(id string) (model.Course, error) {
    for _, course := range Courses {
        if course.Id == id {
            return course, nil
        }
    }

    return model.Course{}, errors.New("No course found")
}

func generateCourseFromBody(course model.Course, r *http.Request) (model.Course) {
    fmt.Println("Update Course ...", course.Id)
    var tmpCourse model.Course

    body, err := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &tmpCourse)

    fmt.Println("Update Course ...", tmpCourse, err)

    return mergeCourses(course, tmpCourse)
}

func mergeCourses(persistedCourse model.Course, newCourse model.Course) model.Course {
    if newCourse.Name != "" {
        persistedCourse.Name = newCourse.Name
    }

    return persistedCourse
}
