import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Layout, Button, Input } from "../../../components";
import { useAddTransaction } from "../../../hooks/useTransactions";
import type { AddTransactionRequest } from "../../../types/transaction.types";

function AddTransaction() {
  const navigate = useNavigate();
  const { isLoading, error, addTransaction, clearError } = useAddTransaction();

  const [formData, setFormData] = useState<AddTransactionRequest>({
    asset_symbol: "",
    asset_type: "STOCK",
    quantity: 0,
    price_in_usd: 0,
    type: "BUY",
    operation_date: "",
  });

  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({});

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;

    // Limpar erro do campo quando o usuário começar a digitar
    if (fieldErrors[name]) {
      setFieldErrors((prev) => ({ ...prev, [name]: "" }));
    }

    setFormData((prev) => ({
      ...prev,
      [name]:
        name === "quantity" || name === "price_in_usd"
          ? parseFloat(value) || 0
          : value,
    }));
  };

  const validateForm = (): boolean => {
    const errors: Record<string, string> = {};

    if (!formData.asset_symbol.trim()) {
      errors.asset_symbol = "Símbolo do ativo é obrigatório";
    }

    if (formData.quantity <= 0) {
      errors.quantity = "Quantidade deve ser maior que zero";
    }

    if (formData.price_in_usd <= 0) {
      errors.price_in_usd = "Preço deve ser maior que zero";
    }

    if (!formData.operation_date) {
      errors.operation_date = "Data da operação é obrigatória";
    }

    setFieldErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateForm()) {
      return;
    }

    clearError();

    const success = await addTransaction(formData);

    if (success) {
      navigate("/transactions");
    }
  };

  const handleCancel = () => {
    navigate("/transactions");
  };

  return (
    <Layout>
      <div className="py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-2xl mx-auto">
          {/* Header */}
          <div className="mb-8">
            <h1 className="text-3xl font-bold text-gray-900">
              Adicionar Transação
            </h1>
            <p className="text-gray-600 mt-2">
              Registre uma nova transação de investimento
            </p>
          </div>

          {/* Form */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <form onSubmit={handleSubmit} className="space-y-6">
              {/* Asset Symbol */}
              <Input
                name="asset_symbol"
                type="text"
                label="Símbolo do Ativo"
                placeholder="Ex: AAPL, BTC, PETR4"
                value={formData.asset_symbol}
                onChange={handleInputChange}
                error={fieldErrors.asset_symbol}
                required
              />

              {/* Asset Type */}
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Tipo de Ativo
                </label>
                <select
                  name="asset_type"
                  value={formData.asset_type}
                  onChange={handleInputChange}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  required
                >
                  <option value="STOCK">Ação (STOCK)</option>
                  <option value="CRYPTO">Criptomoeda (CRYPTO)</option>
                  <option value="ETF">ETF</option>
                </select>
              </div>

              {/* Operation Type */}
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Tipo de Operação
                </label>
                <select
                  name="type"
                  value={formData.type}
                  onChange={handleInputChange}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  required
                >
                  <option value="BUY">Compra (BUY)</option>
                  <option value="SELL">Venda (SELL)</option>
                </select>
              </div>

              {/* Quantity */}
              <Input
                name="quantity"
                type="number"
                label="Quantidade"
                placeholder="0"
                value={formData.quantity.toString()}
                onChange={handleInputChange}
                error={fieldErrors.quantity}
                min="0"
                step="any"
                required
              />

              {/* Price in USD */}
              <Input
                name="price_in_usd"
                type="number"
                label="Preço em USD"
                placeholder="0.00"
                value={formData.price_in_usd.toString()}
                onChange={handleInputChange}
                error={fieldErrors.price_in_usd}
                min="0"
                step="0.01"
                required
              />

              {/* Operation Date */}
              <Input
                name="operation_date"
                type="date"
                label="Data da Operação"
                value={formData.operation_date}
                onChange={handleInputChange}
                error={fieldErrors.operation_date}
                required
              />

              {/* Error Message */}
              {error && (
                <div className="bg-red-50 border border-red-200 rounded-lg p-4">
                  <p className="text-red-600">{error}</p>
                </div>
              )}

              {/* Buttons */}
              <div className="flex gap-4 pt-6">
                <Button type="submit" isLoading={isLoading} className="flex-1">
                  Adicionar Transação
                </Button>
                <Button
                  type="button"
                  variant="secondary"
                  onClick={handleCancel}
                  className="flex-1"
                >
                  Cancelar
                </Button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </Layout>
  );
}

export default AddTransaction;
