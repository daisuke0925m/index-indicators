import { createSelector } from 'reselect';

const tickersSelector = (state) => state.tickers;

export const selectDailyData = createSelector([tickersSelector], (state) => state);
