// import {UpdateStudentInfo} from './mutation_defind.js'
// import Vue from 'vue'

export default {
    // 原则：一般mutation只做同步操作，确保devtools这个工具能正常跟踪调试
　　// 异步操作需要放到actions

    setAuthToken(state, pay_load) {
        state.auth_token = pay_load.auth_token
        // 暂时非常野鸡地实现登录之后设置token到sessionStorage
        sessionStorage.setItem("store",JSON.stringify(state))
    },
}