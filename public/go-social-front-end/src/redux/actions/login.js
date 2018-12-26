import { setItem, removeItem } from '../helpers/storage';
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
    console.log("INSIDE LOGIN REDUX")
    console.log(inputs)
    // not integrated with backend later integrate with axios
    return dispatch => {
        
    }
}

export const signUp = (inputs) => {
    // not integrated with backend later integrate with axios
    return dispatch => {
        axios.put('http://localhost:1337/user/signup', {
            ...inputs
        }).then(function (resp) {
            console.log(resp)
            if (resp.status === 200) {
                // setItem('authToken', resp.data.authToken);
                browserHistory.push('/verify')
                return dispatch(signupSuccess(resp.data));
            } else {
                return dispatch(signupFailure(resp.data))
            }
        }).catch(function (err) {
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