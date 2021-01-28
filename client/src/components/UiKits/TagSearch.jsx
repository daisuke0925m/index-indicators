import React, { useEffect, useState } from 'react';
import Autocomplete from '@material-ui/lab/Autocomplete';
import PropTypes from 'prop-types';
// import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';

// const useStyles = makeStyles((theme) => ({
//     root: {
//         width: 500,
//         '& > * + *': {
//             marginTop: theme.spacing(3),
//         },
//     },
// }));

const tickers = ['spy', 'spxl', '^skew', 'tlt', 'gld', 'gldm'];

const TagSearch = (props) => {
    TagSearch.propTypes = {
        setKeyword: PropTypes.func,
    };
    const setKeyword = props.setKeyword;
    const [value, setValue] = useState([tickers[0]]);

    useEffect(() => {
        setKeyword(value);
    }, [value]);

    return (
        <Autocomplete
            onChange={(event, newValue) => {
                setValue(newValue);
            }}
            multiple
            id="tags-outlined"
            options={tickers}
            getOptionLabel={(option) => option}
            defaultValue={[tickers[0]]}
            filterSelectedOptions
            renderInput={(params) => (
                <TextField {...params} variant="outlined" label="Search Tickers" placeholder="Add" />
            )}
        />
    );
};

export default TagSearch;
