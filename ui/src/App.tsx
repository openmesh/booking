import React from 'react';
import './App.css';
import { Button } from 'antd';
import { Signin } from './signin';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Dashboard } from './dashboard';
import { AppLayout } from './common/components/app-layout';

function App() {
  return (
    <BrowserRouter>
      <Route path="/" exact>
        <Signin />
      </Route>
      <AppLayout>
        <Switch>
          <Route path="/dashboard">
            <Dashboard />
          </Route>
        </Switch>
      </AppLayout>
    </BrowserRouter>
  );
}

export default App;
