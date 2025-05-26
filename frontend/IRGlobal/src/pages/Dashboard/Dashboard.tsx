import { useNavigate } from "react-router-dom";
import { Button } from "../../components/Button/Button";

function Dashboard() {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("authToken");
    navigate("/login");
  };

  const token = localStorage.getItem("authToken") || "null";

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md mx-auto bg-white rounded-lg shadow-md p-6">
        <div className="text-center">
          <h1 className="text-2xl font-bold text-gray-900 mb-4">
            🎉 Dashboard
          </h1>
          <p className="text-gray-600 mb-6">
            Bem-vindo! Você fez login com sucesso.
          </p>

          <div className="bg-green-50 border border-green-200 rounded-md p-4 mb-6">
            <p className="text-green-800 text-sm">
              ✅ Login realizado com sucesso!
            </p>
            {token && (
              <p className="text-green-700 text-xs mt-1">
                Token salvo no localStorage
              </p>
            )}
          </div>

          <Button onClick={handleLogout} variant="secondary" className="w-full">
            Fazer Logout
          </Button>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
