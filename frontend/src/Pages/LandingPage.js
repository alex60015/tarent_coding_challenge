import React, { Component } from "react";
import CourseList from "./CourseList";
import "./Css/shared.css"
import EditCourse from "./EditCourse";

export default class LandingPage extends Component {
    constructor(props) {
        super(props);

        this.state = {
            courses: [],
            selectedCourse: null,
        }
    }

    componentDidMount() {
        this.updateList()
    }

    updateList = () => {
        fetch("http://localhost:8000/courses")
            .then(response => response.json())
            .then(json => this.setState((state, props) => ({
                courses: json
            })));
    }

    selectCourse = (courseId) => {
        this.setState((state, props) => ({
            selectedCourse: courseId
        }))
    }

    endEditing = () => {
        this.updateList()
        this.setState((state, props) => ({
            selectedCourse: null
        }))
    }

    render() {
        let element

        if (this.state.selectedCourse) {
            element = <EditCourse courseId={this.state.selectedCourse} endEditing={this.endEditing}/>
        } else if (this.state.courses.length > 0) {
            element = <CourseList courses={this.state.courses} selectCourse={this.selectCourse}/>
        } else {
            element = <p>L&auml;d...</p>
        }

        return (
            <div className="content">
                {element}
            </div>
        )
    }
}
