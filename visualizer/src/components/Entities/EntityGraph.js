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

function FieldList({fields, midX, midY, rh, rv}) {
  if (fields == undefined) {
    return (
      <g></g>
    );
  }

  const phi = ( Math.PI) / fields.length;

  return fields.map((field, i) => {
    const x = midX + rh * Math.cos(phi*i);
    const y = midY + rv * Math.sin(phi*i);

    const sub_fields = (
      <FieldList fields={field.fields} midX={x} midY={y} rh={180} rv={180}/>
    );

    return (
      <g key={i} className="field">
        {sub_fields}
        <path d={`M ${midX} ${midY} L ${x} ${y}`} />
        <text textAnchor="middle" x={x} y={y}>{field.name}</text>
      </g>
    );
  });
}

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
  const midX = w / 2;
  const midY = 20;
  const fields = <FieldList fields={entityType.fields} midX={midX} midY={midY} rh={w*0.35} rv={h*0.35} />;

  return (
    <div className="scaling-svg-container">
      <svg x="0"
           y="0"
           width={`${w}px`}
           height={`${h}px`}
           viewBox={`0 0 ${w} ${h}`}
           preserveAspectRatio="none">
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