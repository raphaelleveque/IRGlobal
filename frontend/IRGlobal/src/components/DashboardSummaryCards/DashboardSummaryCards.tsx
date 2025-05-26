import React from "react";
import type { DashboardSummary } from "../../types/portfolio.types";

interface DashboardSummaryCardsProps {
  summary: DashboardSummary;
}

export const DashboardSummaryCards: React.FC<DashboardSummaryCardsProps> = ({
  summary,
}) => {
  const formatCurrency = (value: number, currency: "BRL" | "USD" = "BRL") => {
    return value.toLocaleString("pt-BR", {
      style: "currency",
      currency,
      minimumFractionDigits: 2,
    });
  };

  const formatPercentage = (value: number) => {
    return `${value.toFixed(1)}%`;
  };

  const getAssetTypeLabel = (type: string) => {
    const labels = {
      CRYPTO: "Crypto",
      ETF: "ETFs",
      STOCK: "Ações",
    };
    return labels[type as keyof typeof labels] || type;
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      {/* PNL Realizado BRL */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">
          PNL Realizado (BRL)
        </h3>
        <p
          className={`text-2xl font-bold ${
            summary.totalRealizedPnlBrl >= 0 ? "text-green-600" : "text-red-600"
          }`}
        >
          {formatCurrency(summary.totalRealizedPnlBrl, "BRL")}
        </p>
      </div>

      {/* PNL Realizado USD */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">
          PNL Realizado (USD)
        </h3>
        <p
          className={`text-2xl font-bold ${
            summary.totalRealizedPnlUsd >= 0 ? "text-green-600" : "text-red-600"
          }`}
        >
          {formatCurrency(summary.totalRealizedPnlUsd, "USD")}
        </p>
      </div>

      {/* Posições Ativas */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">
          Posições Ativas
        </h3>
        <p className="text-2xl font-bold text-gray-900">
          {summary.totalActivePositions}
        </p>
        <p className="text-sm text-gray-500">ativos em carteira</p>
      </div>

      {/* Alocação de Ativos */}
      <div className="bg-white rounded-lg shadow-md p-6">
        <h3 className="text-sm font-medium text-gray-500 mb-2">
          Alocação de Ativos
        </h3>
        <div className="space-y-2">
          {Object.entries(summary.assetAllocation).map(([type, percentage]) => (
            <div key={type} className="flex justify-between items-center">
              <span className="text-sm text-gray-600">
                {getAssetTypeLabel(type)}:
              </span>
              <span className="text-sm font-medium text-gray-900">
                {formatPercentage(percentage)}
              </span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
