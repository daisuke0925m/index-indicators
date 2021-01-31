import { createSelector } from 'reselect';

const uiStateSelector = (state) => state.uiState;

export const getIsModalOpen = createSelector([uiStateSelector], (state) => state.isModalOpen);
