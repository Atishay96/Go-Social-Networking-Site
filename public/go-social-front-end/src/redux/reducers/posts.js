import { VisibilityFilters } from '../actions/posts';

const { homepageSuccess, homepageError, postSuccess } = VisibilityFilters;

const initialState = {
    homepageData: {
        error: '',
        data: []
    }
}

export const homepage = (state = initialState, action) => {
    switch (action.type) {
        case homepageSuccess:
            return { ...state, homepageData: { error: '', data: action.data } };
        case homepageError:
            return { ...state, homepageData: { error: action.message, data: [] } };
        case postSuccess:
            return { ...state, homepageData: { error: '', data: action.data, post:true } }
        default:
            return state;
    }
}