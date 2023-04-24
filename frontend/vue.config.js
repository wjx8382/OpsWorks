module.exports = {
    publicPath: '/', // 设置根路径
    outputDir: 'dist', // 设置输出目录
    devServer: {
      port: 3000, // 设置开发环境下的端口号
      proxy: {
        '/api': {
          target: 'http://localhost:8080', // 设置后端接口地址
          changeOrigin: true, // 是否改变请求头中的Host值
          pathRewrite: {
            '^/api': '' // 替换掉原来的/api
          }
        }
      }
    }
  }