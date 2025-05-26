import React, { useState } from "react";

// Este é nosso componente LoginPage
function LoginPage() {
  // useState é um Hook do React para gerenciar estado local
  // Aqui criamos estados para armazenar email e senha
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");

  // Função que será executada quando o botão for clicado
  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault(); // Previne o comportamento padrão

    // Limpa erros anteriores
    setError("");

    // Validações básicas
    if (!email || !password) {
      setError("Email e senha são obrigatórios");
      return;
    }

    // Inicia o loading
    setIsLoading(true);

    try {
      // Aqui você faria a chamada para sua API
      console.log("Tentando fazer login com:", { email, password });

      // Simula uma chamada de API
      await new Promise((resolve) => setTimeout(resolve, 2000));

      // Se chegou até aqui, o login foi bem-sucedido
      alert("Login realizado com sucesso!");
    } catch (err) {
      setError("Erro ao fazer login. Tente novamente.");
    } finally {
      // Para o loading independente do resultado
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Faça login na sua conta
          </h2>
        </div>

        <div className="mt-8 space-y-6 text-left">
          <div className="space-y-4">
            {/* Campo de Email */}
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700"
              >
                Email
              </label>
              <input
                id="email"
                name="email"
                type="email"
                required
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="Digite seu email"
              />
            </div>

            {/* Campo de Senha */}
            <div>
              <label
                htmlFor="password"
                className="block text-sm font-medium text-gray-700"
              >
                Senha
              </label>
              <input
                id="password"
                name="password"
                type="password"
                required
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="Digite sua senha"
              />
            </div>
          </div>

          {/* Exibe erros se houver */}
          {error && (
            <div className="text-red-600 text-sm text-center">{error}</div>
          )}

          {/* Botão de Submit */}
          <div>
            <button
              type="submit"
              disabled={isLoading}
              onClick={handleSubmit}
              className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {isLoading ? "Carregando..." : "Entrar"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LoginPage;
