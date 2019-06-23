import axios from 'axios'
axios.defaults.baseURL = '/api'
axios.interceptors.request.use(config => config)
axios.interceptors.response.use(
  res => res.data,
  err => {
    // 错误处理
    if (err) {
      return {}
    }
  }
)

export default axios
