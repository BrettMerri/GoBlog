import React from 'react';
import { Switch, Route } from 'react-router-dom';
import Home from './Home/Home';
import NotFound from './NotFound/NotFound';

const Main = () => (
  <div className="container main">
    <Switch>
      <Route exact path='/' component={Home}/>
      <Route component={NotFound} />
    </Switch>
  </div>
);

export default Main;