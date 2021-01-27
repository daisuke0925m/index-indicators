import { createStore as reduxCreateStore, combineReducers, applyMiddleware } from 'redux';
import { createLogger } from 'redux-logger';
import { connectRouter, routerMiddleware } from 'connected-react-router';
import thunk from 'redux-thunk';

export default function createStore(history) {
    const logger = createLogger({
        collapsed: true,
        diff: true,
    });

    return reduxCreateStore(
        combineReducers({
            router: connectRouter(history),
        }),
        applyMiddleware(logger, routerMiddleware(history), thunk)
    );
}
