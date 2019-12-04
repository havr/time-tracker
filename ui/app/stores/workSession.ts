import {WorkSessionActionNames, WorkSessionActions} from "actions/WorkSession";
import WorkSession from "models/WorkSession";

interface State {
    current: WorkSession
    list: WorkSession[]
    range: string
}

const initialState = {
    current: {
        name: "",
        duration: 0,
    },
    range: "day",
    list: [],
};

export default function workerSessionStore(state: State = initialState, action: WorkSessionActions) {
    switch (action.type) {
    case WorkSessionActionNames.LIST:
        return {...state, list: action.payload.sessions, range: action.payload.range};
    case WorkSessionActionNames.START:
        return {...state, current: {
                ...state.current, startTime: action.payload.startTime,
            }
        };
    case WorkSessionActionNames.SAVE:
        return {...state, current: {}, list: [action.payload.session, ...state.list]};
    case WorkSessionActionNames.RENAME:
        return {...state, current: {...state.current, name: action.payload.name}};
    default:
        return state;
    }
}

