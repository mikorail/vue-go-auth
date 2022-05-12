import Vue from 'vue'
import App from './App.vue'
import store from './common/store/index'
import { router } from './common/router/routes'
import vuetify from './common/lib/vuetify'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
