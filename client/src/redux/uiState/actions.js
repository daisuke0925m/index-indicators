export const ALERT_OPEN = 'ALERT_OPEN';
export const alertOpenAction = (uiState) => {
    return {
        type: 'ALERT_OPEN',
        payload: uiState,
    };
};

export const ALERT_CLOSE = 'ALERT_CLOSE';
export const alertCloseAction = (uiState) => {
    return {
        type: 'ALERT_CLOSE',
        payload: uiState,
    };
};
