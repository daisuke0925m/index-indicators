import React from 'react';
import FormatQuoteIcon from '@material-ui/icons/FormatQuote';
import { makeStyles } from '@material-ui/core/styles';
import PropTypes from 'prop-types';

const useStyles = makeStyles(() => ({
    blockquote: {
        position: 'relative',
        padding: '35px 15px 10px 15px',
        boxSizing: 'border-box',
        fontStyle: 'italic',
        background: '#f5f5f5',
        color: '#777777',
        borderLeft: '4px solid #447B64',
        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.14)',
    },
    iconRoot: {
        transform: 'scale(-1, -1)',
    },
}));

const BlockQuote = (props) => {
    const classes = useStyles();

    BlockQuote.propTypes = {
        children: PropTypes.element,
    };

    return (
        <div className={classes.blockquote}>
            <FormatQuoteIcon className={classes.iconRoot} color="primary" />
            {props.children}
        </div>
    );
};

export default BlockQuote;
