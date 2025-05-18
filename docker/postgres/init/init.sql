-- Criar tabela de usuários
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Criar o enum para tipo de ativo
CREATE TYPE asset_type AS ENUM ('CRYPTO', 'STOCK', 'ETF');

-- Criar o enum para tipo de operação
CREATE TYPE operation_type AS ENUM ('BUY', 'SELL');

-- Criar tabela de transações
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    asset_symbol VARCHAR(50) NOT NULL,
    asset_type asset_type NOT NULL,
    quantity DECIMAL(18, 8) NOT NULL,
    price_in_usd DECIMAL(18, 8) NOT NULL,
    usd_brl_rate DECIMAL(10, 4) NOT NULL,
    price_in_brl DECIMAL(18, 2) NOT NULL,
    type operation_type NOT NULL,
    operation_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Criar índices para melhorar a performance de consultas comuns
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_asset_symbol ON transactions(asset_symbol);
CREATE INDEX idx_transactions_operation_date ON transactions(operation_date); 