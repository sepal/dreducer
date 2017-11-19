import React from 'react';
import ReactDOM from 'react-dom';
import ApolloClient from 'apollo-client'
import {HttpLink, InMemoryCache} from 'apollo-client-preset'
import {ApolloProvider} from 'react-apollo'

import './index.css';
import App from './components/App';
import registerServiceWorker from './registerServiceWorker';

const client = new ApolloClient({
  link: new HttpLink({uri: '/graphql'}),
  cache: new InMemoryCache().restore({})
});

ReactDOM.render(
  <ApolloProvider client={client}>
    <App/>
  </ApolloProvider>,
  document.getElementById('root')
);
registerServiceWorker();
