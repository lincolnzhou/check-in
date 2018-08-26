import axios from "axios"
import port_code from "../common/port_uri"

axios.defaults.baseURL = '/'
axios.defaults.withCredentials = true
axios.defaults.timeout = 100000

// 请求前拦截器
axios.interceptors.request.use(config => {
	config.headers.Authorization = "123456";

	return config;
})

// 请求后拦截器
axios.interceptors.response.use(response => {
	let resData = response.data
	let dataCode = response.code 
	let dataMsg = response.msg

	if (dataCode == port_code.success){
		return Promise.resolve(response)
	}
})

export default axios
