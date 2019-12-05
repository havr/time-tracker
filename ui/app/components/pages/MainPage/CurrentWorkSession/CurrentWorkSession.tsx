import * as React from "react";

import Timer from "components/pages/MainPage/CurrentWorkSession/Timer/Timer";
import WorkSession from "models/WorkSession";
import texts from "constants/texts";

interface Props {
    value: WorkSession;
    onStart: () => void;
    onSave: () => void;
    onRename: (string) => void;
}

export default class CurrentWorkSession extends React.Component<Props> {
    render() {
        const { value } = this.props;
        return (
            <div className={"d-flex flex-row align-items-center justify-content-space-between"}>
                <div className="col">
                    <input className="form-control" placeholder={texts.workSession.namePlaceholder} value={value.name || ""} onChange={this.rename} />
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
        const title = this.tracking ? texts.workSession.actions.done : texts.workSession.actions.start;
        const action = this.tracking ? this.props.onSave : this.props.onStart;
        return (
            <button className="btn btn-primary" onClick={action}>
                {title}
            </button>
        );
    }
}