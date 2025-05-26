import type {
  Transaction,
  AddTransactionRequest,
  TransactionResponse,
  AddTransactionResponse,
} from "../types/transaction.types";
import { appConfig, getEndpointUrl } from "../config/app.config";

class TransactionService {
  // Busca todas as transações
  async getTransactions(): Promise<Transaction[]> {
    const token = localStorage.getItem(appConfig.storage.authToken);
    const url = getEndpointUrl("transactions", "get");

    const response = await fetch(url, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      signal: AbortSignal.timeout(appConfig.api.timeout),
    });

    if (!response.ok) {
      throw new Error(
        `Erro ao buscar transações: ${response.status} ${response.statusText}`
      );
    }

    const data: TransactionResponse = await response.json();
    console.log("API Response - Transactions:", data);

    // Garantir que sempre retornamos um array
    return Array.isArray(data.transactions) ? data.transactions : [];
  }

  // Adiciona uma nova transação
  async addTransaction(
    transactionData: AddTransactionRequest
  ): Promise<AddTransactionResponse> {
    const token = localStorage.getItem(appConfig.storage.authToken);
    const url = getEndpointUrl("transactions", "add");

    const response = await fetch(url, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(transactionData),
      signal: AbortSignal.timeout(appConfig.api.timeout),
    });

    if (!response.ok) {
      throw new Error(
        `Erro ao adicionar transação: ${response.status} ${response.statusText}`
      );
    }

    return response.json();
  }

  // Deleta uma transação
  async deleteTransaction(transactionId: string): Promise<void> {
    const token = localStorage.getItem(appConfig.storage.authToken);
    const url = getEndpointUrl("transactions", "delete");

    const response = await fetch(url, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id: transactionId }),
      signal: AbortSignal.timeout(appConfig.api.timeout),
    });

    if (!response.ok) {
      throw new Error(
        `Erro ao deletar transação: ${response.status} ${response.statusText}`
      );
    }
  }
}

export const transactionService = new TransactionService();
