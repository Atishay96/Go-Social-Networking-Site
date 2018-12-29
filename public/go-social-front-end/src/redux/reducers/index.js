import { combineReducers } from 'redux';
import { login } from './login'
import { homepage } from './posts'

export default combineReducers({
    login,
    homepage
})