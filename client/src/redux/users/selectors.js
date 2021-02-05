import { createSelector } from 'reselect';

const usersSelector = (state) => state.users;

export const getSignedIn = createSelector([usersSelector], (state) => state.isSignedIn);

export const getUserID = createSelector([usersSelector], (state) => state.userID);

export const getUserName = createSelector([usersSelector], (state) => state.userName);

export const getUsersLikes = createSelector([usersSelector], (state) => state.likes);
