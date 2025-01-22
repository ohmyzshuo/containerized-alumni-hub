import './assets/main.css'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import {createApp} from 'vue'
import App from './App.vue'
import router from './router'

import {createPinia} from "pinia";

const app = createApp(App)
const pinia = createPinia()
app.use(pinia)
app.use(router)
app.use(ElementPlus)
import 'virtual:windi.css'

app.mount('#app')
