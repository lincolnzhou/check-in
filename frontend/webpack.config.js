const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin'); 
const ExtractTextPlugin = require('extract-text-webpack-plugin');

let pathsToClean = [
	"static",
]

let cleanOptions = {
	root: path.resolve(__dirname, "../backend")
}

module.exports = {
	devtool: 'eval-source-map',
	entry: {
		main: __dirname + "/app/main.js",
	},
	output: {
		path: path.resolve(__dirname, "../backend/static"),
		filename: "js/[name].js"
	},
	devServer: {
		contentBase: ".", // 本地页面目录
		historyApiFallback: true, // 不跳转
		inline: true, // 实时刷新
		open: false,
		overlay: true,
		proxy: {
			'api/*': {
					target: "http://localhost:1314",
					changeOrigin: true,
				 	secure: false,
				}
		},
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
				test: /\.less$/,
				use: ExtractTextPlugin.extract({
					fallback: 'style-loader',
					use: [{
							loader: "css-loader",
							options: {
								importLoaders: 1,
							},
						}, {
							loader: "less-loader"
						}],
				}),	
				exclude: /node_modules/,
			},
			{ test: /\.css$/, loader: "style-loader!css-loader"},
			{ test: /\.(png|jpg)$/, loader: "url-loader?limit=8192"},
		]
	},
	plugins: [
		new HtmlWebpackPlugin({
			filename: path.resolve(__dirname, "../backend/static/index.html"),
			template: __dirname + "/app/index.tmpl.html",
			favicon: './app/favicon.ico',
		}),
		new webpack.optimize.OccurrenceOrderPlugin(),
		new ExtractTextPlugin({
			filename: 'css/[name].css'
		}),
		new CleanWebpackPlugin(pathsToClean, cleanOptions),
	],
};
