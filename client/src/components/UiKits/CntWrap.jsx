import React from "react";
import CheckIcon from '@material-ui/icons/Check';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import SpaceRow from "./SpaceRow";
import Accordion from '@material-ui/core/Accordion';
import AccordionSummary from '@material-ui/core/AccordionSummary';
import AccordionDetails from '@material-ui/core/AccordionDetails';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

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
    },
    accordionRoot: {
        width: '100%',
    },
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
            <SpaceRow height={20} />
            <div className={classes.accordionRoot}>
                <Accordion>
                    <AccordionSummary
                        expandIcon={<ExpandMoreIcon />}
                    >
                        <h5 style={{ margin: 0 }}>
                            {props.accordionHead}
                        </h5>
                    </AccordionSummary>
                    <AccordionDetails>
                        {props.description}
                    </AccordionDetails>
                </Accordion>
            </div>
        </Paper>
    )
};

export default CntWrap