import React, { useCallback, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import Backdrop from '@material-ui/core/Backdrop';
import Fade from '@material-ui/core/Fade';
import { getAlertState } from '../../redux/uiState/selectors';
import PropTypes from 'prop-types';
import { TextInput } from '../UiKits';
import { Button } from '@material-ui/core';
import { useDispatch, useSelector } from 'react-redux';
import { signUp } from '../../redux/users/operations';

const useStyles = makeStyles((theme) => ({
    modal: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
    paper: {
        backgroundColor: theme.palette.background.paper,
        border: '2px solid #000',
        boxShadow: theme.shadows[5],
        padding: theme.spacing(2, 4, 3),
    },
}));

const SignUpModalForm = (props) => {
    SignUpModalForm.propTypes = {
        text: PropTypes.string,
    };
    const classes = useStyles();
    const dispatch = useDispatch();
    const [open, setOpen] = useState(false);
    const [email, setEmail] = useState('');
    const [pass, setPass] = useState('');
    const [conformPass, setConformPass] = useState('');
    const [userName, setUserName] = useState('');

    const selector = useSelector((state) => state);
    const alertState = getAlertState(selector);

    const inputEmail = useCallback(
        (event) => {
            setEmail(event.target.value);
        },
        [setEmail]
    );

    const inputPass = useCallback(
        (event) => {
            setPass(event.target.value);
        },
        [setPass]
    );

    const inputConformPass = useCallback(
        (event) => {
            setConformPass(event.target.value);
        },
        [setPass]
    );

    const inputUserName = useCallback(
        (event) => {
            setUserName(event.target.value);
        },
        [setUserName]
    );

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    useEffect(() => {
        if (alertState.isOpen == true) {
            setOpen(false);
        }
    }, [alertState]);

    return (
        <div>
            <p onClick={handleOpen} style={{ margin: 0 }}>
                {props.text}
            </p>
            <Modal
                aria-labelledby="transition-modal-title"
                aria-describedby="transition-modal-description"
                className={classes.modal}
                open={open}
                onClose={handleClose}
                closeAfterTransition
                BackdropComponent={Backdrop}
                BackdropProps={{
                    timeout: 500,
                }}
            >
                <Fade in={open}>
                    <div className={classes.paper}>
                        <h2 id="transition-modal-title">{props.text}</h2>
                        <div id="transition-modal-description" style={{ textAlign: 'center' }}>
                            <TextInput
                                fullWidth={true}
                                label={'User Name'}
                                multiline={false}
                                onChange={inputUserName}
                                required={true}
                                rows={1}
                                value={userName}
                                type={'text'}
                            />
                            <TextInput
                                fullWidth={true}
                                label={'Email'}
                                multiline={false}
                                onChange={inputEmail}
                                required={true}
                                rows={1}
                                value={email}
                                type={'text'}
                            />
                            <TextInput
                                fullWidth={true}
                                label={'Password'}
                                multiline={false}
                                onChange={inputPass}
                                required={true}
                                rows={1}
                                value={pass}
                                type={'password'}
                            />
                            <TextInput
                                fullWidth={true}
                                label={'Conform Password'}
                                multiline={false}
                                onChange={inputConformPass}
                                required={true}
                                rows={1}
                                value={conformPass}
                                type={'password'}
                            />
                            <Button
                                color="primary"
                                variant={'outlined'}
                                onClick={() => dispatch(signUp(userName, email, pass, conformPass))}
                            >
                                Sign Up
                            </Button>
                        </div>
                    </div>
                </Fade>
            </Modal>
        </div>
    );
};

export default SignUpModalForm;
