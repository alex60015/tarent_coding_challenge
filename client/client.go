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

func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome User")
    fmt.Println("Endpoint Hit: homePage")

    return
}

func ReturnCourses(w http.ResponseWriter, r *http.Request) {
    var validId = regexp.MustCompile(`^\/courses[\/]?(?P<id>[\d]{0,10})[\/]?$`)
    result := validId.FindStringSubmatch(r.URL.Path)
    fmt.Println(fmt.Sprintf("%#v\n", result))

    if (len(result) == 2 && result[1] != "") {
        id, err := strconv.Atoi(result[1])
        if err != nil {
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        }
        returnCourse(w, r, id)
    } else if (len(result) != 0) {
        returnAllCourses(w, r)
    } else {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    }

    return
}

func returnCourse(w http.ResponseWriter, r *http.Request, id int) {
    fmt.Println("Endpoint Hit: returnCourse")

    course, err := findCourse(id)
    if (err != nil) {
        http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
        return
    }

    fmt.Println(fmt.Sprintf("Course found: %v\n", course.Id))

    json.NewEncoder(w).Encode(course)
    return
}

func returnAllCourses(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllCourses")
    json.NewEncoder(w).Encode(Courses)

    return
}

func findCourse(id int) (model.Course, error) {
    for _, course := range Courses {
        if course.Id == id {
            return course, nil
        }
    }

    return model.Course{}, errors.New("No course found")
}
