import React from 'react';
import Entities from '../Entities/EntityListening';
import List from '../List';


function EntitySelector(props) {
  return (
    <div>
      <h2>Select an entity:</h2>
      <List>
        <Entities />
      </List>
    </div>
  )
}

export default EntitySelector;