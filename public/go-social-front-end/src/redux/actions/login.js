import { setItem } from '../helpers/storage';
import { browserHistory } from 'react-router';
import axios from 'axios'

function loginSuccess(response) {
    // setToken( "token", response.token );
    return { type: 'LOGIN_SUCCESS', message: response.message }
}

function loginFailure(error) {
    return { error: error || 'Login failed', type: 'LOGIN_FAILED', message: error ? error.message : "Login Failed" }
}

function signupSuccess(response) {
    return { type: 'SIGNUP_SUCCESS', message: response.message }
}

function signupFailure(error) {
    console.log("Sign up failure")
    return { error: error || 'Signup failed', type: 'SIGNUP_FAILED', message: error ? error.message : "Signup Failed" }
}

export const login = (inputs) => {
    return dispatch => {
        axios.post('http://localhost:1337/user/login', {
            ...inputs
        }).then((resp) => {
            console.log(resp)
            if (resp.status === 200) {
                setItem('AuthToken', resp.data.data.AuthToken);
                let expiresOn = new Date()
                expiresOn.setHours(expiresOn.getHours() + 3)
                if (inputs.rememberMe) {
                    expiresOn.setHours(expiresOn.getDate() + 30)
                }
                setItem('expiresOn', expiresOn)
                browserHistory.push('/')
                return dispatch(loginSuccess(resp.data.data));
            } else {
                return dispatch(loginFailure(resp.data.data))
            }
        }).catch((err) => {
            console.log(err.response)
            if (!err.response) {
                return dispatch(loginFailure(err))
            } else {
                err.message = err.response.data.message
                return dispatch(loginFailure(err))
            }
        })
    }
}

export const signUp = (inputs) => {
    return dispatch => {
        axios.put('http://localhost:1337/user/signup', {
            ...inputs
        }).then((resp) => {
            console.log(resp)
            if (resp.status === 200) {
                // setItem('authToken', resp.data.authToken);
                browserHistory.push('/verify')
                return dispatch(signupSuccess(resp.data.data));
            } else {
                return dispatch(signupFailure(resp.data.data))
            }
        }).catch((err) => {
            console.log(err.response)
            if (!err.response) {
                return dispatch(signupFailure(err))
            } else {
                err.message = err.response.data.message
                return dispatch(signupFailure(err))
            }
        })
    }
}


export const VisibilityFilters = {
    LOGIN_SUCCESS: 'LOGIN_SUCCESS',
    LOGIN_FAILED: 'LOGIN_FAILED',
    SIGNUP_SUCCESS: 'SIGNUP_SUCCESS',
    SIGNUP_FAILED: 'SIGNUP_FAILED'
}