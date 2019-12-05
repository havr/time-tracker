import * as React from "react";

import {formatInterval, intervalBetween} from "util/time";

interface Props {
    start?: Date;
}

export default class Timer extends React.Component<Props> {
    timer = null;

    constructor(props) {
        super(props);

        if (props.start) {
            this.startTimer();
        }
    }

    componentDidUpdate(oldProps) {
        const { start } = this.props;

        if (oldProps.start != this.props.start) {
            if (start) {
                this.startTimer();
            } else {
                this.stopTimer();
            }
        }
    }

    startTimer = () => {
        this.timer = setInterval(() => {
            this.forceUpdate();
        }, 1000)
    };

    stopTimer = () => {
        clearTimeout(this.timer);
    };

    render() {
        const { start } = this.props;

        const interval = start ? intervalBetween(new Date(), start) : 0;
        const formattedInterval = formatInterval(interval);
        return (
            <div> { formattedInterval } </div>
        )
    }
}

