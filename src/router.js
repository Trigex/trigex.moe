import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'root',
      redirect: '/home'
    },
    {
      path: '/home',
      name: 'Home',
      component: () => import('./views/Home.vue')
    },
    {
      path: '/about',
      name: 'About',
      component: () => import('./views/About.vue')
    },
    {
      path: '/projects',
      name: 'Projects',
      component: () => import('./views/Projects.vue')
    },
    {
      path: '/links',
      name: 'Links',
      component: () => import('./views/Links.vue')
    },
    {
      path: '/blog',
      name: 'Blog',
      component: () => import('./views/Blog.vue')
    }
  ]
})
