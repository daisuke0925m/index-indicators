const initialState = {
    users: {
        isSignedIn: false,
        username: '',
    },
    uiState: {
        isModalOpen: false,
        alert: {
            isOpen: false,
            type: 'success',
            message: 'error',
        },
    },
};

export default initialState;
