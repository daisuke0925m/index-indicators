import * as Actions from './actions';
import initialState from '../store/initialState';

export const TickersReducer = (state = initialState.tickers, action) => {
    switch (action.type) {
        case Actions.GET_ALL_DAILY_DATA:
            return {
                ...state,
                ...action.payload,
            };
        default:
            return state;
    }
};
