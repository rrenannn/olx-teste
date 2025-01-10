-- sql/schema/schema.sql
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT NOT NULL,
    cpf TEXT NOT NULL,
    telefone TEXT NOT NULL,
    dia TEXT NOT NULL,
    mes TEXT NOT NULL,
    ano TEXT NOT NULL,
    chave_pix TEXT NOT NULL
);