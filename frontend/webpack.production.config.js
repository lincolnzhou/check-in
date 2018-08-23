const webpack = require("webpack");
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin'); 
const ExtractTextPlugin = require('extract-text-webpack-plugin');

let pathsToClean = [
  'dist',
  'build',
	"static",
	"views",
]

module.exports = {
	devtool: 'eval-source-map',
	entry: __dirname + "/app/main.js",
	output: {
		path: __dirname + "/static",
		filename: "bundle-[hash:8].js",
	},
	devServer: {
		contentBase: "./views", // 本地页面目录
		historyApiFallback: true, // 不跳转
		inline: true, // 实时刷新
	},
	module: {
		rules:[
			{
				test: /(\.jsx|\.js)$/,
				use: {
					loader: "babel-loader",
				},
				exclude: /node_modules/
			},
			{
				test: /\.css$/,
				use: ExtractTextPlugin.extract({
					fallback: 'style-loader',
					use: [
						{
							loader: 'css-loader',
							options: {
								sourceMap: true,
								importLoaders: 1,
								modules: true,
								camelCase: true,
								localIdentName: '[name]_[local]_[hash:base64:5]',
								minimize: true
							},
						},
						{
							loader: 'postcss-loader'
						}
					]
				}),	
				exclude: /node_modules/,
			}
		]
	},
	plugins: [
		new webpack.BannerPlugin("版权所有，翻版必究"),
		new HtmlWebpackPlugin({
			filename: __dirname + "/index.html",
			template: __dirname + "/app/index.tmpl.html"
		}),
		new webpack.optimize.OccurrenceOrderPlugin(),
		new ExtractTextPlugin({
			filename: '[hash:8].css'
		}),
		new CleanWebpackPlugin(pathsToClean),
	]
};
