import { createWebHistory, createRouter } from 'vue-router'

import StockMarketView from '../views/StockMarketView.vue'
import StockMarketAnalysisView from '../views/StockMarketAnalysisView.vue';

const routes = [
  { path: '/', component: StockMarketView },
  { path: '/analysis/:ticker', name: 'analysis', component: StockMarketAnalysisView }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router;