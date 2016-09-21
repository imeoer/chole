var webpack = require('webpack');
var baseConfig = require('./config.base');
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = Object.assign({}, baseConfig, {
  output: {
    path: __dirname,
    filename: 'build.js'
  },
  plugins: [
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoErrorsPlugin(),
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: 'index.html',
      inject: true
    })
  ],
  vue: {
    postcss: [require('autoprefixer'), require('precss')],
    autoprefixer: {
      browsers: ['last 2 versions']
    }
  },
  devtool: '#eval-source-map'
});
