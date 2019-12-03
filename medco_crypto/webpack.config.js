var webpack = require('webpack');
const path = require('path');

module.exports = {
    entry: {
        'crypto': './src/index.ts',
        'crypto.min': './src/index.ts'
    },
    module: {
        rules: [{
            test: /\.tsx?$/,
            loader: 'awesome-typescript-loader',
            exclude: /node_modules/,
            query: {
                declaration: false,
            }
        }]
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js']
    },
    devtool: 'source-map',
    output: {
        path: path.resolve(__dirname, '_bundles'),
        filename: '[name].js',
        libraryTarget: 'umd',
        library: 'Crypto',
        umdNamedDefine: true
    },
};