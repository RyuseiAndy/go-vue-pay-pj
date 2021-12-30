import Vue from 'vue'
import App from './App.vue'
import router from './router'
import PayjpCheckout from 'vue-payjp-checkout'
import vuetify from './plugins/vuetify'

Vue.config.productionTip = false
Vue.use(PayjpCheckout)

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
