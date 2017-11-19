import React, {Component} from 'react';
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom';
import './App.css';
import EntitySelector from "../Entities/EntitySelector";
import EntityGraph from '../Entities/EntityGraph';


class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <header className="header">
            <h1>DReducer Inspector</h1>
          </header>

          <main>
            <Route exact path="/" component={EntitySelector} />
            <Route path="/entity/:entity/:type" component={EntityGraph} />
          </main>
        </div>
      </Router>
    );
  }
}

export default App;
