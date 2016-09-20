import 'font-awesome/css/font-awesome.css'
import 'source-sans-pro/source-sans-pro.css'

import Vue from 'vue'
import Router from 'vue-router'

import App from 'components/App.vue'
import Main from 'components/Main.vue'
import Editor from 'components/Editor.vue'

Vue.use(Router)

const router = new Router({
  history: true
})

router.map({
  '/': {
    component: Main
  },
  '/edit': {
    component: Editor
  }
})

router.start(App, '#app')
