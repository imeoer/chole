var path = require('path');
var webpack = require('webpack');
var baseConfig = require('./config.base');
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var HtmlWebpackPlugin = require('html-webpack-plugin');
var CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = Object.assign({}, baseConfig, {
  output: {
    path: 'dist',
    filename: 'build.js'
  },
  plugins: [
    new ExtractTextPlugin("style.css"),
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
    new webpack.optimize.OccurenceOrderPlugin(),
    new HtmlWebpackPlugin({
      template: 'index.html',
      filename: 'index.html',
      inject: true,
      minify: {
        removeComments: true,
        collapseWhitespace: true,
        removeAttributeQuotes: true
      },
      chunksSortMode: 'dependency'
    }),
    new CopyWebpackPlugin([
      {
        from: 'node_modules/monaco-editor/min/vs',
        to: 'vs',
      }
    ])
  ],
  vue: {
    loaders: {
      css: ExtractTextPlugin.extract("css!postcss")
    },
    postcss: [require('autoprefixer'), require('precss')],
    autoprefixer: {
      browsers: ['last 2 versions']
    }
  },
});