axios.interceptors.response.use(
  response => response,
  error => {
    if (error.response) {
      const { status, data } = error.response
      let message = '请求失败'
      if (data.message) message = data.message
      
      switch (status) {
        case 400: ElMessage.error(`参数错误: ${message}`) break
        case 401: ElMessage.error('请重新登录') break
        case 500: ElMessage.error('服务器内部错误') break
      }
    }
    return Promise.reject(error)
  }
)