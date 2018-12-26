import { setItem, removeItem } from '../helpers/storage';
import { browserHistory } from 'react-router';
import axios from 'axios'

function loginSuccess(response) {
    // setToken( "token", response.token );
    return { type: 'LOGIN_SUCCESS', message: response.message }
}

function loginFailure(error) {
    return { error: error || 'Login failed', type: 'LOGIN_FAILED', message: error ? error.message : "Login Falied" }
}

export const login = (inputs) => {
    console.log("INSIDE LOGIN REDUX")
    console.log(inputs)
    // not integrated with backend later integrate with axios
    return dispatch => {
        if (inputs.email === 'admin@pokedex.com' && inputs.password === '123456') {
            console.log('correct creds');
            // static response
            let response = { message: "Successfully logged in", data: { authToken: '' } };
            setItem('authToken', 'asdasdasd');
            return dispatch(loginSuccess(response));
        }
        else {
            console.log('wrong creds');
            let error = { message: 'Wrong Credentials' };
            return dispatch(loginFailure(error));
        }
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
                setItem('authToken', resp.data.authToken);
                browserHistory.push('/')
                return dispatch(loginSuccess(resp.data));
            } else {
                return dispatch(loginFailure(resp.data.message))
            }
        }).catch(function (err) {
            if (!err.response) {
                return dispatch(loginFailure(err.response))
            }
            if (err.response.status === 401 || err.response.status === 403) {
                removeItem('authToken')
                return browserHistory.push('/login')
            }

            else
                return dispatch(loginFailure(err.message))
        })
    }
}


export const VisibilityFilters = {
    LOGIN_SUCCESS: 'LOGIN_SUCCESS',
    LOGIN_FAILED: 'LOGIN_FAILED'
}