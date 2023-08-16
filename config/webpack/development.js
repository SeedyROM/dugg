const { merge } = require("webpack-merge");
const BrowserSyncPlugin = require("browser-sync-webpack-plugin");

const common = require("./common.js");

const DEV_WEBSITE_URL = process.env.DEV_WEBSITE_URL || "http://localhost:6661";

module.exports = merge(common, {
  mode: "development",
  devtool: "inline-source-map",
  watch: true,
  plugins: [
    new BrowserSyncPlugin({
      files: ["**/*.html", "**/*.go", "**/*.css", "**/*.js"],
      proxy: DEV_WEBSITE_URL,
      notify: false,
      reloadDelay: 350,
      reloadDebounce: 1000,
    }),
  ],
});
