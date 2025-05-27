import React from "react";
import type { Position, AssetType } from "../../types/portfolio.types";

interface PositionsTableProps {
  positions: Position[];
  title?: string;
  showTitle?: boolean;
}

export const PositionsTable: React.FC<PositionsTableProps> = ({
  positions,
  title = "Posições Atuais",
  showTitle = true,
}) => {
  const formatCurrency = (value: number, currency: "BRL" | "USD" = "BRL") => {
    return value.toLocaleString("pt-BR", {
      style: "currency",
      currency,
      minimumFractionDigits: 2,
    });
  };

  const getAssetTypeLabel = (type: AssetType) => {
    const labels = {
      CRYPTO: "Crypto",
      ETF: "ETF",
      STOCK: "Ação",
    };
    return labels[type];
  };

  const getAssetTypeBadgeColor = (type: AssetType) => {
    const colors = {
      CRYPTO: "bg-orange-100 text-orange-800",
      ETF: "bg-blue-100 text-blue-800",
      STOCK: "bg-green-100 text-green-800",
    };
    return colors[type];
  };

  if (positions.length === 0) {
    return (
      <div className="bg-white rounded-lg shadow-md p-6">
        {showTitle && (
          <div className="px-6 py-4 border-b border-gray-200">
            <h2 className="text-lg font-semibold text-gray-900">{title}</h2>
          </div>
        )}
        <div className="text-center py-8">
          <p className="text-gray-500">Nenhuma posição encontrada</p>
        </div>
      </div>
    );
  }

  return (
    <div className="bg-white rounded-lg shadow-md overflow-hidden">
      {showTitle && (
        <div className="px-6 py-4 border-b border-gray-200">
          <h2 className="text-lg font-semibold text-gray-900">{title}</h2>
        </div>
      )}
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Ativo
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Tipo
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Quantidade
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Preço Médio (BRL)
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Preço Médio (USD)
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Custo Total (BRL)
              </th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Custo Total (USD)
              </th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {positions.map((position) => (
              <tr key={position.id} className="hover:bg-gray-50">
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {position.asset_symbol}
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <span
                    className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${getAssetTypeBadgeColor(
                      position.asset_type
                    )}`}
                  >
                    {getAssetTypeLabel(position.asset_type)}
                  </span>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {position.quantity.toLocaleString()}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {formatCurrency(position.average_cost_brl, "BRL")}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {formatCurrency(position.average_cost_usd, "USD")}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  {formatCurrency(position.total_cost_brl, "BRL")}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  {formatCurrency(position.total_cost_usd, "USD")}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};
 