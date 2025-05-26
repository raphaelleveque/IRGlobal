import { useNavigate } from "react-router-dom";
import {
  Layout,
  Button,
  DashboardSummaryCards,
  PositionsTable,
  RealizedPnlTable,
  LoadingSpinner,
} from "../../components";
import { TabNavigation } from "../../components/TabNavigation/TabNavigation";
import { useDashboard } from "../../hooks/useDashboard";

function Dashboard() {
  const navigate = useNavigate();
  const { summary, positions, realizedPnl, isLoading, error, refetch } =
    useDashboard();

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

  // Estados de loading e erro
  if (isLoading) {
    return (
      <Layout>
        <TabNavigation tabs={tabs} />
        <LoadingSpinner size="lg" message="Carregando dados do dashboard..." />
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
                Erro ao carregar dados
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

  if (!summary) {
    return (
      <Layout>
        <TabNavigation tabs={tabs} />
        <div className="py-8 px-4 sm:px-6 lg:px-8">
          <div className="max-w-7xl mx-auto">
            <div className="text-center py-8">
              <p className="text-gray-500">Nenhum dado disponível</p>
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
          {/* Header com botão de logout */}
          <div className="flex justify-between items-center mb-8">
            <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
            <div className="flex gap-2">
              <Button onClick={refetch} variant="secondary" size="sm">
                Atualizar
              </Button>
              <Button onClick={handleLogout} variant="secondary" size="sm">
                Logout
              </Button>
            </div>
          </div>

          {/* Cards de resumo do Dashboard */}
          <div className="mb-8">
            <DashboardSummaryCards summary={summary} />
          </div>

          {/* Tabela de posições */}
          <div className="mb-8">
            <PositionsTable positions={positions} />
          </div>

          {/* Tabela de PNL realizado */}
          <RealizedPnlTable realizedPnl={realizedPnl} />
        </div>
      </div>
    </Layout>
  );
}

export default Dashboard;
