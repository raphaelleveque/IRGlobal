import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import LoginPage from "./pages/Auth/LoginPage/LoginPage";
import RegisterPage from "./pages/Auth/RegisterPage/RegisterPage";
import Dashboard from "./pages/Dashboard/Dashboard";
import Transactions from "./pages/Transactions/Transactions";
import AddTransaction from "./pages/Transactions/AddTransaction/AddTransaction";
import Positions from "./pages/Positions/Positions";
import PnL from "./pages/PnL/PnL";
// import Reports from "./pages/Reports/Reports";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/transactions" element={<Transactions />} />
        <Route path="/transactions/add" element={<AddTransaction />} />
        <Route path="/positions" element={<Positions />} />
        <Route path="/pnl" element={<PnL />} />
        {/* <Route path="/reports" element={<Reports />} /> */}
        <Route path="/" element={<Navigate to="/login" replace />} />
        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
    </Router>
  );
}

export default App;
