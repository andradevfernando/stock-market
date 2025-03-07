<template>
    <div class="marquee-container overflow-hidden relative">
      <div class="marquee flex gap-6 w-max">
        <div 
          v-for="stock in bestStocks" 
          :key="stock.ticker + '-1'" 
          class="marquee-item bg-white dark:bg-slate-800 rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow min-w-[300px]"
        >
          <div class="space-y-2">
            <p class="text-gray-900 dark:text-white">Ticker: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.ticker }}</span>
            </p>
            <p class="text-gray-900 dark:text-white">Brokerage: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.brokerage }}</span>
            </p>
            <p class="text-gray-900 dark:text-white">Company: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.company }}</span>
            </p>
            <p class="text-gray-900 dark:text-white">Target: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.target_from }} - {{ stock.target_to }}</span>
            </p>
            <div class="mt-5 text-gray-900 dark:text-white">
            <AnalysisButton :currentTicker="stock.ticker" class="text-blue-600 dark:text-blue-400"/>
          </div>
          </div>
        </div>
        
        <div 
          v-for="stock in bestStocks" 
          :key="stock.ticker + '-2'" 
          class="marquee-item bg-white dark:bg-slate-800 rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow min-w-[300px]"
        >
          <div class="space-y-2">
            <p class="text-gray-900 dark:text-white">Ticker: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.ticker }}</span>
            </p>
            <p class="text-gray-900 dark:text-white">Brokerage: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.brokerage }}</span>
            </p>
            <p class="text-gray-900 dark:text-white">Company: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.company }}</span>
            </p>
            <p class="text-gray-900 dark:text-white">Target: 
              <span class="text-blue-600 dark:text-blue-400">{{ stock.target_from }} - {{ stock.target_to }}</span>
            </p>
            <div class="mt-5 text-gray-900 dark:text-white">
            <AnalysisButton :currentTicker="stock.ticker" class="text-blue-600 dark:text-blue-400"/>
          </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { storeToRefs } from 'pinia';
  import { useBestStockStore } from '../stores/bestStocksStore';
  import  AnalysisButton  from '../components/AnalysisButton.vue';
  
  const bestStockStore = useBestStockStore();
  const { bestStocks } = storeToRefs(bestStockStore);
  </script>
  
  <style scoped>
  .marquee-container {
    position: relative;
    width: 100%;
    overflow: hidden;
  }
  
  .marquee {
    animation: marquee 20s linear infinite;
    display: flex;
    flex-wrap: nowrap;
    will-change: transform;
  }
  
  .marquee:hover {
    animation-play-state: paused;
  }
  
  @keyframes marquee {
    0% {
      transform: translateX(0%);
    }
    100% {
      transform: translateX(-50%);
    }
  }
  
  .marquee-item {
    flex: 0 0 auto;
    margin-right: 1.5rem; /* Espaçamento entre os cards */
  }
  
  .dark .marquee-item {
    background-color: #1e293b; /* Cor para dark mode */
  }
  </style>