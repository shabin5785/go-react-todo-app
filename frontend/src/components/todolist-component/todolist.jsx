import React from "react";
import { EditText } from "react-edit-text"
import "./todolist.css"

import axios from "axios"


class TodoList extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            items: {}
        }
    }

    updateTodo = (event, id) => {
        let items = this.state.items;
        let obj = items[id];
        obj.item = event;
        this.setState({
            items: items
        })
    }

    saveUpdatedTodo = (e, item) => {
        let newTodo = e.value;
        let ob = {};
        ob.id = item.id;
        ob.item = newTodo;
        ob.olditem = item.item;
        ob.done = item.done
        let axop = {
            method: 'post',
            url: 'http://localhost:8080/updateItem',
            data: ob
        }

        axios(axop).then(res => {
            this.loadItems()
        }).catch(e => {
            console.log(e)
        })



        //{"id":18, "item":"this is a new one", "done":true, "olditem":"working value"}
    }

    componentDidMount() {

        this.loadItems()

    }

    loadItems() {
        let axop = {
            method: 'get',
            url: 'http://localhost:8080/items',
        }
        axios(axop).then(res => {
            // console.log(res.data)
            let arrItems = res.data.items;
            let items = {}
            for (let a of arrItems) {
                items[a.id] = a
            }
            this.setState({
                items: items
            })
            console.log("items were loaded")
        }).catch(e => {
            console.log(e)
        })
    }

    removeItem = (id) => {
        let axop = {
            method: 'get',
            url: `http://localhost:8080/deleteItem/${id}`
        }
        axios(axop).then(res => {
            this.loadItems()
        }).catch(e => {
            console.log(e)
        })
    }

    toggleTodo = (item) => {
        let ob = {};
        ob.id = item.id;
        ob.item = item.item;;
        ob.olditem = item.item;
        ob.done = !item.done
        let axop = {
            method: 'post',
            url: 'http://localhost:8080/updateItem',
            data: ob
        }

        axios(axop).then(res => {
            this.loadItems()
        }).catch(e => {
            console.log(e)
        })
    }

    createItem(id) {
        let item = this.state.items[id];
        return (
            <div className="ListItem" key={item.id} id={item.id}>
                <div className="RemoveItem" onClick={e => {
                    this.removeItem(id)
                }}>
                    X
                </div>
                <div className="Title">
                    <EditText value={item.item.trim()} onSave={(e) => this.saveUpdatedTodo(e, item)} onChange={(e) => {
                        this.updateTodo(e, id)
                    }} />

                </div>
                <div className="Status" onClick={e => {
                    this.toggleTodo(item)
                }}>
                    {item.done ? "Done" : "Not Done"}
                </div>
            </div>
        )
    }
    render() {
        let items = this.state.items
        return (
            <div className="TodoList">
                <div className="List">
                    {Object.keys(items).map(id => this.createItem(id))}
                </div>
            </div>)
    }
}

export default TodoList