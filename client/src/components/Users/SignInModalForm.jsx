import React, { useCallback, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import Backdrop from '@material-ui/core/Backdrop';
import Fade from '@material-ui/core/Fade';
import PropTypes from 'prop-types';
import { TextInput } from '../UiKits';
import { Button } from '@material-ui/core';
import { useDispatch } from 'react-redux';
import { signIn } from '../../redux/users/operations';

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

const SignInModalForm = (props) => {
    SignInModalForm.propTypes = {
        text: PropTypes.string,
    };
    const classes = useStyles();
    const dispatch = useDispatch();
    const [open, setOpen] = useState(false);
    const [email, setEmail] = useState('');
    const [pass, setPass] = useState('');

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

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

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
                                label={'Email'}
                                multiline={false}
                                onChange={inputEmail}
                                required={false}
                                rows={1}
                                value={email}
                                type={'text'}
                            />
                            <TextInput
                                fullWidth={true}
                                label={'Password'}
                                multiline={false}
                                onChange={inputPass}
                                required={false}
                                rows={1}
                                value={pass}
                                type={'password'}
                            />
                            <Button color="primary" variant={'outlined'} onClick={() => dispatch(signIn(email, pass))}>
                                Sign In
                            </Button>
                        </div>
                    </div>
                </Fade>
            </Modal>
        </div>
    );
};

export default SignInModalForm;
