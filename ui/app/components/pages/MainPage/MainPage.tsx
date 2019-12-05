import * as React from "react";
import { connect } from "react-redux";
import { ThunkDispatch } from "redux-thunk";

import CurrentWorkSession from "components/pages/MainPage/CurrentWorkSession/CurrentWorkSession";
import WorkSessionListRange from "components/pages/MainPage/WorkSessionListRange/WorkSessionListRange";
import WorkSessionList from "components/pages/MainPage/WorkSessionList/WorkSessionList";
import {
    listWorkSessions,
    renameWorkSession,
    saveWorkSession,
    startWorkSession
} from "actions/WorkSession";
import WorkSession from "models/WorkSession";
import {intervalBetween} from "util/time";
import texts from "constants/texts";

interface Props {
    currentWorkSession: WorkSession;
    workSessionList: WorkSession[];
    range: string;

    onStart: (name: string) => void;
    onList: (range: string) => void;
    onSave: (session: WorkSession) => void;
    onRename: (name: string) => void;
}

const listRanges = [
    {name: "day", title: texts.workSession.listRange.types.day},
    {name: "week", title: texts.workSession.listRange.types.week},
    {name: "month", title: texts.workSession.listRange.types.month},
];

class MainPage extends React.Component<Props> {
    componentDidMount() {
        const { onList, range } = this.props;

        onList(range);
    }

    render() {
        const { range, onList, onRename, workSessionList, currentWorkSession, } = this.props;

        return (
            <div className="container">
                <div className="row">
                    <div className="col-md-12">
                        <div className="mt-2">
                          <CurrentWorkSession
                              value={currentWorkSession}
                              onStart={this.start}
                              onSave={this.save}
                              onRename={onRename}
                          />
                        </div>
                        <div className="mt-3">
                            <WorkSessionListRange current={range} available={listRanges} onChange={onList}/>
                        </div>
                        <div className="mt-1">
                            <WorkSessionList list={workSessionList} />
                        </div>
                    </div>
                </div>
            </div>
        )
    }

    start = () => {
        const { onStart, currentWorkSession } = this.props;

        onStart(currentWorkSession.name);
    };

    save = () => {
        const { onSave, currentWorkSession } = this.props;

        const duration = intervalBetween(new Date (), currentWorkSession.startTime);
        onSave({...currentWorkSession, duration: duration});
    };
}

function inject({ workSession }) {
    return {
        currentWorkSession: workSession.current,
        workSessionList: workSession.list,
        range: workSession.range,
    }
}

function actions(dispatch: ThunkDispatch<{}, {}, any>) {
    return {
        onStart: () => dispatch(startWorkSession()),
        onSave: (session) => dispatch(saveWorkSession(session)),
        onRename: (newName) => dispatch(renameWorkSession(newName)),
        onList: (range) => dispatch(listWorkSessions(range))
    }
}

export default connect(inject, actions) (MainPage);
