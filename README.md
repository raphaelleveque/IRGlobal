# IRGlobal - Sistema para Declaração de Investimentos no Exterior

Sistema web para auxiliar brasileiros na declaração do Imposto de Renda sobre investimentos no exterior, especialmente em criptomoedas, ações e ETFs.

## Requisitos

- Docker
- Docker Compose

## Configuração e Execução

1. Clone o repositório:

```
git clone [URL_DO_REPOSITORIO]
cd IRGlobal
```

2. Inicie os containers Docker:

```
docker-compose up -d
```

3. Acesse o pgAdmin para gerenciar o banco de dados (opcional):

   - URL: http://localhost:5050
   - Email: admin@irglobal.com
   - Senha: admin123

   Para conectar ao servidor PostgreSQL no pgAdmin:

   - Host: postgres
   - Porta: 5432
   - Usuário: irglobal
   - Senha: irglobal123
   - Banco de dados: irglobal

## Estrutura do Banco de Dados

### Tabela `users`

Armazena os dados do usuário autenticado.

| Campo      | Tipo          | Descrição                 |
| ---------- | ------------- | ------------------------- |
| id         | UUID          | Identificador único       |
| email      | string        | E-mail do usuário (único) |
| password   | string (hash) | Senha (hash)              |
| created_at | timestamp     | Data de criação           |

### Tabela `transactions`

Registra cada compra ou venda de ativo.

| Campo          | Tipo      | Descrição                           |
| -------------- | --------- | ----------------------------------- |
| id             | UUID      | Identificador único                 |
| user_id        | UUID      | FK para users                       |
| asset_symbol   | string    | Ex: BTC, ETH, AAPL, VOO             |
| tipo_ativo     | enum      | CRYPTO, STOCK, ETF                  |
| quantity       | float     | Quantidade comprada ou vendida      |
| price_in_usd   | float     | Preço unitário em dólar na data     |
| usd_brl_rate   | float     | Cotação do dólar no dia da operação |
| price_in_brl   | float     | Valor total em reais                |
| type           | enum      | BUY ou SELL                         |
| operation_date | date      | Data da operação                    |
| created_at     | timestamp | Data de inserção do registro        |
