import { useNavigate } from "react-router-dom";
import {
  Layout,
  LoadingSpinner,
  Button,
  TabNavigation,
} from "../../components";
import { TransactionsTable } from "../../components/TransactionsTable/TransactionsTable";
import {
  useTransactions,
  useDeleteTransaction,
} from "../../hooks/useTransactions";

function Transactions() {
  const navigate = useNavigate();
  const { transactions, isLoading, error, refetch } = useTransactions();
  const { deleteTransaction } = useDeleteTransaction();

  const tabs = [
    { name: "Dashboard", path: "/dashboard" },
    { name: "Transações", path: "/transactions" },
    { name: "Posições", path: "/positions" },
    { name: "Ganhos e Perdas", path: "/pnl" },
    // { name: "Relatórios", path: "/reports" },
  ];

  const handleAddTransaction = () => {
    navigate("/transactions/add");
  };

  const handleDeleteTransaction = async (transactionId: string) => {
    const success = await deleteTransaction(transactionId);
    if (success) {
      // Recarregar a lista de transações após deletar
      await refetch();
    }
  };

  if (isLoading) {
    return (
      <Layout>
        <TabNavigation tabs={tabs} />
        <div className="py-12 px-4 sm:px-6 lg:px-8">
          <div className="max-w-7xl mx-auto">
            <LoadingSpinner size="lg" message="Carregando transações..." />
          </div>
        </div>
      </Layout>
    );
  }

  if (error) {
    return (
      <Layout>
        <TabNavigation tabs={tabs} />
        <div className="py-12 px-4 sm:px-6 lg:px-8">
          <div className="max-w-7xl mx-auto">
            <div className="bg-red-50 border border-red-200 rounded-lg p-6">
              <h3 className="text-lg font-medium text-red-800 mb-2">
                Erro ao carregar transações
              </h3>
              <p className="text-red-600 mb-4">{error}</p>
              <Button onClick={refetch} variant="secondary">
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
      <div className="py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          {/* Header */}
          <div className="flex justify-between items-center mb-8">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">Transações</h1>
              <p className="text-gray-600 mt-2">
                Gerencie todas as suas transações de investimentos
              </p>
            </div>
            <Button onClick={handleAddTransaction}>Adicionar Transação</Button>
          </div>

          {/* Statistics */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                Total de Transações
              </h3>
              <p className="text-3xl font-bold text-blue-600">
                {transactions.length}
              </p>
            </div>
            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                Compras
              </h3>
              <p className="text-3xl font-bold text-green-600">
                {transactions.filter((t) => t.type === "BUY").length}
              </p>
            </div>
            <div className="bg-white rounded-lg shadow-md p-6">
              <h3 className="text-lg font-medium text-gray-900 mb-2">Vendas</h3>
              <p className="text-3xl font-bold text-red-600">
                {transactions.filter((t) => t.type === "SELL").length}
              </p>
            </div>
          </div>

          {/* Transactions Table */}
          <TransactionsTable
            transactions={transactions}
            onDeleteTransaction={handleDeleteTransaction}
          />
        </div>
      </div>
    </Layout>
  );
}

export default Transactions;
