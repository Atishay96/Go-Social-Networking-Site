import React, { Component } from 'react';
import { Router, Route, browserHistory } from 'react-router';

import landingPage from './components/landingPage';
import loginPage from './components/login';
import signUpPage from './components/signup';
import verifyPage from './components/verifyPage';
import profilePage from './components/profile';

class App extends Component {
  render() {
    return (
      <Router history={browserHistory}>
        <Route path='/' component={landingPage}></Route>
        <Route path='/profile' component={profilePage}></Route>
        <Route path='/profile/:id' component={profilePage}></Route>
        <Route path='/login' component={loginPage}></Route>
        <Route path='/signup' component={signUpPage}></Route>
        <Route path='/verify' component={verifyPage}></Route>
      </Router>
    );
  }
}

export default App;