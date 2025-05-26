import React from "react";
import type { PortfolioSummary } from "../../types/portfolio.types";

interface PortfolioSummaryCardsProps {
  summary: PortfolioSummary;
}

export const PortfolioSummaryCards: React.FC<PortfolioSummaryCardsProps> = ({
  summary,
}) => {
  const formatCurrency = (value: number) => {
    return value.toLocaleString("pt-BR", {
      style: "currency",
      currency: "BRL",
      minimumFractionDigits: 2,
    });
  };

  const formatPercentage = (value: number) => {
    return `${value >= 0 ? "+" : ""}${value.toFixed(2)}%`;
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">
          Valor Total da Carteira
        </h3>
        <p className="text-2xl font-bold text-gray-900">
          {formatCurrency(summary.totalValue)}
        </p>
        <p className="text-sm text-gray-500 mt-1">
          Investido: {formatCurrency(summary.totalInvested)}
        </p>
      </div>

      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">PNL Total</h3>
        <p
          className={`text-2xl font-bold ${summary.totalPnL >= 0 ? "text-green-600" : "text-red-600"}`}
        >
          {formatCurrency(summary.totalPnL)}
        </p>
        <p
          className={`text-sm ${summary.pnlPercentage >= 0 ? "text-green-600" : "text-red-600"}`}
        >
          {formatPercentage(summary.pnlPercentage)}
        </p>
      </div>

      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">PNL do Dia</h3>
        <p
          className={`text-2xl font-bold ${summary.dailyPnL >= 0 ? "text-green-600" : "text-red-600"}`}
        >
          {formatCurrency(summary.dailyPnL)}
        </p>
        <p
          className={`text-sm ${summary.dailyPnLPercentage >= 0 ? "text-green-600" : "text-red-600"}`}
        >
          {formatPercentage(summary.dailyPnLPercentage)}
        </p>
      </div>

      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">
          Posições Ativas
        </h3>
        <p className="text-2xl font-bold text-gray-900">
          {summary.totalPositions}
        </p>
        <p className="text-sm text-gray-500">ativos em carteira</p>
      </div>
    </div>
  );
};
