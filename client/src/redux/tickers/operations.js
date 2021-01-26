import { getAllDailyDataAction } from './actions';
import axios from 'axios';

export const getAllDailyData = (symbol) => {
    return async (dispatch) => {
        try {
            const response = await axios.get(`/ticker?symbol=${symbol}`);
            dispatch(getAllDailyDataAction(response.data));
        } catch (error) {
            console.log(error);
        }
    };
};
