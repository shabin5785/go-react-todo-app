import React from "react";

import axios from "axios"

import "./addbar.css"

class AddBar extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            todo: "",
            message: ""
        }
    }

    setTodo = (event) => {
        this.setState({
            todo: event.target.value
        })
    }


    handleSubmit = async (event) => {
        event.preventDefault()
        console.log(this.state);
        let axop = {
            method: 'post',
            url: 'http://localhost:8080/createItem',
            data: {
                item: this.state.todo
            }
        }
        try {
            let res = await axios(axop)
            console.log(res);
            this.setState({...this.state, message: "A todo was created"})
        }
        catch (e) {
            console.log(e)
        }
    }
    render() {
        return (
            <div>
                <form onSubmit={this.handleSubmit}>
                    <div className="AddBar">
                        <input id="todo" name="todo"
                            value={this.state.todo}
                            type="text" onChange={this.setTodo} />
                        <button type="submit">Create ToDo</button>
                    </div>
                </form>
                <div>{this.state.message}</div>
            </div>

        )
    }
}

export default AddBar