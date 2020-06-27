import Vue from 'vue'
import Vuex from 'vuex'

import mutations from './mutations.js'
import getters from './getters.js'
import actions from './actions.js'
import modules from './modules.js'

// 1. 安装插件
Vue.use(Vuex)

const state = {
    auth_token: "", 
}

// 2. 创建对象
const store = new Vuex.Store({
    state,
    mutations,
    actions,
    getters,
    modules,
})

// 3. 导出
export default store