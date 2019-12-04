import moment from "moment";

export function formatInterval(duration) {
    const value = moment.duration(duration);
    return [value.hours(), value.minutes(), value.seconds()].map(padZeroes).join(":");
}

export function intervalBetween(end: Date, start: Date) {
    return (+ end) - (+ start)
}

function padZeroes(value) {
    let stringValue = "" + value;
    if (stringValue.length == 1) {
        stringValue = "0" + stringValue;
    }

    return stringValue;
}
