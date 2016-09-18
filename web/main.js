import 'font-awesome/css/font-awesome.css'
import 'source-sans-pro/source-sans-pro.css'

import Vue from 'vue'
import Router from 'vue-router'

import App from './components/App.vue'

Vue.use(Router)

const router = new Router()

router.start(App, '#app')
