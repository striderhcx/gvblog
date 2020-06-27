import axios from 'axios'
import Vue from "../main"

// axios.defaults.baseURL = 'https://api.example.com';
// 这里改成vuex存储token
// axios.defaults.headers.common['Authorization'] = (window.localStorage.getItem('ihxn-blog-token') ||  "");

const ExternalIpPort = "http://120.78.76.74"
const TestEnvIpPort = "http://192.168.238.133:9999"

var host = ""
var prefix = ""

function get_config(env) {
    let config = {}
    if (env == "product") {
        config["host"] = ExternalIpPort
        config["prefix"] = "/api/"
    } else {
        config["host"] = TestEnvIpPort
        config["prefix"] = "/"
    }
    return config
}

({host, prefix} =  get_config("test"))


function request(config) {
    // 1. 创建axios的实例
    const instance = axios.create({
        baseURL: host+prefix,
        headers: { 
            'Content-Type': 'application/json',
            'Authorization': Vue.$store.state.auth_token,
        },
        timeout: 5000,
        withCredentials: true,
    })
    // 2. axios的拦截器
    // 2.1.请求拦截器 
    instance.interceptors.request.use(
        config => {
            // 1. 比如config中的一些　请求信息不符合服务器的需求格式

            // 2. 请求较慢时，可以在这加段动画

            // 3. 某些网络请求，比如登录ｔｏｋｅｎ,必须携带一些特殊信息
            
            //　注意处理完之后要返回 
            return config
        },
        err => {
            console.log(err)
        },
        
    )

    instance.interceptors.response.use(
        res => {
            // 未登录跳转
            let code  = res.data.code
            if (code === 1001) { // 未登录，需要跳转
              Vue.$router.push("/bloglogin")
            }

            // 注意处理完之后要返回
            return res.data
        },
        err => {
            console.log(err)
        },
    )

    return instance(config)
}

export default request
