import * as Actions from './actions';
import initialState from '../store/initialState';

export const UiStateReducer = (state = initialState.uiState, action) => {
    switch (action.type) {
        case Actions.ALERT_OPEN:
            return {
                ...action.payload,
            };
        case Actions.ALERT_CLOSE:
            return {
                ...action.payload,
            };
        default:
            return state;
    }
};
