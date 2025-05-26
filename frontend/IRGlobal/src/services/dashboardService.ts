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
  // Garantir que temos arrays válidos
  const safePositions = Array.isArray(positions) ? positions : [];
  const safeRealizedPnl = Array.isArray(realizedPnl) ? realizedPnl : [];

  // Calcular PNL realizado total
  const totalRealizedPnlBrl = safeRealizedPnl.reduce(
    (sum, pnl) => sum + (pnl.realized_profit_brl || 0),
    0
  );
  const totalRealizedPnlUsd = safeRealizedPnl.reduce(
    (sum, pnl) => sum + (pnl.realized_profit_usd || 0),
    0
  );

  // Calcular alocação por tipo de ativo
  const totalCostBrl = safePositions.reduce(
    (sum, pos) => sum + (pos.total_cost_brl || 0),
    0
  );

  const assetAllocation = {
    CRYPTO: 0,
    ETF: 0,
    STOCK: 0,
  };

  if (totalCostBrl > 0) {
    safePositions.forEach((position) => {
      const percentage = ((position.total_cost_brl || 0) / totalCostBrl) * 100;
      if (
        position.asset_type &&
        assetAllocation[position.asset_type] !== undefined
      ) {
        assetAllocation[position.asset_type] += percentage;
      }
    });
  }

  return {
    totalRealizedPnlBrl,
    totalRealizedPnlUsd,
    totalActivePositions: safePositions.length,
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

    const data = await response.json();
    console.log("API Response - Positions:", data);

    // Verificar se a resposta tem a estrutura esperada
    if (data && Array.isArray(data.positions)) {
      return data.positions;
    }

    // Se a resposta for um array direto
    if (Array.isArray(data)) {
      return data;
    }

    // Caso contrário, retornar array vazio
    console.warn("Formato inesperado da resposta de posições:", data);
    return [];
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

    const data = await response.json();
    console.log("API Response - Realized PNL:", data);

    // Verificar se a resposta tem a estrutura esperada (pnls)
    if (data && Array.isArray(data.pnls)) {
      return data.pnls;
    }

    // Verificar se a resposta tem a estrutura alternativa (realized_pnl)
    if (data && Array.isArray(data.realized_pnl)) {
      return data.realized_pnl;
    }

    // Se a resposta for um array direto
    if (Array.isArray(data)) {
      return data;
    }

    // Caso contrário, retornar array vazio
    console.warn("Formato inesperado da resposta de PNL realizado:", data);
    return [];
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
