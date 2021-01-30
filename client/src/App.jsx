import React from 'react';
import axios from 'axios';
import { ThemeProvider } from '@material-ui/core/styles';
import './style/index.css';
import Main from './components/Main/Main';
import Header from './components/Header/Header';
import { createMuiTheme } from '@material-ui/core/styles';

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
