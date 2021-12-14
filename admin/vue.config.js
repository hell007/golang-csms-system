/*
 * @Descripttion: 
 * @Author: zenghua.wang
 * @Date: 2019-09-19 21:49:36
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2021-12-14 09:35:51
 */
// 参考文档 https://cli.vuejs.org/zh/config/#integrity
const moment = require('moment');
const path = require('path');

const buildDate = moment().format('YYYYMMDDhhmmss');

module.exports = {
  publicPath: './',
  outputDir: 'dist',
  //assetsDir: '/assets',
  indexPath: 'index.html',
  lintOnSave: process.env.NODE_ENV === 'development',
  productionSourceMap: false,
  runtimeCompiler: true,//热更新
  devServer: {
    port: 9528,
    open: true,
    hot: true,
    overlay: {
      warnings: false,
      errors: true
    },
    proxy: {
      [process.env.VUE_APP_BASE_API]: {
        target: "http://127.0.0.1:5000",
        changeOrigin: true,
        ws: true,
        pathRewrite: {
          ['^' + process.env.VUE_APP_BASE_API]: ''
        }
      }
    },
  },
  css: {
    extract: {
      filename: `css/[name].${buildDate}.css`,
      chunkFilename: `css/[name].${buildDate}.css`,
    },
  },
  configureWebpack: {
    output: {
      filename: `js/[name].${buildDate}.js`,
      chunkFilename: `js/[name].${buildDate}.js`,
    },
  },
  chainWebpack: config => {
    // 修复HMR
    //config.resolve.symlinks(true);
    const oneOfsMap = config.module.rule('scss').oneOfs.store
    oneOfsMap.forEach(item => {
      item
        .use('sass-resources-loader')
        .loader('sass-resources-loader')
        .options({
          resources: path.resolve(__dirname, './src/styles/_global.scss'),
        })
        .end()
    })
  },
}