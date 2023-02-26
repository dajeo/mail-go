import { createRouter, createMemoryHistory } from 'vue-router'
import TrackView from '../views/TrackView.vue'

const router = createRouter({
  history: createMemoryHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'track',
      component: TrackView
    },
    {
      path: '/createParcel',
      name: 'createParcel',
      component: () => import('../views/AddView.vue')
    },
    {
      path: '/items/:id',
      name: 'items',
      component: () => import('../views/ItemsView.vue')
    },
    {
      path: '/auth',
      name: 'auth',
      component: () => import('../views/AuthView.vue')
    },
    {
      path: '/findItems',
      name: 'findItems',
      component: () => import('../views/FindItemsView.vue')
    },
    {
      path: '/findStatus',
      name: 'findStatus',
      component: () => import('../views/FindParcelView.vue')
    },
    {
      path: '/status/:id',
      name: 'status',
      component: () => import('../views/StatusView.vue')
    }
  ]
})

export default router
