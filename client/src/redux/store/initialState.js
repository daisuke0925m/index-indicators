const initialState = {
    users: {
        isSignedIn: false,
        username: '',
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
