export interface StockData {
    ticker: string;
    target_from: string;
    target_to: string;
    company: string;
    action: string;
    brokerage: string;
    rating_from: string;
    rating_to: string;
    time: string;
  }
  
  export interface StockQueryParams {
    page: number;
    items: number;
    startDate: string;
    endDate: string;
    company: string;
  }