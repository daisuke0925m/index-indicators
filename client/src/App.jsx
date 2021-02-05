import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import './style/index.css';
import { CustomizedSnackbar } from './components/UiKits';
import { createMuiTheme } from '@material-ui/core/styles';
import { getSignedIn, getUserID, getUsersLikes } from './redux/users/selectors';
import Header from './components/Header/Header';
import Main from './components/Main/Main';
import { ThemeProvider } from '@material-ui/core/styles';
import { fetchUsersLikes, listenAuthState } from './redux/users/operations';

const theme = createMuiTheme({
    palette: {
        primary: {
            main: '#447B64',
        },
    },
});

const App = () => {
    const dispatch = useDispatch();
    const selector = useSelector((state) => state);
    const isSignedIn = getSignedIn(selector);
    const userID = getUserID(selector);
    const likes = getUsersLikes(selector);
    console.log('isSignedIn', isSignedIn);
    console.log('userID', userID);
    console.log('likes', likes);

    useEffect(() => {
        dispatch(listenAuthState());
        if (userID) {
            dispatch(fetchUsersLikes(userID));
        }
    }, [userID]);

    return (
        <ThemeProvider theme={theme}>
            <CustomizedSnackbar />
            <Header />
            <Main />
        </ThemeProvider>
    );
};

export default App;
