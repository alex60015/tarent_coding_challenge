package main

import (
    "log"
    "net/http"
)

import "TCC/model"
import "TCC/client"

var Courses = []model.Course{
    model.Course{
       Id: 1,
        Name: "How to create a Restful API",
        ProfName: "Goddart Goethe",
        Description: "This time we will discuss how we can plan and create a restful API",
        Price: "10.00€",
        Online: true,
        Dates: []int{1621949400},
    },
    model.Course{
        Id: 2,
        Name: "How to create Linux",
        ProfName: "Linus Torvalds",
        Description: "This time we will discuss how to create Linux",
        Price: "00.00€",
        Online: true,
        Dates: []int{1621941000},
    },
    model.Course{
        Id: 3,
        Name: "How to Eat",
        ProfName: "Anon",
        Description: "We. Will. Eat",
        Price: "100.00€",
        Online: true,
        Dates: []int{1621942465},
    },
}

func setupEndpoints() {
    http.HandleFunc("/courses/", client.HandleCourses)
}

func handleRequests() {
    client.Courses = Courses
    setupEndpoints()
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    handleRequests()
}
