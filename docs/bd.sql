CREATE TABLE usuario (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    login TEXT NOT NULL UNIQUE,
    senha TEXT NOT NULL,
    role VARCHAR(10) NOT NULL CHECK (role IN ('USER','ADMIN'))
);

CREATE TABLE produto (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    descricao TEXT NOT NULL,
    valor_unitario DECIMAL(10,2) NOT NULL,
    unidade_medida VARCHAR(10) NOT NULL,
    quantidade INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL REFERENCES usuario(id)
);