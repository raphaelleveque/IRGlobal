import type {
  PortfolioData,
  Position,
  PortfolioSummary,
} from "../types/portfolio.types";

// Mock data - será substituído por chamadas reais de API
const mockPositions: Position[] = [
  {
    symbol: "PETR4",
    quantity: 100,
    avgPrice: 28.5,
    currentPrice: 31.2,
    pnl: 270.0,
    pnlPercentage: 9.47,
    marketValue: 3120.0,
    totalCost: 2850.0,
  },
  {
    symbol: "VALE3",
    quantity: 200,
    avgPrice: 65.8,
    currentPrice: 68.45,
    pnl: 530.0,
    pnlPercentage: 4.03,
    marketValue: 13690.0,
    totalCost: 13160.0,
  },
  {
    symbol: "ITUB4",
    quantity: 150,
    avgPrice: 24.3,
    currentPrice: 23.85,
    pnl: -67.5,
    pnlPercentage: -1.85,
    marketValue: 3577.5,
    totalCost: 3645.0,
  },
  {
    symbol: "BBDC4",
    quantity: 80,
    avgPrice: 18.75,
    currentPrice: 19.9,
    pnl: 92.0,
    pnlPercentage: 6.13,
    marketValue: 1592.0,
    totalCost: 1500.0,
  },
];

// Função para calcular o resumo do portfolio baseado nas posições
const calculatePortfolioSummary = (positions: Position[]): PortfolioSummary => {
  const totalValue = positions.reduce((sum, pos) => sum + pos.marketValue, 0);
  const totalInvested = positions.reduce((sum, pos) => sum + pos.totalCost, 0);
  const totalPnL = positions.reduce((sum, pos) => sum + pos.pnl, 0);
  const pnlPercentage =
    totalInvested > 0 ? (totalPnL / totalInvested) * 100 : 0;

  return {
    totalValue,
    totalPnL,
    pnlPercentage,
    dailyPnL: 1250.75, // Mock - seria calculado baseado em dados históricos
    dailyPnLPercentage: 1.01, // Mock
    totalPositions: positions.length,
    totalInvested,
  };
};

class PortfolioService {
  // Simula uma chamada de API para buscar dados do portfolio
  async getPortfolioData(): Promise<PortfolioData> {
    // Simula delay de rede
    await new Promise((resolve) => setTimeout(resolve, 500));

    const positions = mockPositions;
    const summary = calculatePortfolioSummary(positions);

    return {
      summary,
      positions,
      lastUpdated: new Date(),
    };
  }

  // Simula uma chamada de API para buscar apenas as posições
  async getPositions(): Promise<Position[]> {
    await new Promise((resolve) => setTimeout(resolve, 300));
    return mockPositions;
  }

  // Simula uma chamada de API para buscar apenas o resumo
  async getPortfolioSummary(): Promise<PortfolioSummary> {
    await new Promise((resolve) => setTimeout(resolve, 200));
    return calculatePortfolioSummary(mockPositions);
  }

  // Método para quando integrar com API real
  async getPortfolioDataFromAPI(): Promise<PortfolioData> {
    const token = localStorage.getItem("authToken");

    const response = await fetch("/api/portfolio", {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Erro ao buscar dados do portfolio");
    }

    return response.json();
  }
}

export const portfolioService = new PortfolioService();
