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
    enableCors(&w)
    json.NewEncoder(w).Encode(Courses)
}

func ReturnCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnCourse")
    enableCors(&w)

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
    enableCors(&w)

    course, err := findCourse(servemux.Value(r, "id"))
    if err != nil {
        courseNotFound(w, r, err)
        return
    }

    newCourse, err := generateCourseFromBody(course, r)
    if err != nil {
        fmt.Println("Error while generating course:", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    }

    for index, course := range Courses {
        if course.Id == newCourse.Id {
            Courses[index] = newCourse
        }
    }

    json.NewEncoder(w).Encode(newCourse)

    fmt.Println("Course Updated! ID:", newCourse.Id)
}

func HandleOptionsCall(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Options Hit: HandleOptionsCall")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, OPTIONS")
    w.WriteHeader(http.StatusOK)
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
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

func generateCourseFromBody(course model.Course, r *http.Request) (model.Course, error) {
    fmt.Println("Update Course ... ID:", course.Id)
    var tmpCourse model.Course

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return model.Course{}, err
    }

    json.Unmarshal(body, &tmpCourse)

    return mergeCourses(course, tmpCourse), nil
}

func mergeCourses(persistedCourse model.Course, newCourse model.Course) model.Course {
    if newCourse.Name != "" {
        persistedCourse.Name = newCourse.Name
    }
    if newCourse.ProfName != "" {
        persistedCourse.ProfName = newCourse.ProfName
    }
    if newCourse.Description != "" {
        persistedCourse.Description = newCourse.Description
    }
    if newCourse.Price != "" {
        persistedCourse.Price = newCourse.Price
    }
    persistedCourse.Online = newCourse.Online

    return persistedCourse
}
