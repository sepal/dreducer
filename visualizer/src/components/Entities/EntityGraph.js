import React from 'react';
import {graphql} from 'react-apollo';
import gql from 'graphql-tag';

const ENTITY_TYPE_QUERY = gql`
  query root($entity: String!, $type: String!) {
    entityType (entity: $entity, type: $type) {
      type
      fields {
        name
        fields {
          name
        }
      }
    }
  }
`;

function EntityGraph(props) {
  console.log(props);
  const w = window.innerWidth;
  const h = window.innerHeight - 75;
  const aspectRation = w/h;

  return (
    <div className="scaling-svg-container">
      <svg x="0"
           y="0"
           width={`${w}px`}
           height={`${h}px`}
           viewBox={`0 0 ${w} ${h}`}
           preserveAspectRatio="none">
      </svg>
    </div>
  )
}

export default graphql(ENTITY_TYPE_QUERY, {
  options: (props) => {
    return {
      variables: {
        entity: props.match.params.entity,
        type: props.match.params.type
      }
    }
  },
  props: ({data: {loading, entityType}}) => ({
    loading,
    entityType,
  })
})(EntityGraph);