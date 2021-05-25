package main

import (
    "testing"
    "strings"
    "net/http"
    "net/http/httptest"
    "go.reizu.org/servemux"
)

import "TCC/client"

func TestReturnAllCourses(t *testing.T) {
    expected := `[{"id":"1","name":"How to create a Restful API","prof_name":"Goddart Goethe","description":"This time we will discuss how we can plan and create a restful API","price":"10.00€","online":true,"dates":[1621949400000]},{"id":"2","name":"How to create Linux","prof_name":"Linus Torvalds","description":"This time we will discuss how to create Linux","price":"00.00€","online":true,"dates":[1621941000000,1621941500000,1621942000000]},{"id":"3","name":"How to Eat","prof_name":"Anon","description":"We. Will. Eat","price":"100.00€","online":false,"dates":[1621942465000]}]`
    client.Courses = Courses

    mux := initMux()
    request, err := http.NewRequest("GET", "/courses", nil)
    if err != nil {
        t.Fatal(err)
    }

    requestRecorder := httptest.NewRecorder()
    mux.ServeHTTP(requestRecorder, request)

    if status := requestRecorder.Code; status != http.StatusOK {
        t.Errorf("Returned wrong status code: \n wanted %v \n got %v", http.StatusOK, status)
    }

    got := strings.TrimRight(requestRecorder.Body.String(), "\n")

    if expected != got {
        t.Errorf("Returned wrong response: \n wanted %v \n got %v", expected, got)
    }
}

func TestReturnCourse(t *testing.T) {
    expected := `{"id":"1","name":"How to create a Restful API","prof_name":"Goddart Goethe","description":"This time we will discuss how we can plan and create a restful API","price":"10.00€","online":true,"dates":[1621949400000]}`
    client.Courses = Courses

    mux := initMux()
    request, err := http.NewRequest("GET", "/courses/1", nil)

    if err != nil {
        t.Fatal(err)
    }

    requestRecorder := httptest.NewRecorder()
    mux.ServeHTTP(requestRecorder, request)

    if status := requestRecorder.Code; status != http.StatusOK {
        t.Errorf("Returned wrong status code: \n wanted %v \n got %v", http.StatusOK, status)
        return
    }

    got := strings.TrimRight(requestRecorder.Body.String(), "\n")

    if expected != got {
        t.Errorf("Returned wrong response: \n wanted %v \n got %v", expected, got)
    }
}

func TestUpdateCourse(t *testing.T) {
    sending := `{"id":"2","name":"Captured!","prof_name":"Tester","description":"Captured from the testing unit","price":"UNLIMITED","online":true,"dates":[1621949400000]}`
    expected := `{"id":"1","name":"Captured!","prof_name":"Tester","description":"Captured from the testing unit","price":"UNLIMITED","online":true,"dates":[1621949400000]}`
    client.Courses = Courses

    mux := initMux()
    request, err := http.NewRequest("PUT", "/courses/1", strings.NewReader(sending))

    if err != nil {
        t.Fatal(err)
    }

    requestRecorder := httptest.NewRecorder()
    mux.ServeHTTP(requestRecorder, request)

    if status := requestRecorder.Code; status != http.StatusOK {
        t.Errorf("Returned wrong status code: \n wanted %v \n got %v", http.StatusOK, status)
        return
    }

    got := strings.TrimRight(requestRecorder.Body.String(), "\n")

    if expected != got {
        t.Errorf("Returned wrong response: \n wanted %v \n got %v", expected, got)
    }
}

func initMux() (*servemux.ServeMux) {
    mux := servemux.New()
    mux.HandleFunc("/courses", client.ReturnAllCourses)
    mux.Handle("/courses/:id", servemux.MethodFuncs{
        http.MethodGet: client.ReturnCourse,
        http.MethodPut: client.UpdateCourse,
    })
    return mux
}
