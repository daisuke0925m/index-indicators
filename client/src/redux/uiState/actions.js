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

export const ALERT_OPEN = 'ALERT_OPEN';
export const alertOpenAction = (uiState) => {
    return {
        type: 'ALERT_OPEN',
        payload: uiState,
    };
};

export const ALERT_CLOSE = 'ALERT_CLOSE';
export const alertCloseAction = () => {
    return {
        type: 'MODAL_CLOSE',
        payload: {
            alert: {
                isOpen: false,
                type: 'success',
                message: '',
            },
        },
    };
};
