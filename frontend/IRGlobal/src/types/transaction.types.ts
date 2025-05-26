export type AssetType = "CRYPTO" | "STOCK" | "ETF";
export type OperationType = "BUY" | "SELL";

export interface Transaction {
  id: string;
  user_id: string;
  asset_symbol: string;
  asset_type: AssetType;
  quantity: number;
  price_in_usd: number;
  usd_brl_rate: number;
  price_in_brl: number;
  total_cost_usd: number;
  total_cost_brl: number;
  type: OperationType;
  operation_date: string;
  created_at: string;
}

export interface AddTransactionRequest {
  asset_symbol: string;
  asset_type: AssetType;
  quantity: number;
  price_in_usd: number;
  type: OperationType;
  operation_date: string;
}

export interface TransactionResponse {
  transactions: Transaction[];
}

export interface AddTransactionResponse {
  transaction: Transaction;
  message: string;
}
