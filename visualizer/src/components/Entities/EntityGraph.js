import React from 'react';
import {graphql} from 'react-apollo';
import gql from 'graphql-tag';

const ENTITY_TYPE_QUERY = gql`
  query root($entity: String!, $type: String!) {
    entityType (entity: $entity, type: $type) {
      entity
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

function EntityGraph({loading, entityType}) {
  if (loading) {
    return (
      <div>
        Loading data...
      </div>
    )
  }

  const w = window.innerWidth;
  const h = window.innerHeight - 75;
  const aspectRation = w / h;

  const midX = w / 2;
  const midY = h / 2;

  const phi = (2 * Math.PI) /  entityType.fields.length;


  const fields = entityType.fields.map((field, i) => {
    const x = midX + w * 0.35 * Math.cos(phi*i);
    const y = midY + h * 0.35 * Math.sin(phi*i);
    return (
      <text textAnchor="middle" key={i} x={x} y={y}>{field.name}</text>
    );
  });

  const lines =  entityType.fields.map((field, i) => {
    const x = midX + w * 0.35 * Math.cos(phi*i);
    const y = midY + h * 0.35 * Math.sin(phi*i);
    return (
      <path d={`M ${midX} ${midY} L ${x} ${y}`} />
    );
  });

  return (
    <div className="scaling-svg-container">
      <svg x="0"
           y="0"
           width={`${w}px`}
           height={`${h}px`}
           viewBox={`0 0 ${w} ${h}`}
           preserveAspectRatio="none">

        {lines}
        {fields}

        <text textAnchor="middle"
              x={midX}
              y={midY}>{`${entityType.entity}:${entityType['type']}`}</text>
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