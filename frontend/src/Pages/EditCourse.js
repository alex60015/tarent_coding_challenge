import React, { Component } from "react";
import "./Css/shared.css"

export default class EditCourse extends Component {
    constructor(props) {
        super(props);

        this.state = {
            course: null,
        }
    }

    componentDidMount() {
        fetch("http://localhost:8000/courses/" + this.props.courseId)
            .then(response => response.json())
            .then(json => this.setState((state, props) => ({
                course: json
            })));
    }

    onChangeHandler = (event) => {
        const { name, value, type, checked } = event.target
        let course = this.state.course
        console.log("Update: ", name, value, type, checked)

        if (type === "checkbox") {
            course[name] = checked
        } else {
            course[name] = value
        }

        this.setState((state, props) => ({
            course: course
        }))
    }

    saveCourse = (event) => {
        const options = {
            method: "PUT",
            body: JSON.stringify(this.state.course)
        }

        fetch("http://localhost:8000/courses/" + this.props.courseId, options)
            .then(response => response.json())
            .then(json => {
                console.log(json)
                this.props.endEditing()
            })
    }

    render() {
        const course = this.state.course

        if (course === null) {
            return (
                <div className="editCourse">
                    <h1>Kurs bearbeiten</h1>
                    <p>L&auml;d...</p>
                </div>
            )
        }

        return (
            <div className="editCourse">
                <h1>Kurs bearbeiten</h1>
                <div className="courseEntry">
                    <div className="picture">
                        <span>Bild<br />vom<br />Kurs</span>
                    </div>
                    <div className="data">
                        <EditableProperty name="name" label="Name" value={course.name} onChangeHandler={this.onChangeHandler} />
                        <EditableProperty name="prof_name" label="Dozent" value={course.prof_name} onChangeHandler={this.onChangeHandler} />
                        <EditableProperty name="description" label="Beschreibung" value={course.description} onChangeHandler={this.onChangeHandler} />
                        <EditableProperty name="price" label="Preis" value={course.price} onChangeHandler={this.onChangeHandler} />
                        <p className="online">
                            Online:
                            <input
                                type="checkbox"
                                name="online"
                                checked={course.online ? "selected" : ""}
                                alt="Online"
                                onChange={this.onChangeHandler}
                            />
                        </p>

                        <p className="date">Zeitpunkt:
                            <input
                                name="date"
                                type="number"
                                value={course.date}
                                alt="Zeitpunkt"
                                onChange={this.onChangeHandler}
                            />
                        </p>
                    </div>
                    <button onClick={this.saveCourse}>Speichern</button>
                </div>
            </div>
        )
    }
}

class EditableProperty extends Component {
    render() {
        const { name, label, value, onChangeHandler } = this.props

        return (
            <p className={name}>{label}:
                <input
                    name={name}
                    value={value}
                    alt={label}
                    onChange={onChangeHandler}
                />
            </p>
        )
    }
}
