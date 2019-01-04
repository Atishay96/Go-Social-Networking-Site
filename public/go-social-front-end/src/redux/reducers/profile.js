import { VisibilityFilters } from '../actions/profile';

const { profileSuccess, profileError } = VisibilityFilters;

const initialState = {
    userData: {
        error: '',
        data: []
    }
}

export const homepage = (state = initialState, action) => {
    switch (action.type) {
        case profileSuccess:
            return { ...state, userData: { error: '', data: action.data } }
        case profileError:
            return { ...state, userData: { error: action.message, data: [] } }
        default:
            return state;
    }
}