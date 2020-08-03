const path = require("path")

module.exports = {
    entry: {
        bundle: "./src/index.js"
    },

    output: {
        filename: "[name].js",
        path: path.resolve(__dirname, "public")
    },

    mode: "development",
    devtool: "#eval-source-map",

    //mode: "production",
    //devtool: "source-map",

    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: [
                    /node_modules/
                ],
                use: [
                    { loader: "babel-loader" }
                ]
            },
            {
                test: /\.css$/,
                use: [
                    'style-loader',
                    'css-loader'
                ]
            }
        ]
    }
}
