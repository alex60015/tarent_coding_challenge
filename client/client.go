package client

import (
    "fmt"
    "errors"
    "regexp"
    "strconv"
    "encoding/json"
    "net/http"
)

import "TCC/model"

var Courses []model.Course

func HandleCourses(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ReturnCourses(w, r)
    }
}

func ReturnCourses(w http.ResponseWriter, r *http.Request) {
    id, err := getIdFromPath(r.URL.Path)

    if err != nil {
        fmt.Println("Error in ReturnCourses:", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }
    if (id != 0) {
        returnCourse(w, r, id)
        return
    }

    returnAllCourses(w, r)

    return
}

func returnCourse(w http.ResponseWriter, r *http.Request, id int) {
    fmt.Println("Endpoint Hit: returnCourse")

    course, err := findCourse(id)
    if (err != nil) {
        fmt.Println("Course not found in returnCourse:", err)
        http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    }

    fmt.Println(fmt.Sprintf("Course found: %v\n", course.Id))

    json.NewEncoder(w).Encode(course)
}

func returnAllCourses(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllCourses")
    json.NewEncoder(w).Encode(Courses)
}

func findCourse(id int) (model.Course, error) {
    for _, course := range Courses {
        if course.Id == id {
            return course, nil
        }
    }

    return model.Course{}, errors.New("No course found")
}

func getIdFromPath(path string) (int, error) {
    validId := regexp.MustCompile(`^\/courses[\/]?(?P<id>[\d]{0,10})[\/]?$`)
    result := validId.FindStringSubmatch(path)

    fmt.Println(fmt.Sprintf("%#v\n", result))

    if len(result) == 0 {
        return 0, errors.New("Invalid Request")
    } else if (len(result) == 2 && result[1] != "") {
        id, err := strconv.Atoi(result[1])

        if err != nil {
            return 0, err
        }

        return id, nil
    }

    return 0, nil
}
