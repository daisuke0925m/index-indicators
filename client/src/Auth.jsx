import { useSelector } from 'react-redux';
import { getSignedIn } from './redux/users/selectors';
import PropTypes from 'prop-types';

const Auth = (props) => {
    Auth.propTypes = {
        children: PropTypes.element,
    };

    const selector = useSelector((state) => state);
    const isSignedIn = getSignedIn(selector);

    return <div>{isSignedIn ? props.children : 'ログインしているーザーのみ使用できます。'}</div>;
};

export default Auth;
