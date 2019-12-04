import {applyMiddleware, combineReducers, createStore} from "redux";
import thunk from "redux-thunk";

import workSession from "stores/workSession";

export default createStore(combineReducers(
    { workSession },
), {}, applyMiddleware(thunk));

