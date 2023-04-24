import Vue from 'vue'
import List from './components/list.vue'

Vue.config.productionTip = false

new Vue({
  render: h => h(List),
}).$mount('#app')