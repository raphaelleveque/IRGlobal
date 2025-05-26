import type { LoginFormData } from "../types/auth.types";

export interface ValidationResult {
  isValid: boolean;
  errors: Record<string, string>;
}

export const validateLoginForm = (data: LoginFormData): ValidationResult => {
  const errors: Record<string, string> = {};

  // Validação de email
  if (!data.email) {
    errors.email = "Email é obrigatório";
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(data.email)) {
    errors.email = "Email inválido";
  }

  // Validação de senha
  if (!data.password) {
    errors.password = "Senha é obrigatória";
  } else if (data.password.length < 6) {
    errors.password = "Senha deve ter pelo menos 6 caracteres";
  }

  return {
    isValid: Object.keys(errors).length === 0,
    errors,
  };
};
