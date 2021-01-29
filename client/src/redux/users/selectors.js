import { createSelector } from 'reselect';

const usersSelector = (state) => state.users;

export const getSignedIn = createSelector([usersSelector], (state) => state.isSignedIn);
