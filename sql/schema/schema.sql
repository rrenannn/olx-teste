-- sql/schema/schema.sql
CREATE OR REPLACE FUNCTION inserir_ou_atualizar_usuario(
    p_email VARCHAR(255),
    p_senha VARCHAR(255),
    p_cpf VARCHAR(14),
    p_telefone VARCHAR(15),
    p_nome VARCHAR(50),
    p_dia VARCHAR(2),
    p_mes VARCHAR(2),
    p_ano VARCHAR(4),
    p_chave_pix VARCHAR(50)
)
RETURNS VOID AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM usuarios WHERE cpf = p_cpf OR telefone = p_telefone) THEN
        UPDATE usuarios
        SET email = p_email,
            senha = p_senha,
            nome = p_nome,
            dia = p_dia,
            mes = p_mes,
            ano = p_ano,
            chave_pix = p_chave_pix
        WHERE cpf = p_cpf OR telefone = p_telefone;
    ELSE
        INSERT INTO usuarios (email, senha, cpf, telefone, nome, dia, mes, ano, chave_pix)
        VALUES (p_email, p_senha, p_cpf, p_telefone, p_nome, p_dia, p_mes, p_ano, p_chave_pix);
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    cpf TEXT NOT NULL,
    telefone TEXT NOT NULL,
    dia TEXT NOT NULL,
    mes TEXT NOT NULL,
    ano TEXT NOT NULL,
    chave_pix TEXT NOT NULL,
    email TEXT NOT NULL,
    senha TEXT NOT NULL
);

DELETE FROM usuarios
WHERE id IN (
  SELECT id
  FROM (
    SELECT id, ROW_NUMBER() OVER (PARTITION BY cpf, telefone ORDER BY id) AS row_num
    FROM usuarios
  ) AS subquery
  WHERE row_num > 1
);