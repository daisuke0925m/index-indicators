import React, { useEffect, useState } from 'react';
import { withStyles } from '@material-ui/core/styles';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import { likePost, likeDelete } from '../../redux/users/operations';
import { useDispatch, useSelector } from 'react-redux';
import { getUserID } from '../../redux/users/selectors';
import PropTypes from 'prop-types';

const StyledSwitch = withStyles((theme) => ({
    root: {
        width: 42,
        height: 26,
        padding: 0,
        margin: theme.spacing(1),
    },
    switchBase: {
        padding: 1,
        '&$checked': {
            transform: 'translateX(16px)',
            color: theme.palette.common.white,
            '& + $track': {
                backgroundColor: '#52d869',
                opacity: 1,
                border: 'none',
            },
        },
        '&$focusVisible $thumb': {
            color: '#52d869',
            border: '6px solid #fff',
        },
    },
    thumb: {
        width: 24,
        height: 24,
    },
    track: {
        borderRadius: 26 / 2,
        border: `1px solid ${theme.palette.grey[400]}`,
        backgroundColor: theme.palette.grey[50],
        opacity: 1,
        transition: theme.transitions.create(['background-color', 'border']),
    },
    checked: {},
    focusVisible: {},
}))(({ classes, ...props }) => {
    return (
        <Switch
            focusVisibleClassName={classes.focusVisible}
            disableRipple
            classes={{
                root: classes.root,
                switchBase: classes.switchBase,
                thumb: classes.thumb,
                track: classes.track,
                checked: classes.checked,
            }}
            {...props}
        />
    );
});

const LikeSwitch = (props) => {
    LikeSwitch.propTypes = {
        symbol: PropTypes.string,
        flag: PropTypes.bool,
        likeID: PropTypes.number,
    };

    const likeID = props.likeID;
    const flag = props.flag;
    const symbol = props.symbol;
    const dispatch = useDispatch();
    const selector = useSelector((state) => state);
    const userID = getUserID(selector);
    const [switchState, setSwitchState] = useState({
        checked: flag,
    });

    const handleChange = (event) => {
        setSwitchState({ ...switchState, [event.target.name]: event.target.checked });
        const checked = event.target.checked;
        if (checked) {
            dispatch(likePost(userID, symbol));
        } else if (!checked) {
            console.log('done');
            dispatch(likeDelete(userID, likeID, symbol));
        }
    };

    useEffect(() => {
        setSwitchState({ checked: flag });
    }, [flag]);
    return (
        <FormGroup>
            <FormControlLabel
                control={<StyledSwitch checked={switchState.checked} onChange={handleChange} name="checked" />}
                label=""
                style={{ margin: ' 0 0 0 auto' }}
            />
        </FormGroup>
    );
};

export default LikeSwitch;
