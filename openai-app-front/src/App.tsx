import React from 'react';
import { Authenticator } from '@aws-amplify/ui-react';
import '@aws-amplify/ui-react/styles.css';

import './App.css';
import Home from './Pages/Home';

function App() {
  return (
    <Authenticator signUpAttributes={['email']}>
      {({ signOut, user }) => (
        <Home />
      )}
    </Authenticator>
  );
}

export default App;