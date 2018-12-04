import React, { Component } from 'react';
import { Router, Route, browserHistory } from 'react-router';

import landingPage from './components/landingPage';
import loginPage from './components/login';

class App extends Component {
  render() {
    return (
      <Router history = { browserHistory }>
        <Route path ='/' component = { landingPage }></Route>
        <Route path ='/login' component = { loginPage }></Route>
      </Router>
    );
  }
}

export default App;