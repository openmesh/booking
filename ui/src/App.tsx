import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Home } from './home/components';
import { Signup } from './signup/components';

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/" exact>
          <Home />
        </Route>
        <Route path="/signup">
          <Signup />
        </Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
