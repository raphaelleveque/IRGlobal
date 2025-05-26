import type {
  DashboardData,
  Position,
  RealizedPNL,
  DashboardSummary,
} from "../types/portfolio.types";

// Mock data - será substituído por chamadas reais de API
const mockPositions: Position[] = [
  {
    id: "1",
    asset_symbol: "BTC",
    asset_type: "CRYPTO",
    quantity: 0.5,
    average_cost_brl: 150000,
    average_cost_usd: 30000,
    total_cost_brl: 75000,
    total_cost_usd: 15000,
    created_at: "2024-01-15T10:00:00Z",
    user_id: "user123",
  },
  {
    id: "2",
    asset_symbol: "PETR4",
    asset_type: "STOCK",
    quantity: 100,
    average_cost_brl: 28.5,
    average_cost_usd: 5.7,
    total_cost_brl: 2850,
    total_cost_usd: 570,
    created_at: "2024-01-10T14:30:00Z",
    user_id: "user123",
  },
  {
    id: "3",
    asset_symbol: "IVVB11",
    asset_type: "ETF",
    quantity: 50,
    average_cost_brl: 280,
    average_cost_usd: 56,
    total_cost_brl: 14000,
    total_cost_usd: 2800,
    created_at: "2024-01-05T09:15:00Z",
    user_id: "user123",
  },
];

const mockRealizedPnl: RealizedPNL[] = [
  {
    id: "1",
    asset_symbol: "ETH",
    asset_type: "CRYPTO",
    quantity: 2,
    average_cost_brl: 8000,
    average_cost_usd: 1600,
    selling_price_brl: 9500,
    selling_price_usd: 1900,
    total_cost_brl: 16000,
    total_cost_usd: 3200,
    total_value_sold_brl: 19000,
    total_value_sold_usd: 3800,
    realized_profit_brl: 3000,
    realized_profit_usd: 600,
    created_at: "2024-01-20T16:45:00Z",
    user_id: "user123",
  },
  {
    id: "2",
    asset_symbol: "VALE3",
    asset_type: "STOCK",
    quantity: 200,
    average_cost_brl: 65.8,
    average_cost_usd: 13.16,
    selling_price_brl: 68.45,
    selling_price_usd: 13.69,
    total_cost_brl: 13160,
    total_cost_usd: 2632,
    total_value_sold_brl: 13690,
    total_value_sold_usd: 2738,
    realized_profit_brl: 530,
    realized_profit_usd: 106,
    created_at: "2024-01-18T11:20:00Z",
    user_id: "user123",
  },
];

// Função para calcular o resumo do dashboard baseado nas posições e PNL realizado
const calculateDashboardSummary = (
  positions: Position[],
  realizedPnl: RealizedPNL[]
): DashboardSummary => {
  // Calcular PNL realizado total
  const totalRealizedPnlBrl = realizedPnl.reduce(
    (sum, pnl) => sum + pnl.realized_profit_brl,
    0
  );
  const totalRealizedPnlUsd = realizedPnl.reduce(
    (sum, pnl) => sum + pnl.realized_profit_usd,
    0
  );

  // Calcular alocação por tipo de ativo
  const totalCostBrl = positions.reduce(
    (sum, pos) => sum + pos.total_cost_brl,
    0
  );

  const assetAllocation = {
    CRYPTO: 0,
    ETF: 0,
    STOCK: 0,
  };

  if (totalCostBrl > 0) {
    positions.forEach((position) => {
      const percentage = (position.total_cost_brl / totalCostBrl) * 100;
      assetAllocation[position.asset_type] += percentage;
    });
  }

  return {
    totalRealizedPnlBrl,
    totalRealizedPnlUsd,
    totalActivePositions: positions.length,
    assetAllocation,
  };
};

class DashboardService {
  // Simula uma chamada de API para buscar dados do dashboard
  async getDashboardData(): Promise<DashboardData> {
    // Simula delay de rede
    await new Promise((resolve) => setTimeout(resolve, 500));

    const positions = mockPositions;
    const realizedPnl = mockRealizedPnl;
    const summary = calculateDashboardSummary(positions, realizedPnl);

    return {
      summary,
      positions,
      realizedPnl,
      lastUpdated: new Date(),
    };
  }

  // Simula uma chamada de API para buscar apenas as posições
  async getPositions(): Promise<Position[]> {
    await new Promise((resolve) => setTimeout(resolve, 300));
    return mockPositions;
  }

  // Simula uma chamada de API para buscar apenas o PNL realizado
  async getRealizedPnl(): Promise<RealizedPNL[]> {
    await new Promise((resolve) => setTimeout(resolve, 300));
    return mockRealizedPnl;
  }

  // Simula uma chamada de API para buscar apenas o resumo
  async getDashboardSummary(): Promise<DashboardSummary> {
    await new Promise((resolve) => setTimeout(resolve, 200));
    return calculateDashboardSummary(mockPositions, mockRealizedPnl);
  }

  // Métodos para quando integrar com API real
  async getPositionsFromAPI(): Promise<Position[]> {
    const token = localStorage.getItem("authToken");

    const response = await fetch("/api/positions", {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Erro ao buscar posições");
    }

    return response.json();
  }

  async getRealizedPnlFromAPI(): Promise<RealizedPNL[]> {
    const token = localStorage.getItem("authToken");

    const response = await fetch("/api/realized-pnl", {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Erro ao buscar PNL realizado");
    }

    return response.json();
  }

  async getDashboardDataFromAPI(): Promise<DashboardData> {
    const [positions, realizedPnl] = await Promise.all([
      this.getPositionsFromAPI(),
      this.getRealizedPnlFromAPI(),
    ]);

    const summary = calculateDashboardSummary(positions, realizedPnl);

    return {
      summary,
      positions,
      realizedPnl,
      lastUpdated: new Date(),
    };
  }
}

export const dashboardService = new DashboardService();
