# IRGlobal - Sistema de Imposto de Renda para Investidores

Sistema web para gerenciamento de impostos sobre investimentos, desenvolvido com React/TypeScript no frontend e Go no backend.

## 🚀 Configuração da API

O sistema está configurado para usar a API real em `localhost:8080` com os seguintes endpoints:

### Endpoints de Autenticação

- `POST /auth/login` - Login do usuário
- `POST /auth/register` - Registro de novo usuário

### Endpoints de Dados

- `GET /RealizedPNL/get` - Lista de PNL realizado
- `GET /position/get` - Lista de posições ativas

## 🛠️ Tecnologias

### Frontend

- React 18
- TypeScript
- Tailwind CSS
- React Router DOM
- Vite

### Backend

- Go
- JWT para autenticação

## 📦 Instalação e Execução

### Frontend

```bash
cd frontend/IRGlobal
npm install
npm run dev
```

### Backend

Certifique-se de que o backend Go está rodando em `localhost:8080`

## 🔧 Configuração

### API Configuration

O arquivo `src/config/app.config.ts` contém todas as configurações da API:

```typescript
export const appConfig = {
  api: {
    baseUrl: "http://localhost:8080",
    timeout: 10000,
    endpoints: {
      auth: {
        login: "/auth/login",
        register: "/auth/register",
      },
      realizedPnl: "/RealizedPNL/get",
      positions: "/position/get",
    },
  },
  // ... outras configurações
};
```

## 📊 Estrutura de Dados

### Position

```json
{
  "id": "string",
  "asset_symbol": "string",
  "asset_type": "CRYPTO" | "ETF" | "STOCK",
  "quantity": 0,
  "average_cost_brl": 0,
  "average_cost_usd": 0,
  "total_cost_brl": 0,
  "total_cost_usd": 0,
  "created_at": "string",
  "user_id": "string"
}
```

### RealizedPNL

```json
{
  "id": "string",
  "asset_symbol": "string",
  "asset_type": "CRYPTO" | "ETF" | "STOCK",
  "quantity": 0,
  "average_cost_brl": 0,
  "average_cost_usd": 0,
  "selling_price_brl": 0,
  "selling_price_usd": 0,
  "total_cost_brl": 0,
  "total_cost_usd": 0,
  "total_value_sold_brl": 0,
  "total_value_sold_usd": 0,
  "realized_profit_brl": 0,
  "realized_profit_usd": 0,
  "created_at": "string",
  "user_id": "string"
}
```

## 🔐 Autenticação

O sistema usa JWT tokens para autenticação. O token é armazenado no localStorage e enviado automaticamente em todas as requisições para endpoints protegidos.

## 📱 Funcionalidades

- **Dashboard**: Visão geral com PNL realizado e alocação de ativos
- **Posições**: Lista de posições ativas por tipo de ativo
- **PNL**: Histórico completo de PNL realizado
- **Autenticação**: Login e registro de usuários

## 🏗️ Arquitetura

### Serviços

- `authService`: Gerencia autenticação (login, register, logout)
- `dashboardService`: Gerencia dados do dashboard (posições, PNL realizado)

### Hooks

- `useAuth`: Hook para autenticação
- `useDashboard`: Hook para dados do dashboard
- `usePositions`: Hook específico para posições
- `useRealizedPnl`: Hook específico para PNL realizado

### Componentes Reutilizáveis

- `DashboardSummaryCards`: Cards de resumo financeiro
- `PositionsTable`: Tabela de posições
- `RealizedPnlTable`: Tabela de PNL realizado
- `LoadingSpinner`: Componente de loading

## 🚀 Deploy

Para produção, altere a `baseUrl` no arquivo de configuração para apontar para o servidor de produção.
