import type {
  DashboardData,
  Position,
  RealizedPNL,
  DashboardSummary,
} from "../types/portfolio.types";
import { appConfig, getEndpointUrl } from "../config/app.config";

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
  // Busca dados do dashboard
  async getDashboardData(): Promise<DashboardData> {
    const [positions, realizedPnl] = await Promise.all([
      this.getPositions(),
      this.getRealizedPnl(),
    ]);

    const summary = calculateDashboardSummary(positions, realizedPnl);

    return {
      summary,
      positions,
      realizedPnl,
      lastUpdated: new Date(),
    };
  }

  // Busca posições
  async getPositions(): Promise<Position[]> {
    const token = localStorage.getItem(appConfig.storage.authToken);
    const url = getEndpointUrl("positions");

    const response = await fetch(url, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      signal: AbortSignal.timeout(appConfig.api.timeout),
    });

    if (!response.ok) {
      throw new Error(
        `Erro ao buscar posições: ${response.status} ${response.statusText}`
      );
    }

    return response.json();
  }

  // Busca PNL realizado
  async getRealizedPnl(): Promise<RealizedPNL[]> {
    const token = localStorage.getItem(appConfig.storage.authToken);
    const url = getEndpointUrl("realizedPnl");

    const response = await fetch(url, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      signal: AbortSignal.timeout(appConfig.api.timeout),
    });

    if (!response.ok) {
      throw new Error(
        `Erro ao buscar PNL realizado: ${response.status} ${response.statusText}`
      );
    }

    return response.json();
  }

  // Busca resumo do dashboard
  async getDashboardSummary(): Promise<DashboardSummary> {
    const [positions, realizedPnl] = await Promise.all([
      this.getPositions(),
      this.getRealizedPnl(),
    ]);
    return calculateDashboardSummary(positions, realizedPnl);
  }
}

export const dashboardService = new DashboardService();
