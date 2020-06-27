import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/index.js'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

Vue.config.productionTip = false

Vue.use(hljs)
Vue.directive('highlight', function (el) {
  let blocks = el.querySelectorAll('pre');
  blocks.forEach((block) => {
    hljs.highlightBlock(block)
  })
})

Vue.use(mavonEditor)

export default new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app')