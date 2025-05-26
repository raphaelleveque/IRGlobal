import { useState, useEffect } from "react";
import type {
  DashboardData,
  Position,
  RealizedPNL,
  DashboardSummary,
} from "../types/portfolio.types";
import { dashboardService } from "../services/dashboardService";

interface UseDashboardReturn {
  dashboardData: DashboardData | null;
  summary: DashboardSummary | null;
  positions: Position[];
  realizedPnl: RealizedPNL[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

export const useDashboard = (): UseDashboardReturn => {
  const [dashboardData, setDashboardData] = useState<DashboardData | null>(
    null
  );
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchDashboardData = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await dashboardService.getDashboardData();
      setDashboardData(data);
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : "Erro ao carregar dados do dashboard"
      );
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchDashboardData();
  }, []);

  return {
    dashboardData,
    summary: dashboardData?.summary || null,
    positions: dashboardData?.positions || [],
    realizedPnl: dashboardData?.realizedPnl || [],
    isLoading,
    error,
    refetch: fetchDashboardData,
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
      const data = await dashboardService.getPositions();
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

// Hook específico para PNL realizado
export const useRealizedPnl = () => {
  const [realizedPnl, setRealizedPnl] = useState<RealizedPNL[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchRealizedPnl = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await dashboardService.getRealizedPnl();
      setRealizedPnl(data);
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Erro ao carregar PNL realizado"
      );
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchRealizedPnl();
  }, []);

  return {
    realizedPnl,
    isLoading,
    error,
    refetch: fetchRealizedPnl,
  };
};
