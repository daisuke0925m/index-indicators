import React from "react";
import CheckIcon from '@material-ui/icons/Check';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundColor: '#E7E7E7',
        borderRadius: 18,
        padding: 21,
    },
    title: {
        fontFamily: 'Arial',
        fontSize: 16,
        fontWeight: 'bold',
        height: 23,
        display: 'flex',
    },
    secondPaper: {
        padding: 10,
        borderColor: '#707070'
    }
}));

const CntWrap = (props) => {
    const classes = useStyles();

    return (
        <Paper elevation={0} classes={{ root: classes.root }} >
            <div className={classes.title}>
                <CheckIcon style={{ fontSize: 18 }} />
                {props.title}
            </div>
            <Paper classes={{ root: classes.secondPaper }} elevation={0} square variant="outlined">
                {props.children}
            </Paper>
        </Paper>
    )
};

export default CntWrap