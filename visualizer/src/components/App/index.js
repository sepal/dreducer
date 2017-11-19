import React, {Component} from 'react';
import './App.css';
import EntitySelector from "../Entities/EntitySelector";


class App extends Component {
  render() {
    return (
      <div>
        <header className="header">
          <h1>DReducer Inspector</h1>
        </header>
        <EntitySelector/>
      </div>
    );
  }
}

export default App;
