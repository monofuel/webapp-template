module.exports = {
	entry: './frontend/index.tsx',
	output: {
		filename: 'index.js',
		path: __dirname + '/public/js/'
	},

	devtool: 'source-map',

	resolve: {
		extensions: ['.ts', '.tsx', '.js', '.json']
	},

	module: {
		rules: [
			{ test: /\.tsx?$/, loader: "awesome-typescript-loader" },
			{ enforce: "pre", test: /\.js$/, loader: "source-map-loader" },
			{
				test: /\.css$/,
				use: ['style-loader', 'css-loader']
			}
		]
	},
};