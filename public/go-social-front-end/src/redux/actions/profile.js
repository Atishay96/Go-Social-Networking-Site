import axios from 'axios'

function success(response) {
    return { type: response.statusData, message: response.message, data: response.data }
}

function failed(error) {
    return { error: error, type: error.statusData, message: error ? error.message : "Something wrong happened" }
}

export const getUser = (authToken, userId) => {
    return dispatch => {
        axios({
            method: 'get',
            url: 'http://localhost:1337/user/profile/' + userId,
            headers: {
                'Authorization': authToken
            }
        }).then((resp) => {
            if (resp.status === 200) {
                resp.data.statusData = "profileSuccess"
                return dispatch(success(resp.data));
            } else {
                resp.data.statusData = "profileError"
                return dispatch(failed(resp.data))
            }
        }).catch((err) => {
            console.log(err.response)
            err.statusData = 'profileError'
            return dispatch(failed(err))
        })
    }
}

export const VisibilityFilters = {
    profileSuccess: "profileSuccess",
    profileError: "profileError"
}