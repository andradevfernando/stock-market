<template>
    <!-- <div v-if="loading">Loading...</div>
    <div v-else-if="error">{{ error }}</div> -->
    <div class="max-w-6xl mx-auto px-4 py-2">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    <div 
      v-for="stock in stocks" 
      :key="stock.ticker" 
      class="bg-white dark:bg-slate-800 rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow"
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
          <p class="text-gray-900 dark:text-white">Action: 
            <span class="text-blue-600 dark:text-blue-400">{{ stock.action }}</span>
          </p>
          <p class="text-gray-900 dark:text-white">Rating: 
            <span class="text-blue-600 dark:text-blue-400">{{ stock.rating_from }} → {{ stock.rating_to }}</span>
          </p>
          <p class="text-gray-900 dark:text-white">Target: 
            <span class="text-blue-600 dark:text-blue-400">{{ stock.target_from }} - {{ stock.target_to }}</span>
          </p>
          <p class="text-gray-900 dark:text-white">Date: 
            <span class="text-blue-600 dark:text-blue-400">{{ stock.time }}</span>
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
  import { useStockStore } from '../stores/stockStore';
  import  AnalysisButton  from '../components/AnalysisButton.vue';
  const stockStore = useStockStore();
  const { stocks } = storeToRefs(stockStore);
  </script>