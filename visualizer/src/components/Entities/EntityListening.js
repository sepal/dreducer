import React from 'react';
import {graphql} from 'react-apollo';
import gql from 'graphql-tag';
import List from "../List";

const ENTITIES_QUERY = gql`query root { 
  entities {
    id
    name 
    types {
      id
      name
    }
  }
}`;


function EntityListening(props) {
  if (props.loading) {
    return (
      <li>Waiting for data...</li>
    )
  }

  return props.entities.map((entity) => {
    const types = entity.types.map(et => <li key={et.id}><a href="">{et.name}</a></li>);

    return (<li key={entity.id}>
      <a href="/">{entity.name}</a>
      <List>
        {types}
      </List>
    </li>)
  });
}

export default graphql(ENTITIES_QUERY, {
  options: {
    fetchPolicy: 'cache-and-network',
  },
  props: ({data: {loading, entities}}) => ({
    loading,
    entities,
  }),
})(EntityListening)