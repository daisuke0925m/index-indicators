import React, { useEffect } from 'react';
import { useSelector } from 'react-redux';
import './style/index.css';
import { createMuiTheme } from '@material-ui/core/styles';
import { getSignedIn } from './redux/users/selectors';
import Header from './components/Header/Header';
import Main from './components/Main/Main';
import { ThemeProvider } from '@material-ui/core/styles';

const theme = createMuiTheme({
    palette: {
        primary: {
            main: '#447B64',
        },
    },
});

const App = () => {
    // const dispatch = useDispatch();
    const selector = useSelector((state) => state);
    const isSignedIn = getSignedIn(selector);
    console.log(isSignedIn);

    useEffect(() => {
        // dispatch(listenAuthState());
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
