import { useNavigate } from "react-router-dom";
import {
  Layout,
  Button,
  PositionsTable,
  LoadingSpinner,
} from "../../components";
import { TabNavigation } from "../../components/TabNavigation/TabNavigation";
import { usePositions } from "../../hooks/useDashboard";

function Positions() {
  const navigate = useNavigate();
  const { positions, isLoading, error, refetch } = usePositions();

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
        <LoadingSpinner size="lg" message="Carregando posições..." />
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
                Erro ao carregar posições
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
              <h1 className="text-3xl font-bold text-gray-900">Posições</h1>
              <p className="text-gray-600 mt-1">
                Visualize todas as suas posições em ações
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

          {/* Estatísticas rápidas */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-sm font-medium text-gray-500 mb-2">
                Total de Posições
              </h3>
              <p className="text-2xl font-bold text-gray-900">
                {positions.length}
              </p>
            </div>

            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-sm font-medium text-gray-500 mb-2">
                Posições em Crypto
              </h3>
              <p className="text-2xl font-bold text-orange-600">
                {positions.filter((p) => p.asset_type === "CRYPTO").length}
              </p>
            </div>

            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-sm font-medium text-gray-500 mb-2">
                Posições em Ações
              </h3>
              <p className="text-2xl font-bold text-green-600">
                {positions.filter((p) => p.asset_type === "STOCK").length}
              </p>
            </div>
          </div>

          {/* Tabela de posições */}
          <PositionsTable positions={positions} title="Todas as Posições" />
        </div>
      </div>
    </Layout>
  );
}

export default Positions;
