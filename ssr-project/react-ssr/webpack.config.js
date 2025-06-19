const path = require("node:path");

module.exports = {
  entry: { client: "./src/app/client-jsx-entry.tsx" },
  mode: "development",
  output: {
    publicPath: "/public/",
    filename: "[name].bundle.js",
    chunkFilename: "[name].bundle.js",
    path: path.resolve(__dirname, "dist/public"),
  },
  resolve: {
    extensions: [".ts", ".tsx", "..."],
  },
  devtool: "eval-source-map",
  module: {
    rules: [
      {
        test: /\.(ts|tsx|js)$/,
        loader: "ts-loader",
        exclude: /node_modules/,
      },
    ],
  },
};
