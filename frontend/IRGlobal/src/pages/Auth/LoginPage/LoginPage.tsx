import React, { useState } from "react";
import { Link } from "react-router-dom";
import { useAuthNavigation } from "../../../hooks/useAuthNavigation";
import { Button, Input, Logo } from "../../../components";
import type { LoginFormData } from "../../../types/auth.types";

// Este é nosso componente LoginPage
function LoginPage() {
  const [formData, setFormData] = useState<LoginFormData>({
    email: "",
    password: "",
  });

  const { isLoading, error, fieldErrors, loginWithNavigation } =
    useAuthNavigation();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    await loginWithNavigation(formData);
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 px-4 sm:px-6 lg:px-8 relative">
      <div
        className="absolute top-0 left-0 right-0 flex justify-center"
        style={{ top: "calc(50% - 300px)" }}
      >
        <Logo size="xlg" />
      </div>

      <div className="max-w-md w-full space-y-8">
        <div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Faça login na sua conta
          </h2>
        </div>

        <form onSubmit={handleSubmit} className="mt-8 space-y-6">
          <div className="space-y-4">
            <Input
              name="email"
              type="email"
              label="Email"
              placeholder="Digite seu email"
              value={formData.email}
              onChange={handleInputChange}
              error={fieldErrors.email}
              required
            />

            <Input
              name="password"
              type="password"
              label="Senha"
              placeholder="Digite sua senha"
              value={formData.password}
              onChange={handleInputChange}
              error={fieldErrors.password}
              required
            />
          </div>

          {error && (
            <div className="text-red-600 text-sm text-center">{error}</div>
          )}

          <Button
            type="submit"
            isLoading={isLoading}
            className="w-full flex justify-center py-2"
          >
            Entrar
          </Button>

          <div className="text-center">
            <p className="text-sm text-gray-600">
              Não possui uma conta?{" "}
              <Link
                to="/register"
                className="font-medium text-indigo-600 hover:text-indigo-500"
              >
                Cadastre-se
              </Link>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
}

export default LoginPage;
