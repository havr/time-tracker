import Api from "api/Api";
import WorkSession from "models/WorkSession";

export enum WorkSessionActionNames {
    LIST = "WORK_SESSION/LIST",
    START = "WORK_SESSION/START",
    SAVE = "WORK_SESSION/SAVE",
    RENAME = "WORK_SESSION/RENAME",
}

export interface ListWorkSessionsAction {
    type: WorkSessionActionNames.LIST,
    payload: {
        sessions: WorkSession[],
        range: string,
    },
};

export function listWorkSessions(range) {
    return (dispatch) => {
        return Api.listWorkSessions(range).then((workSessions) => {
            dispatch({
                payload: {
                    sessions: workSessions,
                    range,
                },
                type: WorkSessionActionNames.LIST,
            })
        })
    }
}

export type StartWorkSessionAction = {
    type: WorkSessionActionNames.START
    payload: {
        startTime: Date
    }
};

export function startWorkSession() {
    return {
        type: WorkSessionActionNames.START,
        payload: {
            startTime: new Date(),
        },
    };
}

export type SaveWorkSessionAction = {
    type: WorkSessionActionNames.SAVE,
    payload: {
        session: WorkSession
    }
};

export function saveWorkSession(session: WorkSession) {
    return dispatch => {
        return Api.saveWorkSession(session).then(() => {
            dispatch({
                type: WorkSessionActionNames.SAVE,
                payload: {
                    session: session,
                }
            })
        })
    };
}

export type RenameWorkSessionAction = {
    type: WorkSessionActionNames.RENAME,
    payload: {
        name: string
    }
};

export function renameWorkSession(name) {
    return {
        type: WorkSessionActionNames.RENAME,
        payload: {
            name
        }
    };
}

export type WorkSessionActions = ListWorkSessionsAction
    | StartWorkSessionAction
    | SaveWorkSessionAction
    | RenameWorkSessionAction;

