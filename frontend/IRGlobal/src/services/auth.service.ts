import type {
  LoginCredentials,
  LoginResponse,
  AuthError,
  RegisterCredentials,
  RegisterResponse,
} from "../types/auth.types";
import { appConfig, getEndpointUrl } from "../config/app.config";

class AuthService {
  async login(credentials: LoginCredentials): Promise<LoginResponse> {
    try {
      const url = getEndpointUrl("auth", "login");
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(credentials),
        signal: AbortSignal.timeout(appConfig.api.timeout),
      });

      if (!response.ok) {
        throw new Error("Erro na autenticação");
      }

      const data = await response.json();

      // Salvar token no localStorage
      if (data.token) {
        localStorage.setItem(appConfig.storage.authToken, data.token);
      }

      return data;
    } catch (error) {
      const authError: AuthError = {
        message: error instanceof Error ? error.message : "Erro desconhecido",
      };
      throw authError;
    }
  }

  async register(credentials: RegisterCredentials): Promise<RegisterResponse> {
    try {
      const url = getEndpointUrl("auth", "register");
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(credentials),
        signal: AbortSignal.timeout(appConfig.api.timeout),
      });

      if (!response.ok) {
        throw new Error("Erro na autenticação");
      }

      const data = await response.json();

      // Salvar token no localStorage
      if (data.token) {
        localStorage.setItem(appConfig.storage.authToken, data.token);
      }

      return data;
    } catch (error) {
      const authError: AuthError = {
        message: error instanceof Error ? error.message : "Erro desconhecido",
      };
      throw authError;
    }
  }

  async logout(): Promise<void> {
    localStorage.removeItem(appConfig.storage.authToken);
    localStorage.removeItem(appConfig.storage.refreshToken);
  }

  isAuthenticated(): boolean {
    return !!localStorage.getItem(appConfig.storage.authToken);
  }

  getToken(): string | null {
    return localStorage.getItem(appConfig.storage.authToken);
  }
}

export const authService = new AuthService();
