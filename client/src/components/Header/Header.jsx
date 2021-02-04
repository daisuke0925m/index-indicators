import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getSignedIn, getUserName } from '../../redux/users/selectors';
import AccountCircle from '@material-ui/icons/AccountCircle';
import { AppBar, Button, Divider } from '@material-ui/core';
import IconButton from '@material-ui/core/IconButton';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import { makeStyles } from '@material-ui/core/styles';
import { signOut } from '../../redux/users/operations';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Title from '../../assets/img/Index_logo.svg';
import { SignInModalForm, SignUpModalForm, DeleteUserModalForm, UpdateModalForm } from '../Users/index';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
        lineHeight: '50%',
    },
}));

const Header = () => {
    const classes = useStyles();

    const dispatch = useDispatch();
    const selector = useSelector((state) => state);
    const isSignedIn = getSignedIn(selector);
    const userName = getUserName(selector);

    const [anchorEl, setAnchorEl] = React.useState(null);
    const open = Boolean(anchorEl);

    const handleMenu = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    return (
        <div className={classes.root}>
            <AppBar position="static" color="default">
                <Toolbar>
                    <Typography variant="h6" className={classes.title} color="primary">
                        <img src={Title} alt="title img" />
                    </Typography>
                    <div>
                        {isSignedIn ? (
                            <>
                                <IconButton
                                    aria-label="account of current user"
                                    aria-controls="menu-appbar"
                                    aria-haspopup="true"
                                    onClick={handleMenu}
                                    color="primary"
                                >
                                    <AccountCircle />
                                </IconButton>
                                <Menu
                                    id="menu-appbar"
                                    anchorEl={anchorEl}
                                    anchorOrigin={{
                                        vertical: 'top',
                                        horizontal: 'right',
                                    }}
                                    keepMounted
                                    transformOrigin={{
                                        vertical: 'top',
                                        horizontal: 'right',
                                    }}
                                    open={open}
                                    onClose={handleClose}
                                >
                                    <MenuItem disabled>{userName}</MenuItem>
                                    <Divider />
                                    <MenuItem onClick={handleClose}>
                                        <span onClick={() => dispatch(signOut())}>Sign Out</span>
                                    </MenuItem>
                                    <MenuItem onClick={handleClose}>
                                        <UpdateModalForm text={'Update User'} />
                                    </MenuItem>
                                    <MenuItem onClick={handleClose}>
                                        <DeleteUserModalForm text={'Delete User'} />
                                    </MenuItem>
                                </Menu>
                            </>
                        ) : (
                            <div style={{ display: 'flex' }}>
                                <Button>
                                    <SignInModalForm text={'Sign In'} />
                                </Button>
                                <Button variant="contained" color="primary" style={{ marginLeft: 10 }}>
                                    <SignUpModalForm text={'Sign Up'} />
                                </Button>
                            </div>
                        )}
                    </div>
                </Toolbar>
            </AppBar>
        </div>
    );
};

export default Header;
