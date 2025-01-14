-- sql/queries/usuarios.sql
-- name: CriarUsuario :exec
INSERT INTO usuarios (nome, cpf, telefone, dia, mes, ano, chave_pix, email, senha)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);