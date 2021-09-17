import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import axios from "axios"
import vuetify from '@/plugins/vuetify'
import qs from 'qs'

import VueCodeMirror from 'vue-codemirror'
import 'codemirror/lib/codemirror.css'
import 'xterm/css/xterm.css'
import 'xterm-addon-fit'


Vue.use(VueCodeMirror)

Vue.prototype.$ajax = axios
Vue.prototype.$qs = qs
Vue.config.productionTip = false
Vue.use(ElementUI,{size:'small',zIndex:2000})


if (process.env.NODE_ENV==='development'){
  axios.defaults.baseURL = 'http://10.135.168.120:9000/'
  // axios.defaults.baseURL = 'http://192.168.70.137:9000/'
  // axios.defaults.baseURL = 'http://192.168.50.168:9000/'
}else{
  axios.defaults.baseURL='/'
}
console.log(process.env)

new Vue({
  vuetify,
  render: h => h(App),
}).$mount('#app')
