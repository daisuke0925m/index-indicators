import React, { useCallback, useState } from 'react';
import { Button } from '@material-ui/core';
import Backdrop from '@material-ui/core/Backdrop';
import { updateUser } from '../../redux/users/operations';
import Fade from '@material-ui/core/Fade';
import { getUserID } from '../../redux/users/selectors';
import { makeStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import PropTypes from 'prop-types';
import { TextInput } from '../UiKits';
import { useDispatch, useSelector } from 'react-redux';

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

const UpdateModalForm = (props) => {
    UpdateModalForm.propTypes = {
        text: PropTypes.string,
    };
    const classes = useStyles();
    const dispatch = useDispatch();
    const [newName, setNewName] = useState('');
    const [newEmail, setNewEmail] = useState('');
    const [newPass, setNewPass] = useState('');
    const [open, setOpen] = useState(false);
    const [pass, setPass] = useState('');
    const selector = useSelector((state) => state);
    const userID = getUserID(selector);

    const inputNewName = useCallback(
        (event) => {
            setNewName(event.target.value);
        },
        [setNewName]
    );

    const inputNewEmail = useCallback(
        (event) => {
            setNewEmail(event.target.value);
        },
        [setNewEmail]
    );

    const inputNewPass = useCallback(
        (event) => {
            setNewPass(event.target.value);
        },
        [setNewPass]
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
                                label={'Change User Name'}
                                multiline={false}
                                onChange={inputNewName}
                                required={false}
                                rows={1}
                                value={newName}
                                type={'text'}
                            />
                            <TextInput
                                fullWidth={true}
                                label={'Change Email'}
                                multiline={false}
                                onChange={inputNewEmail}
                                required={false}
                                rows={1}
                                value={newEmail}
                                type={'text'}
                            />
                            <TextInput
                                fullWidth={true}
                                label={'Change Email'}
                                multiline={false}
                                onChange={inputNewPass}
                                required={false}
                                rows={1}
                                value={newPass}
                                type={'text'}
                            />
                            <TextInput
                                fullWidth={true}
                                label={'Current Password'}
                                multiline={false}
                                onChange={inputPass}
                                required={true}
                                rows={1}
                                value={pass}
                                type={'password'}
                            />
                            <Button
                                color="primary"
                                variant={'outlined'}
                                onClick={() => dispatch(updateUser(newName, newEmail, newPass, pass, userID))}
                            >
                                Update User
                            </Button>
                        </div>
                    </div>
                </Fade>
            </Modal>
        </div>
    );
};

export default UpdateModalForm;
