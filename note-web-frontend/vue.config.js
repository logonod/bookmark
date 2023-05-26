module.exports = {
  devServer: {
    proxy: {
      '/api/user': {
        target: 'http://localhost:8000/',
        changeOrigin: true
      },
      '/api/link': {
      	target: 'http://localhost:8002/',
        changeOrigin: true
      }
    }
  }
}
