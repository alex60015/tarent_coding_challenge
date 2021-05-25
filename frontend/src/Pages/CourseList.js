import React, { Component } from "react";
import "./Css/shared.css"

export default class CourseList extends Component {
    renderChildren() {
        let children = [];
        this.props.courses.forEach((course, index) => (
            children.push(
                <a onClick={(e) => this.props.selectCourse(course.id)} key={index}>{CourseEntry(course)}</a>
            )
        ))

        return children
    }

    render() {
        return (
            <div className="courseList">
                <h1>Eine wundervolle Liste an Kursen</h1>
                {this.renderChildren()}
            </div>
        )
    }
}

const CourseEntry = (course) => (
    <div className="courseEntry" key={course.id}>
        <div className="picture">
            <span>Bild<br />vom<br />Kurs</span>
        </div>
        <div className="data">
            <p className="name">Name: {course.name}</p>
            <p className="prof">Dozent: {course.prof_name}</p>
            <p className="price">Preis: {course.price}</p>
            <p className="online">
                Online: {course.online ? <span className="check" /> : <span className="cross">X</span>}</p>
            <p className="date">
                Zeitpunkt: <span key={course.date}>{new Date(course.date).toLocaleString("de-de")}</span>
            </p>
        </div>
    </div>
)
