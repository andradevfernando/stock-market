import { createWebHistory, createRouter } from 'vue-router'

import StockMarketView from '../views/StockMarketView.vue'
import StockMarketAnalysis from '../views/StockMarketAnalysis.vue';

const routes = [
  { path: '/', component: StockMarketView },
  { path: '/analysis/:ticker', name: 'analysis', component: StockMarketAnalysis }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;