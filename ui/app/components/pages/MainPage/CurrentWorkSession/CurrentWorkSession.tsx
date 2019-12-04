import * as React from "react";

import Timer from "components/pages/MainPage/CurrentWorkSession/Timer/Timer";
import WorkSession from "models/WorkSession";

interface Props {
    value: WorkSession;
    onStart: () => any;
    onSave: () => any;
    onRename: (string) => any;
}

export default class extends React.Component<Props> {
    render() {
        const { value } = this.props;
        return (
            <div className={"d-flex flex-row align-items-center justify-content-space-between"}>
                <div className="col">
                    <input className="form-control" placeholder={"Session name"} value={value.name || ""} onChange={this.rename} />
                </div>
                <div className="align-self-center">
                    <Timer start={value.startTime} />
                </div>
                <div className="col">
                    <div className="d-flex flew-row">
                        {this.renderActionButton()}
                    </div>
                </div>
            </div>
        )
    }

    rename = (e) => {
        const { onRename } = this.props;

        onRename(e.target.value);
    };

    get tracking() {
        return !! this.props.value.startTime;
    }

    renderActionButton() {
        const title = this.tracking ? "Done" : "Start";
        const action = this.tracking ? this.props.onSave : this.props.onStart;
        return (
            <button className="btn btn-primary" onClick={action}>
                {title}
            </button>
        );
    }
}