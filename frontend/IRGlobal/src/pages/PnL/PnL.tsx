import { useNavigate } from "react-router-dom";
import {
  Layout,
  Button,
  RealizedPnlTable,
  LoadingSpinner,
} from "../../components";
import { TabNavigation } from "../../components/TabNavigation/TabNavigation";
import { useRealizedPnl } from "../../hooks/useDashboard";

function PnL() {
  const navigate = useNavigate();
  const { realizedPnl, isLoading, error, refetch } = useRealizedPnl();

  const handleLogout = () => {
    localStorage.removeItem("authToken");
    navigate("/login");
  };

  const tabs = [
    { name: "Dashboard", path: "/dashboard" },
    { name: "Transações", path: "/transactions" },
    { name: "Posições", path: "/positions" },
    { name: "Ganhos e Perdas", path: "/pnl" },
  ];

  // Calcular estatísticas
  const totalPnlBrl = realizedPnl.reduce(
    (sum, pnl) => sum + pnl.realized_profit_brl,
    0
  );
  const totalPnlUsd = realizedPnl.reduce(
    (sum, pnl) => sum + pnl.realized_profit_usd,
    0
  );
  const positivePnl = realizedPnl.filter(
    (pnl) => pnl.realized_profit_brl > 0
  ).length;
  const negativePnl = realizedPnl.filter(
    (pnl) => pnl.realized_profit_brl < 0
  ).length;

  const formatCurrency = (value: number, currency: "BRL" | "USD" = "BRL") => {
    return value.toLocaleString("pt-BR", {
      style: "currency",
      currency,
      minimumFractionDigits: 2,
    });
  };

  // Estados de loading e erro
  if (isLoading) {
    return (
      <Layout>
        <TabNavigation tabs={tabs} />
        <LoadingSpinner size="lg" message="Carregando PNL realizado..." />
      </Layout>
    );
  }

  if (error) {
    return (
      <Layout>
        <TabNavigation tabs={tabs} />
        <div className="py-8 px-4 sm:px-6 lg:px-8">
          <div className="max-w-7xl mx-auto">
            <div className="bg-red-50 border border-red-200 rounded-lg p-6">
              <h3 className="text-lg font-medium text-red-800 mb-2">
                Erro ao carregar PNL realizado
              </h3>
              <p className="text-red-600 mb-4">{error}</p>
              <Button onClick={refetch} variant="primary" size="sm">
                Tentar novamente
              </Button>
            </div>
          </div>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <TabNavigation tabs={tabs} />

      <div className="py-8 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          {/* Header */}
          <div className="flex justify-between items-center mb-8">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">
                Ganhos e Perdas
              </h1>
              <p className="text-gray-600 mt-1">
                Visualize todo o seu PNL realizado
              </p>
            </div>
            <div className="flex gap-2">
              <Button onClick={refetch} variant="secondary" size="sm">
                Atualizar
              </Button>
              <Button onClick={handleLogout} variant="secondary" size="sm">
                Logout
              </Button>
            </div>
          </div>

          {/* Estatísticas de PNL */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-sm font-medium text-gray-500 mb-2">
                PNL Total (BRL)
              </h3>
              <p
                className={`text-2xl font-bold ${totalPnlBrl >= 0 ? "text-green-600" : "text-red-600"}`}
              >
                {formatCurrency(totalPnlBrl, "BRL")}
              </p>
            </div>

            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-sm font-medium text-gray-500 mb-2">
                PNL Total (USD)
              </h3>
              <p
                className={`text-2xl font-bold ${totalPnlUsd >= 0 ? "text-green-600" : "text-red-600"}`}
              >
                {formatCurrency(totalPnlUsd, "USD")}
              </p>
            </div>

            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-sm font-medium text-gray-500 mb-2">
                Operações Positivas
              </h3>
              <p className="text-2xl font-bold text-green-600">{positivePnl}</p>
            </div>

            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-sm font-medium text-gray-500 mb-2">
                Operações Negativas
              </h3>
              <p className="text-2xl font-bold text-red-600">{negativePnl}</p>
            </div>
          </div>

          {/* Tabela de PNL realizado */}
          <RealizedPnlTable
            realizedPnl={realizedPnl}
            title="Histórico de PNL Realizado"
          />
        </div>
      </div>
    </Layout>
  );
}

export default PnL;
