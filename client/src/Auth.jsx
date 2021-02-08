import { useSelector } from 'react-redux';
import { getSignedIn } from './redux/users/selectors';
import PropTypes from 'prop-types';

const Auth = (props) => {
    Auth.propTypes = {
        enableEle: PropTypes.element,
        disableEle: PropTypes.element,
    };

    const selector = useSelector((state) => state);
    const isSignedIn = getSignedIn(selector);
    console.log(isSignedIn);

    return <div>{isSignedIn ? props.enableEle : props.disableEle}</div>;
};

export default Auth;
