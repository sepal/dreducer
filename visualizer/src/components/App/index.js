import React, {Component} from 'react';
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom';
import './App.css';
import EntitySelector from "../Entities/EntitySelector";


class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <header className="header">
            <h1>DReducer Inspector</h1>
          </header>

          <Route exact path="/" component={EntitySelector}/>
        </div>
      </Router>
    );
  }
}

export default App;
