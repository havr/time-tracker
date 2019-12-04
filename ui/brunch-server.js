var express = require('express');
var httpProxy = require('http-proxy');
var app = express();

// TODO: use environment variable?
const API = 'http://localhost:8080';

module.exports = (config, callback) => {
    var proxy = httpProxy.createProxyServer({
        target: API
    });
    app.all('/api/*', proxy.web.bind(proxy));
    app.use(express.static(__dirname + '/' + config.path));
    app.listen(config.port, function () {
        console.log('serving on port', config.port);
        callback();
    });

    return app;
};

