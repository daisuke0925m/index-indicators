import React, { useCallback, useState } from 'react';
import { Button } from '@material-ui/core';
import Backdrop from '@material-ui/core/Backdrop';
import { deleteUser } from '../../redux/users/operations';
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

const DeleteUserModalForm = (props) => {
    DeleteUserModalForm.propTypes = {
        text: PropTypes.string,
    };
    const classes = useStyles();
    const dispatch = useDispatch();
    const [open, setOpen] = useState(false);
    const [pass, setPass] = useState('');
    const selector = useSelector((state) => state);
    const userID = getUserID(selector);

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
                                label={'Password'}
                                multiline={false}
                                onChange={inputPass}
                                required={false}
                                rows={1}
                                value={pass}
                                type={'password'}
                            />
                            <Button
                                color="primary"
                                variant={'outlined'}
                                onClick={() => dispatch(deleteUser(pass, userID))}
                            >
                                Delete User
                            </Button>
                        </div>
                    </div>
                </Fade>
            </Modal>
        </div>
    );
};

export default DeleteUserModalForm;
