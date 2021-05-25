package main

import (
    "log"
    "net/http"
    "go.reizu.org/servemux"
)

import "TCC/model"
import "TCC/client"

var Courses = []model.Course{
    model.Course{
        Id: "1",
        Name: "How to create a Restful API",
        ProfName: "Goddart Goethe",
        Description: "This time we will discuss how we can plan and create a restful API",
        Price: "10.00€",
        Online: true,
        Date: 1621949400000,
    },
    model.Course{
        Id: "2",
        Name: "How to create Linux",
        ProfName: "Linus Torvalds",
        Description: "This time we will discuss how to create Linux",
        Price: "00.00€",
        Online: true,
        Date: 1621941000000,
    },
    model.Course{
        Id: "3",
        Name: "How to Eat",
        ProfName: "Anon",
        Description: "We. Will. Eat",
        Price: "100.00€",
        Online: false,
        Date: 1621942465000,
    },
}

func main() {
    client.Courses = Courses

    mux := servemux.New()
    mux.HandleFunc("/courses", client.ReturnAllCourses)
    mux.Handle("/courses/:id", servemux.MethodFuncs{
        http.MethodGet: client.ReturnCourse,
        http.MethodPut: client.UpdateCourse,
        http.MethodOptions: client.HandleOptionsCall,
    })

    log.Fatal(http.ListenAndServe(":8000", mux))
}
