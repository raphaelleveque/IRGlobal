import { useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./useAuth";
import type {
  LoginCredentials,
  RegisterCredentials,
} from "../types/auth.types";

export const useAuthNavigation = () => {
  const navigate = useNavigate();
  const auth = useAuth();
  const previousUser = useRef(auth.user);

  // Monitora mudanças no usuário para redirecionar após login/registro
  useEffect(() => {
    if (!previousUser.current && auth.user && !auth.error) {
      console.log("Autenticação realizada com sucesso! Redirecionando...");
      navigate("/dashboard");
    }
    previousUser.current = auth.user;
  }, [auth.user, auth.error, navigate]);

  const loginWithNavigation = async (
    credentials: LoginCredentials
  ): Promise<void> => {
    await auth.login(credentials);
  };

  const registerWithNavigation = async (
    credentials: RegisterCredentials
  ): Promise<void> => {
    await auth.register(credentials);
  };

  return {
    ...auth,
    loginWithNavigation,
    registerWithNavigation,
  };
};
