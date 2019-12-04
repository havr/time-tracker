import * as React from "react";
import * as ReactDOM from "react-dom";
import {
    HashRouter as Router,
    Switch,
    Route,
} from "react-router-dom";

import { Provider } from "react-redux";
import MainPage from "components/pages/MainPage/MainPage";
import store from "stores/combined";

class Index extends React.Component<{}> {
    render() {
        return (
            <Provider store={store}>
                <Router>
                    <Switch>
                        <Route path={"/"}> <MainPage/> </Route>
                    </Switch>
                </Router>
            </Provider>
        )
    }
}

ReactDOM.render(<Index />, document.getElementById("root"),);

