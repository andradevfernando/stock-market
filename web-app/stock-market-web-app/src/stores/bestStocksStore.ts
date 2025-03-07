import { defineStore } from 'pinia';
import { ref, shallowRef } from 'vue';
import { fetchBestStockData } from '../services/stockService';
import type { StockData } from '../types/stock.types';

export const useBestStockStore = defineStore('bestStock', () => {
  const bestStocks = ref<StockData[]>([]);
  const loading = shallowRef(false);
  const error = shallowRef<string | null>(null);

  const fetchBestStocks = async (): Promise<StockData[]> => {
    loading.value = true;
    try {
        const response = await fetchBestStockData();
        bestStocks.value = response;
        return response;
      } catch (err) {
      error.value = 'Erro ao buscar dados das ações alternativas';
      console.error(err);
      return [];
    } finally {
      loading.value = false;
    }
  };

  return { bestStocks, loading, error, fetchBestStocks };
});