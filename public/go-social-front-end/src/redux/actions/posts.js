// import { browserHistory } from 'react-router';
import axios from 'axios'

function success(response) {
    // setToken( "token", response.token );
    return { type: response.statusData, message: response.message, data: response.data }
}

function failed(error) {
    return { error: error, type: error.statusData, message: error ? error.message : "Something wrong happened" }
}

export const homepage = (authToken) => {
    return dispatch => {
        const inputs = {
            limit: 1000,
            IDs: []
        }
        axios({
            method: 'post',
            url: 'http://localhost:1337/homepage',
            headers: {
                'Authorization': authToken
            },
            data: inputs
        }).then((resp) => {
            if (resp.status === 200) {
                resp.data.statusData = "homepageSuccess"
                return dispatch(success(resp.data));
            } else {
                resp.data.statusData = "homepageError"
                return dispatch(failed(resp.data))
            }
        }).catch((err) => {
            console.log(err.response)
            err.statusData = 'homepageError'
            return dispatch(failed(err))
        })
    }
}

export const postData = (text, authToken) => {
    return dispatch => {
        const inputs = {
            text
        }
        axios({
            method: 'put',
            url: 'http://localhost:1337/post',
            headers: {
                'Authorization': authToken
            },
            data: inputs
        }).then((resp) => {
            if (resp.status === 200) {
                resp.data.statusData = "postSuccess"
                return dispatch(success(resp.data));
            } else {
                resp.data.statusData = "homepageError"
                return dispatch(failed(resp.data))
            }
        }).catch((err) => {
            console.log(err.response)
            err.statusData = 'homepageError'
            return dispatch(failed(err))
        })
    }
}
export const VisibilityFilters = {
    homepageSuccess: "homepageSuccess",
    homepageError: "homepageError",
    postSuccess: "postSuccess"
}