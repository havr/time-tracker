import * as React from "react";

import texts from "constants/texts";

export interface Element {
    name: string;
    title: string;
}

interface Props {
    current: string;
    available: Element[]
    onChange: (string) => void;
}

export default class WorkSessionListRange extends React.Component<Props> {
    render() {
        const { current, available, onChange } = this.props;

        const elements = available.map(this.renderElement);
        const change = (e) => onChange(e.target.value);
        return (
              <div className="form-group row">
                  <label className={"col-form-label"}> {texts.workSession.listRange.title} </label>
                  <div className={"col-sm-2"}>
                      <select className="form-control" value={current} onChange={change}>
                          {elements}
                      </select>
                  </div>
            </div>
        )
    }

    renderElement = (element: Element) => {
        return (
            <option key={element.name}
                    value={element.name} >
                {element.title}
            </option>
        );
    }
}