import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import './style/index.css';
import axios from 'axios';
import { createMuiTheme } from '@material-ui/core/styles';
import { getSignedIn } from './redux/users/selectors';
import Header from './components/Header/Header';
import Main from './components/Main/Main';
import { ThemeProvider } from '@material-ui/core/styles';
import { listenAuthState } from './redux/users/operations';

axios.defaults.baseURL = 'http://localhost:8080';
axios.defaults.headers.post['Content-Type'] = 'application/json';
axios.defaults.withCredentials = true;

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
    console.log(isSignedIn);

    useEffect(() => {
        dispatch(listenAuthState());
    }, []);

    return (
        <ThemeProvider theme={theme}>
            <div>
                <Header />
                <Main />
            </div>
        </ThemeProvider>
    );
};

export default App;
