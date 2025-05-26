import { useState, useEffect } from "react";
import type {
  PortfolioData,
  Position,
  PortfolioSummary,
} from "../types/portfolio.types";
import { portfolioService } from "../services/portfolioService";

interface UsePortfolioReturn {
  portfolioData: PortfolioData | null;
  summary: PortfolioSummary | null;
  positions: Position[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

export const usePortfolio = (): UsePortfolioReturn => {
  const [portfolioData, setPortfolioData] = useState<PortfolioData | null>(
    null
  );
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchPortfolioData = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await portfolioService.getPortfolioData();
      setPortfolioData(data);
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : "Erro ao carregar dados do portfolio"
      );
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchPortfolioData();
  }, []);

  return {
    portfolioData,
    summary: portfolioData?.summary || null,
    positions: portfolioData?.positions || [],
    isLoading,
    error,
    refetch: fetchPortfolioData,
  };
};

// Hook específico para posições (pode ser usado na página de Posições)
export const usePositions = () => {
  const [positions, setPositions] = useState<Position[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchPositions = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await portfolioService.getPositions();
      setPositions(data);
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Erro ao carregar posições"
      );
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchPositions();
  }, []);

  return {
    positions,
    isLoading,
    error,
    refetch: fetchPositions,
  };
};
