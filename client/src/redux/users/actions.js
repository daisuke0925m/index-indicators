export const SIGN_IN = 'SIGN_IN';
export const signInAction = (userState) => {
    return {
        type: 'SIGN_IN',
        payload: userState,
    };
};

export const SIGN_OUT = 'SIGN_OUT';
export const signOutAction = () => {
    return {
        type: 'SIGN_OUT',
        payload: {
            isSignedIn: false,
            userName: '',
            userID: 0,
            likes: [],
        },
    };
};

export const FETCH_LIKES = 'FETCH_LIKES';
export const fetchLikesAction = (likes) => {
    return {
        type: 'FETCH_LIKES',
        payload: likes,
    };
};
