import axios from 'axios'
import VueAxios from 'vue-axios'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'

createApp(App)
.use(VueAxios, axios)
.use(router)
.mount('#app')

