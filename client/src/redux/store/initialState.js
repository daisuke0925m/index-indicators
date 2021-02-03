const initialState = {
    users: {
        userID: 0,
        userName: '',
        isSignedIn: false,
    },
    uiState: {
        alert: {
            isOpen: false,
            type: 'success',
            message: 'error',
        },
    },
};

export default initialState;
