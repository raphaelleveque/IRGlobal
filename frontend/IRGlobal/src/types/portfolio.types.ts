export interface Position {
  symbol: string;
  quantity: number;
  avgPrice: number;
  currentPrice: number;
  pnl: number;
  pnlPercentage: number;
  marketValue: number;
  totalCost: number;
}

export interface PortfolioSummary {
  totalValue: number;
  totalPnL: number;
  pnlPercentage: number;
  dailyPnL: number;
  dailyPnLPercentage: number;
  totalPositions: number;
  totalInvested: number;
}

export interface PortfolioData {
  summary: PortfolioSummary;
  positions: Position[];
  lastUpdated: Date;
}
