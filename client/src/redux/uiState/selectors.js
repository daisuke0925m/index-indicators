import { createSelector } from 'reselect';

const uiStateSelector = (state) => state.uiState;

export const getIsModalOpen = createSelector([uiStateSelector], (state) => state.isModalOpen);

export const getAlertState = createSelector([uiStateSelector], (state) => state.alert);
