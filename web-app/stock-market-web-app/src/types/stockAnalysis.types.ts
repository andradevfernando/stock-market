export interface AnalysisData {
    shouldInvest: boolean;
    reason: string;
    recommendation: string;
    comment: string;
    score: number;
    ticker: string;
    pe_ratio: number;
    pb_ratio: number;
    dividend_yield: number;
    price: number;
    debt_to_equity: number;
    market_cap: number;
  }