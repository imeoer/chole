var webpack = require('webpack');
var baseConfig = require('./config.base');
var ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = Object.assign({}, baseConfig, {
  output: {
    path: __dirname,
    filename: '../dist/build.js'
  },
  vue: {
    loaders: {
      css: ExtractTextPlugin.extract("css")
    }
  },
  plugins: [
    new ExtractTextPlugin("../dist/style.css"),
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: '"production"'
      }
    }),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      }
    }),
    new webpack.optimize.OccurenceOrderPlugin()
  ]
});