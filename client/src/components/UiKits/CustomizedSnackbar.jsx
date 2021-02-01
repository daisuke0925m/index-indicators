import React from 'react';
import { alertCloseAction } from '../../redux/uiState/actions';
import { getAlertState } from '../../redux/uiState/selectors';
import MuiAlert from '@material-ui/lab/Alert';
import { makeStyles } from '@material-ui/core/styles';
import Snackbar from '@material-ui/core/Snackbar';
import { useDispatch, useSelector } from 'react-redux';

function Alert(props) {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        '& > * + *': {
            marginTop: theme.spacing(2),
        },
    },
}));

export default function CustomizedSnackbar() {
    const classes = useStyles();
    const dispatch = useDispatch();
    const selector = useSelector((state) => state);
    const alertState = getAlertState(selector);

    const handleClose = (event, reason) => {
        if (reason === 'clickaway') {
            return;
        }
        dispatch(
            alertCloseAction({
                alert: {
                    isOpen: false,
                    type: alertState.type,
                    message: '',
                },
            })
        );
    };

    return (
        <div className={classes.root}>
            <Snackbar
                open={alertState.isOpen}
                autoHideDuration={3000}
                onClose={handleClose}
                anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
            >
                <Alert onClose={handleClose} severity={alertState.type}>
                    {alertState.message}
                </Alert>
            </Snackbar>
        </div>
    );
}
