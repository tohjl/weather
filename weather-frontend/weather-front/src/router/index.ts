import { createRouter, createWebHistory } from 'vue-router'
import WeatherView from '../views/WeatherView.vue';  // Import your custom view

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: WeatherView, // Use your custom WeatherView instead
    },
  ],
})

export default router
