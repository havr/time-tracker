exports.files = {
    javascripts: {
        joinTo: 'index.js',
    },
    stylesheets: {
        joinTo: 'index.css'
    }
};

exports.modules = {
    autoRequire: {
        'index.js': ['index.tsx']
    }
};

exports.paths = {
    public: 'build'
};

exports.plugins = {
    sass: {
        options: {
            includePaths: ['node_modules/bootstrap/scss'],
        }
    },
    copycat: {}
};

