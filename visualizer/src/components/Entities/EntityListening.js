import React from 'react';
import {graphql} from 'react-apollo';
import gql from 'graphql-tag';

const ENTITIES_QUERY = gql`query root { 
  entities {
    id
    name 
  }
}`;

function EntityListening(props) {
  if (props.loading) {
    return (
      <li>Waiting for data...</li>
    )
  }

  return props.entities.map(entity => <li key={entity.id}>
    <a href="/">{entity.name}</a>
  </li>);
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