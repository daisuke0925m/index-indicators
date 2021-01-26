export const GET_ALL_DAILY_DATA = 'GET_ALL_DAILY_DATA';
export const getAllDailyDataAction = (tickers) => {
    return {
        type: 'GET_ALL_DAILY_DATA',
        payload: tickers,
    };
};
