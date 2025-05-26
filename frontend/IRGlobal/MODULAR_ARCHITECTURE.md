# Arquitetura Modular - Portfolio Management

## Visão Geral

O código foi refatorado para seguir uma arquitetura mais modular e escalável, preparando-se para integração com APIs reais. A nova estrutura separa responsabilidades e permite reutilização de código entre diferentes páginas.

## Estrutura da Arquitetura

### 1. **Types** (`src/types/`)

Define interfaces TypeScript para tipagem forte dos dados:

- `portfolio.types.ts`: Tipos para dados do portfolio
  - `Position`: Representa uma posição individual
  - `PortfolioSummary`: Resumo consolidado do portfolio
  - `PortfolioData`: Dados completos do portfolio

### 2. **Services** (`src/services/`)

Camada de serviços para comunicação com APIs:

- `portfolioService.ts`: Gerencia todas as operações relacionadas ao portfolio
  - Métodos mock para desenvolvimento
  - Método preparado para API real (`getPortfolioDataFromAPI`)
  - Cálculos automáticos de resumo baseados nas posições

### 3. **Hooks** (`src/hooks/`)

Hooks customizados para gerenciamento de estado:

- `usePortfolio.ts`: Hook principal para dados do portfolio
  - `usePortfolio()`: Dados completos com resumo e posições
  - `usePositions()`: Hook específico apenas para posições
  - Gerenciamento automático de loading, erro e refetch

### 4. **Components** (`src/components/`)

Componentes reutilizáveis:

- `PortfolioSummaryCards`: Cards de resumo financeiro
- `PositionsTable`: Tabela de posições reutilizável
- `LoadingSpinner`: Componente de loading
- `TabNavigation`: Navegação por abas

## Benefícios da Nova Arquitetura

### ✅ **Modularidade**

- Componentes reutilizáveis entre páginas
- Lógica de negócio separada da apresentação
- Fácil manutenção e teste

### ✅ **Preparação para API Real**

- Serviços centralizados para chamadas de API
- Fácil substituição de dados mock por dados reais
- Tratamento consistente de loading e erros

### ✅ **Reutilização de Código**

- Mesmos componentes usados no Dashboard e na página de Posições
- Hooks compartilhados entre diferentes páginas
- Formatação consistente de dados

### ✅ **Tipagem Forte**

- Interfaces TypeScript bem definidas
- Detecção de erros em tempo de desenvolvimento
- Melhor IntelliSense no editor

### ✅ **Gerenciamento de Estado**

- Estados de loading, erro e dados centralizados
- Função de refetch para atualização manual
- Tratamento consistente de erros

## Como Migrar para API Real

### 1. **Atualizar o Service**

```typescript
// Em portfolioService.ts
async getPortfolioData(): Promise<PortfolioData> {
  // Substituir por:
  return this.getPortfolioDataFromAPI();
}
```

### 2. **Configurar Endpoints**

```typescript
// Adicionar configuração de API
const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

async getPortfolioDataFromAPI(): Promise<PortfolioData> {
  const response = await fetch(`${API_BASE_URL}/api/portfolio`, {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('authToken')}`,
      'Content-Type': 'application/json'
    }
  });
  // ... resto da implementação
}
```

### 3. **Ajustar Tipos se Necessário**

- Atualizar interfaces em `portfolio.types.ts` conforme API real
- Adicionar novos campos ou modificar existentes

## Exemplo de Uso

### Dashboard

```typescript
function Dashboard() {
  const { summary, positions, isLoading, error, refetch } = usePortfolio();

  if (isLoading) return <LoadingSpinner />;
  if (error) return <ErrorComponent error={error} onRetry={refetch} />;

  return (
    <Layout>
      <PortfolioSummaryCards summary={summary} />
      <PositionsTable positions={positions} />
    </Layout>
  );
}
```

### Página de Posições

```typescript
function Positions() {
  const { positions, isLoading, error, refetch } = usePositions();

  return (
    <Layout>
      <PositionsTable positions={positions} title="Todas as Posições" />
    </Layout>
  );
}
```

## Próximos Passos

1. **Integração com API Real**: Substituir dados mock por chamadas reais
2. **Cache de Dados**: Implementar cache para melhorar performance
3. **Websockets**: Adicionar atualizações em tempo real
4. **Testes**: Criar testes unitários para hooks e componentes
5. **Otimizações**: Implementar lazy loading e memoização

## Estrutura de Arquivos

```
src/
├── types/
│   └── portfolio.types.ts
├── services/
│   └── portfolioService.ts
├── hooks/
│   └── usePortfolio.ts
├── components/
│   ├── PortfolioSummaryCards/
│   ├── PositionsTable/
│   ├── LoadingSpinner/
│   └── index.ts
└── pages/
    ├── Dashboard/
    └── Positions/
```

Esta arquitetura garante que o código seja escalável, maintível e pronto para integração com APIs reais.
