import { defineStore } from 'pinia';
import { ref, shallowRef } from 'vue';
import { fetchStockData } from '../services/stockService';
import type { StockData, StockQueryParams } from '../types/stock.types';

export const useStockStore = defineStore('stock', () => {
  const stocks = ref<StockData[]>([]);
  const loading = shallowRef(false);
  const error = shallowRef<string | null>(null);
  
  var lastPage = false;

  const fetchStocks = async (params: StockQueryParams) => {

    loading.value = true;
    try {
      if(!lastPage || params.page === 1)
      {
        lastPage = false;
        var response = await fetchStockData(params);
        if(response.length === 0 )
        {
          lastPage = true;
        }
        console.log(params.page);
        
        if(params.page === 1)
        {
          stocks.value = [];	
        }
        stocks.value.push(...(response));

      }
    } catch (err) {

      console.log(error);
      error.value = 'Erro ao buscar dados das ações';
      console.error(err);
      return [];
    } finally {
      loading.value = false;
    }
  };
  console.log("fetch", stocks)
  return { stocks, loading, error, fetchStocks };
});