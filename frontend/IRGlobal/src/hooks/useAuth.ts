import { useState } from "react";
import type {
  LoginCredentials,
  AuthUser,
  AuthError,
  RegisterCredentials,
} from "../types/auth.types";
import { authService } from "../services/auth.service";
import { validateLoginForm } from "../utils/validation";

interface UseAuthReturn {
  user: AuthUser | null;
  isLoading: boolean;
  error: string | null;
  fieldErrors: Record<string, string>;
  register: (credentials: RegisterCredentials) => Promise<void>;
  login: (credentials: LoginCredentials) => Promise<void>;
  logout: () => void;
  clearError: () => void;
}

export const useAuth = (): UseAuthReturn => {
  const [user, setUser] = useState<AuthUser | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({});

  const register = async (credentials: RegisterCredentials): Promise<void> => {
    // Limpa erros anteriores
    setError(null);
    setFieldErrors({});

    // Validação
    const validation = validateLoginForm(credentials);
    if (!validation.isValid) {
      setFieldErrors(validation.errors);
      return;
    }

    setIsLoading(true);

    try {
      const response = await authService.register(credentials);
      setUser(response.user);
      localStorage.setItem("authToken", response.token);
      console.log("Register successful");
    } catch (err) {
      const authError = err as AuthError;
      setError(authError.message);
    } finally {
      setIsLoading(false);
    }
  };

  const login = async (credentials: LoginCredentials): Promise<void> => {
    // Limpa erros anteriores
    setError(null);
    setFieldErrors({});

    // Validação
    const validation = validateLoginForm(credentials);
    if (!validation.isValid) {
      setFieldErrors(validation.errors);
      return;
    }

    setIsLoading(true);

    try {
      const response = await authService.login(credentials);
      setUser(response.user);
      localStorage.setItem("authToken", response.token);
      console.log("Login successful");
    } catch (err) {
      const authError = err as AuthError;
      setError(authError.message);
    } finally {
      setIsLoading(false);
    }
  };

  const logout = (): void => {
    setUser(null);
    authService.logout();
  };

  const clearError = (): void => {
    setError(null);
    setFieldErrors({});
  };

  return {
    user,
    isLoading,
    error,
    fieldErrors,
    register,
    login,
    logout,
    clearError,
  };
};
