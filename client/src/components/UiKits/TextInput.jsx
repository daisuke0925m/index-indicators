import React from 'react';
import PropTypes from 'prop-types';
import TextField from '@material-ui/core/TextField';

const TextInput = (props) => {
    TextInput.propTypes = {
        fullWidth: PropTypes.bool,
        label: PropTypes.string,
        multiline: PropTypes.bool,
        required: PropTypes.bool,
        rows: PropTypes.number,
        value: PropTypes.string,
        type: PropTypes.string,
        onChange: PropTypes.func,
    };
    return (
        <TextField
            fullWidth={props.fullWidth}
            label={props.label}
            margin="dense"
            multiline={props.multiline}
            required={props.required}
            rows={props.rows}
            value={props.value}
            type={props.type}
            onChange={props.onChange}
        />
    );
};

export default TextInput;
