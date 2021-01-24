import React from 'react';
import PropTypes from 'prop-types';

const SpaceRow = (props) => {
    SpaceRow.propTypes = {
        height: PropTypes.number,
    };

    return <div style={{ margin: props.height }} />;
};

export default SpaceRow;
