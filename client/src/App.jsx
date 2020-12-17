import React from 'react'
import { ThemeProvider } from '@material-ui/core/styles';
import './style/index.css'
import Fgi from './components/Fgi/Fgi'
import Header from './components/Header/Header'
import { createMuiTheme } from '@material-ui/core/styles';

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
                <Fgi />
            </div>
        </ThemeProvider>
    );
}

export default App;
