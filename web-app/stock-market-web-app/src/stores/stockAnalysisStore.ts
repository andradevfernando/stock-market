import { defineStore } from 'pinia';
import { ref, shallowRef } from 'vue';
import { fetchAnalysis } from '../services/stockService';
import type { AnalysisData } from '../types/stockAnalysis.types';

export const useStockAnalysisStore = defineStore('stockAnalysisStore', () => {
  const analysisData = ref<AnalysisData>();
  const loading = shallowRef(false);
  const error = shallowRef<string | null>(null);

  const fetchStockAnalysis = async (ticker: string): Promise<AnalysisData> => {
    loading.value = true;
    try {
        const response = await fetchAnalysis(ticker);
        analysisData.value = response;
        return response;
      } catch (err) {
      error.value = 'Error analysing stock';
      console.error(err);
      return analysisData.value ?? {} as AnalysisData;
    } finally {
      loading.value = false;
    }
  };

  return { analysisData, loading, error, fetchStockAnalysis };
});