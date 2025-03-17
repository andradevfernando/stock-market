import httpClient from './httpClient';
import type { StockData, StockQueryParams } from '../types/stock.types';
import type { AnalysisData } from '../types/stockAnalysis.types';

export const fetchStockData = async (params: StockQueryParams): Promise<StockData[]> => {
  const response = await httpClient.get('/stocks', { params });
  return response.status, response.data;
};
export const fetchBestStockData = async (): Promise<StockData[]> => {
  const response = await httpClient.get('/best-investments');
  return response.status, response.data;
};
export const fetchAnalysis = async (ticker: string): Promise<AnalysisData> => {
console.log(ticker);

  const response = await httpClient.get('/analysis/' + ticker);
  console.log(response);
  
  return response.status, response.data
}