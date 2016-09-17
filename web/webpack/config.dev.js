var baseConfig = require('./config.base');

module.exports = Object.assign({}, baseConfig, {
  output: {
    path: __dirname,
    filename: 'build.js'
  }
});