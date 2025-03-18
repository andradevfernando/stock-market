<template>
  <StockFilter
    class="sticky top-0 z-50 text-slate-900 dark:text-white transition-all duration-300"
    @filter="handleFilter" 
  />

  <div class="background bg-gray-200 bg-opacity-50 dark:bg-slate-900 dark:bg-opacity-50 backdrop-blur-sm text-slate-900 dark:text-white min-h-screen pt-4 transition-all">
    <div class="text-center text-4xl font-bold mb-4 text-gray-900 dark:text-white">
    Stay Ahead with the Best Stock Picks!
  </div>
    <BestStockList/>
    <div class="text-center text-4xl font-bold  pt-4 text-gray-900 dark:text-white">
    All stocks
  </div>
    <StockList/>
  </div>
</template>

<script setup lang="ts">
import StockFilter from '../components/StockFilter.vue';
import StockList from '../components/StockList.vue';
import BestStockList from '../components/BestStockList.vue';
import { useStockStore } from '../stores/stockStore';
import {useBestStockStore} from '../stores/bestStocksStore';
import { ref, onMounted, onUnmounted } from 'vue';

const stockStore = useStockStore();
const bestStockStore = useBestStockStore();
// const stocks = ref<any[]>([]);
const currentPage = ref(1);
const isLoadingMore = ref(false);
const scrolled = ref(false);
const scrollY = ref(0);
const now = new Date();
const startDate = ref(new Date(now.getTime() - (48 * 60 * 60 * 1000)));
const endDate = ref(now);
const company = ref('');

const handleScroll = async () => {
  scrollY.value = window.scrollY || window.pageYOffset;
  scrolled.value = scrollY.value > 50;

  if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight - 100 && !isLoadingMore.value) {
    currentPage.value += 1;
    await stockStore.fetchStocks({
      page : currentPage.value,
      items: 12,
      company: company.value,
      startDate: startDate.value.toISOString(),
      endDate: endDate.value.toISOString(),
    });
  }
};
onMounted(async () => {
  window.addEventListener('scroll', handleScroll);
  await bestStockStore.fetchBestStocks();
  await stockStore.fetchStocks({
    page: 1,
    items: 12,
    company: '',
    startDate: startDate.value.toISOString(),
    endDate: endDate.value.toISOString(),
  }).then(() => {
  });
});

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll);
});

const handleFilter = async (filters: { company: string; startDate: string; endDate: string }) => {
  currentPage.value = 1;
  company.value = filters.company;
  startDate.value = new Date(filters.startDate);
  endDate.value = new Date(filters.endDate);
  await stockStore.fetchStocks(
    {
    page: 1,
    items: 12,
    company: filters.company,
    startDate: filters.startDate,
    endDate: filters.endDate,
  }
  );
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  });
};
</script>