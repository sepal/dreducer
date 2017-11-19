import React from 'react';

function Graph(props) {
  return (
    <svg version="1.1"
         baseProfile="full"
         preserveAspectRatio="none"
         viewBox="0 0 5 3">
         xmlns="http://www.w3.org/2000/svg">
      {props.children}
    </svg>
  );
}