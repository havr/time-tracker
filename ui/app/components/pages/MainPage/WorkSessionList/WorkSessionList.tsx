import * as React from "react";

import moment from "moment";

import WorkSession from "models/WorkSession";
import { formatInterval } from "util/time";
import texts from "constants/texts";

interface Props {
    list: WorkSession[]
}

export default class WorkSessionList extends React.Component<Props> {
    render() {
        const { list } = this.props;
        const sessions = list.map(this.renderWorkSession);

        return (
            <table className="table">
              <thead>
                <tr>
                  <th scope="col">{texts.workSession.fields.name}</th>
                  <th scope="col">{texts.workSession.fields.duration}</th>
                  <th scope="col">{texts.workSession.fields.started}</th>
                </tr>
              </thead>
              <tbody>
                {sessions}
              </tbody>
            </table>
        )
    }

    renderWorkSession(session: WorkSession) {
        const name = session.name || <i> {texts.workSession.unnamed} </i>;
        const startTime = moment(session.startTime).format("DD/MM/YYYY HH:mm:ss");
        const duration = formatInterval(session.duration);

        return (
            <tr key={session.id}>
                <td> {name} </td>
                <td> {duration} </td>
                <td> {startTime} </td>
            </tr>
        )
    }
}