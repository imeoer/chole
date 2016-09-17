module.exports = {
  entry: './main.js',
  module: {
    preLoaders: [
      {
        test: /.vue$/,
        loader: 'eslint',
        exclude: /node_modules/
      }
    ],
    loaders: [
      {
        test: /\.vue$/,
        loader: 'vue'
      },
      {
        test: /\.js$/,
        loader: 'babel',
        exclude: /node_modules/
      },
      {
        test: /\.(png|jpg|gif)$/,
        loader: 'url',
        query: {
          limit: 10000,
          name: '[name].[ext]?[hash]'
        }
      }
    ]
  },
  babel: {
    presets: ["es2015"],
    plugins: ["transform-runtime"]
  },
  vue: {
    postcss: [require('autoprefixer'), require('precss')],
    autoprefixer: {
      browsers: ['last 2 versions']
    }
  }
}