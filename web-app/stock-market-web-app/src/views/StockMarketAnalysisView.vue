<template>
    <div class="min-h-screen bg-gray-50 dark:bg-gray-900">
      <div class="container mx-auto p-6 max-w-7xl">
      
        <div class="mb-8 flex items-center justify-between">
          <button 
            @click="goBack"
            class="flex cursor-pointer items-center gap-2 px-4 py-2 bg-white dark:bg-gray-800 rounded-lg shadow-md hover:shadow-lg transition-shadow"
          >
            <svg class="w-5 h-5 text-gray-600 dark:text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
            </svg>
            <span class="text-gray-600 dark:text-gray-300">Go back</span>
          </button>
          
          <h1 class="text-4xl font-bold text-gray-900 dark:text-white text-center">
            📈 {{ analysisData?.ticker }} Analysis
          </h1>
          <div class="relative group flex items-center gap-2 px-4 py-2 bg-gradient-to-r from-purple-500 to-blue-500 rounded-full shadow-md">
            <svg class="w-5 h-5 text-white animate-spin-slow" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                    d="M13 10V3L4 14h7v7l9-11h-7z M19 14l-7 7m0 0l-7-7m7 7V3"/>
            </svg>
            <span class="text-white font-medium text-sm">Powered by AI</span>
            
        
            <div class="absolute inset-0 rounded-full bg-white/20 blur-md opacity-0 group-hover:opacity-30 transition-opacity"></div>
          </div> 
          <div class="w-24"></div> 
        </div>
  
    
        <div v-if="loading" class="text-center py-12">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
        </div>
  
        <div v-else-if="error" class="bg-red-100 dark:bg-red-900 text-red-700 dark:text-red-100 p-4 rounded-lg">
          Error: {{ error }}
        </div>
  
        <div v-else-if="analysisData" class="bg-white dark:bg-gray-800 shadow-xl rounded-xl p-8 space-y-8">
      
          <div :class="recommendationClasses" class="p-6 rounded-xl text-white">
            <div class="flex items-center gap-4">
              <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path v-if="analysisData.shouldInvest" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
              <div>
                <h2 class="text-2xl font-bold">Recommendation: {{ analysisData.recommendation }}</h2>
                <p class="text-lg mt-2">{{ analysisData.reason }}</p>
              </div>
            </div>
          </div>
  
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">

            <div class="bg-gray-50 dark:bg-gray-700 p-6 rounded-xl">
              <h3 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white border-b pb-2">📊 Key Metrics</h3>
              <div class="space-y-4">
                <MetricItem label="Investment Score" :value="analysisData.score" color="text-blue-500"/>
                <MetricItem label="Current Price" :value="analysisData.price" prefix="$"/>
                <MetricItem label="Market Cap" :value="analysisData.market_cap"/>
                <MetricItem label="Dividend Yield" :value="analysisData.dividend_yield" suffix="%"/>
              </div>
            </div>
  
            <div class="bg-gray-50 dark:bg-gray-700 p-6 rounded-xl">
              <h3 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white border-b pb-2">📈 Valuation Ratios</h3>
              <div class="space-y-4">
                <MetricItem label="P/E Ratio" :value="analysisData.pe_ratio"/>
                <MetricItem label="P/B Ratio" :value="analysisData.pb_ratio"/>
                <MetricItem label="Debt/Equity" :value="analysisData.debt_to_equity"/>
              </div>
            </div>
  
            <div class="bg-gray-50 dark:bg-gray-700 p-6 rounded-xl">
              <h3 class="text-xl font-semibold mb-4 text-gray-900 dark:text-white border-b pb-2">💼 Financial Health</h3>
              <div class="space-y-4">
                <div class="space-y-2">
                  <p class="text-gray-700 dark:text-gray-300">
                    <span class="font-medium">Analyst Consensus:</span>
                    <span class="ml-2 px-3 py-1 rounded-full bg-green-100 dark:bg-green-900 text-green-700 dark:text-green-100">
                      {{ analysisData.recommendation }}
                    </span>
                  </p>
                </div>
              </div>
            </div>
          </div>
  
          <div class="bg-blue-50 dark:bg-blue-900 p-6 rounded-xl">
            <h3 class="text-xl font-semibold mb-4 text-blue-800 dark:text-blue-100">💡 Expert Analysis</h3>
            <p class="text-gray-700 dark:text-blue-200 leading-relaxed">
              {{ analysisData.comment }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </template>

<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia';
import { useStockAnalysisStore } from '../stores/stockAnalysisStore';
import { onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import MetricItem from '../components/MetricItem.vue';


const route = useRoute();
const analysisStore = useStockAnalysisStore();
const { analysisData } = storeToRefs(analysisStore);
const { loading } = storeToRefs(analysisStore);
const { error } = storeToRefs(analysisStore);
const router = useRouter();
const goBack = () => {
  router.go(-1)
}

const recommendationClasses = computed(() => {
  return analysisData.value?.shouldInvest 
    ? 'bg-gradient-to-r from-green-500 to-emerald-600'
    : 'bg-gradient-to-r from-red-500 to-rose-600'
})

onMounted(() => {
    const ticker = route.params.ticker as string;
  if (ticker) {
    analysisStore.fetchStockAnalysis(ticker);
  }
});
</script>

<style>
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fade-in {
  animation: fadeIn 0.5s ease-out;
}

@keyframes pulse-slow {
  0%, 100% { transform: scale(1); opacity: 0.9; }
  50% { transform: scale(1.02); opacity: 1; }
}

.animate-pulse-slow {
  animation: pulse-slow 3s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.ai-particles::before {
  content: "";
  position: absolute;
  width: 150%;
  height: 150%;
  background: 
    radial-gradient(circle at 30% 30%, rgba(255,255,255,0.1) 0%, transparent 50%),
    radial-gradient(circle at 70% 70%, rgba(255,255,255,0.1) 0%, transparent 50%);
  animation: particle-drift 20s linear infinite;
}

@keyframes particle-drift {
  0% { transform: translate(-5%, -5%); }
  25% { transform: translate(5%, 5%); }
  50% { transform: translate(-5%, 5%); }
  75% { transform: translate(5%, -5%); }
  100% { transform: translate(-5%, -5%); }
}

</style>