import './App.css';

import Header from "./components/header-component"
import AddBar from './components/addBar-component/addBar';
import TodoList from './components/todolist-component/todolist';
import React from 'react';

class App extends React.Component {

  render() {
    return (
      <div className="App">
       <Header/>
       <AddBar/>
       <TodoList/>
      </div>
    );
  }
  
}

export default App;
