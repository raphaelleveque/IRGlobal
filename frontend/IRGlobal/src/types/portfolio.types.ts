export type AssetType = "CRYPTO" | "ETF" | "STOCK";

export interface Position {
  id: string;
  asset_symbol: string;
  asset_type: AssetType;
  quantity: number;
  average_cost_brl: number;
  average_cost_usd: number;
  total_cost_brl: number;
  total_cost_usd: number;
  created_at: string;
  user_id: string;
}

export interface RealizedPNL {
  id: string;
  asset_symbol: string;
  asset_type: AssetType;
  quantity: number;
  average_cost_brl: number;
  average_cost_usd: number;
  selling_price_brl: number;
  selling_price_usd: number;
  total_cost_brl: number;
  total_cost_usd: number;
  total_value_sold_brl: number;
  total_value_sold_usd: number;
  realized_profit_brl: number;
  realized_profit_usd: number;
  created_at: string;
  user_id: string;
}

export interface DashboardSummary {
  totalRealizedPnlBrl: number;
  totalRealizedPnlUsd: number;
  totalActivePositions: number;
  assetAllocation: {
    CRYPTO: number;
    ETF: number;
    STOCK: number;
  };
}

export interface DashboardData {
  summary: DashboardSummary;
  positions: Position[];
  realizedPnl: RealizedPNL[];
  lastUpdated: Date;
}
