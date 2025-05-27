import { useState, useEffect } from "react";
import type {
  Transaction,
  AddTransactionRequest,
} from "../types/transaction.types";
import { transactionService } from "../services/transactionService";

interface UseTransactionsReturn {
  transactions: Transaction[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

export const useTransactions = (): UseTransactionsReturn => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchTransactions = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await transactionService.getTransactions();
      setTransactions(data);
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Erro ao carregar transações"
      );
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchTransactions();
  }, []);

  return {
    transactions,
    isLoading,
    error,
    refetch: fetchTransactions,
  };
};

interface UseAddTransactionReturn {
  isLoading: boolean;
  error: string | null;
  addTransaction: (data: AddTransactionRequest) => Promise<boolean>;
  clearError: () => void;
}

export const useAddTransaction = (): UseAddTransactionReturn => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const addTransaction = async (
    data: AddTransactionRequest
  ): Promise<boolean> => {
    try {
      setIsLoading(true);
      setError(null);
      await transactionService.addTransaction(data);
      return true;
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Erro ao adicionar transação"
      );
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  const clearError = () => {
    setError(null);
  };

  return {
    isLoading,
    error,
    addTransaction,
    clearError,
  };
};

interface UseDeleteTransactionReturn {
  isLoading: boolean;
  error: string | null;
  deleteTransaction: (transactionId: string) => Promise<boolean>;
  clearError: () => void;
}

export const useDeleteTransaction = (): UseDeleteTransactionReturn => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const deleteTransaction = async (transactionId: string): Promise<boolean> => {
    try {
      setIsLoading(true);
      setError(null);
      await transactionService.deleteTransaction(transactionId);
      return true;
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Erro ao deletar transação"
      );
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  const clearError = () => {
    setError(null);
  };

  return {
    isLoading,
    error,
    deleteTransaction,
    clearError,
  };
};
