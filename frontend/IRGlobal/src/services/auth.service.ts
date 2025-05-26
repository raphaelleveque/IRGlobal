import type {
  LoginCredentials,
  LoginResponse,
  AuthError,
  RegisterCredentials,
  RegisterResponse,
} from "../types/auth.types";

class AuthService {
  private baseUrl = "http://localhost:8080"; // Configure conforme sua API

  async login(credentials: LoginCredentials): Promise<LoginResponse> {
    try {
      const response = await fetch(`${this.baseUrl}/auth/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(credentials),
      });

      if (!response.ok) {
        throw new Error("Erro na autenticação");
      }

      return await response.json();
    } catch (error) {
      const authError: AuthError = {
        message: error instanceof Error ? error.message : "Erro desconhecido",
      };
      throw authError;
    }
  }

  async register(credentials: RegisterCredentials): Promise<RegisterResponse> {
    try {
      const response = await fetch(`${this.baseUrl}/auth/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(credentials),
      });

      if (!response.ok) {
        throw new Error("Erro na autenticação");
      }

      return await response.json();
    } catch (error) {
      const authError: AuthError = {
        message: error instanceof Error ? error.message : "Erro desconhecido",
      };
      throw authError;
    }
  }

  async logout(): Promise<void> {
    // Implementar logout
    localStorage.removeItem("authToken");
  }

  isAuthenticated(): boolean {
    return !!localStorage.getItem("authToken");
  }
}

export const authService = new AuthService();
