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
	entry: __dirname + "/app/main.js",
	output: {
		path: path.resolve(__dirname, "../backend/static"),
		filename: "js/[chunkhash:8].js"
	},
	devServer: {
		contentBase: ".", // 本地页面目录
		historyApiFallback: true, // 不跳转
		inline: true, // 实时刷新
		open: false,
		overlay: true,
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
			filename: path.resolve(__dirname, "../backend/static/index.html"),
			template: __dirname + "/app/index.tmpl.html"
		}),
		new webpack.optimize.OccurrenceOrderPlugin(),
		new ExtractTextPlugin({
			filename: 'css/[hash:8].css'
		}),
		new CleanWebpackPlugin(pathsToClean, cleanOptions),
	]
};
