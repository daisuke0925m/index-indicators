export const MODAL_OPEN = 'MODAL_OPEN';
export const modalOpenAction = (uiState) => {
    return {
        type: 'MODAL_OPEN',
        payload: uiState,
    };
};

export const MODAL_CLOSE = 'MODAL_CLOSE';
export const modalCloseAction = () => {
    return {
        type: 'MODAL_CLOSE',
        payload: {
            isModalOpen: false,
        },
    };
};
