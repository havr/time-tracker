import WorkSession from "models/WorkSession";

export class ApiClient {
    _url = "";

    constructor(url: string) {
        this._url = url;
    }

    listWorkSessions(range) {
        return this.query("GET", "/api/v1/sessions?range=" + range);
    }

    saveWorkSession(session: WorkSession) {
        return this.query("POST", "/api/v1/sessions", session);
    }

    query(method, url, body = null) {
        const params: RequestInit = {
            method,
        };

        if (body) {
            params.body = JSON.stringify(body);
        }

        return fetch(this._url + url, params)
            .then(response => {
                return response.text().then((value) => {
                    if (response.status < 200 || response.status >= 300) {
                        throw `Http Error ${response.status}: ${value}`;
                    } else {
                        if (value) {
                            return JSON.parse(value);
                        }
                        return null;
                    }
                })
            });
    }
}

export default new ApiClient("http://" + window.location.host);