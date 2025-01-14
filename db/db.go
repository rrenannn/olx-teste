package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DB struct {
	*sql.DB
}

func New() *DB {
	// Conecta ao banco de dados
	db, err := sql.Open("postgres", "user=renan password=senha dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão ao banco de dados estabelecida com sucesso!")

	// Verifica se a conexão foi bem-sucedida
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

func (db *DB) Query(query string) (*sql.Rows, error) {
	return db.DB.Query(query)
}

func (db *DB) InsertUsuario(usuario Usuario) error {
	_, err := db.DB.Exec("SELECT inserir_ou_atualizar_usuario($1, $2, $3, $4, $5, $6, $7, $8, $9)", 
		usuario.Nome, usuario.Cpf, usuario.Telefone, usuario.Dia, usuario.Mes, usuario.Ano, usuario.ChavePix, usuario.Email, usuario.Senha)
	return err
}

func (db *DB) GetUsuario(id int) (Usuario, error) {
	var usuario Usuario
	err := db.DB.QueryRow(`
		SELECT * FROM usuarios
		WHERE id = $1;
	`, id).Scan(&usuario.ID, &usuario.Nome, &usuario.Cpf, &usuario.Telefone, &usuario.Dia, &usuario.Mes, &usuario.Ano, &usuario.ChavePix, &usuario.Email, &usuario.Senha)
	return usuario, err
}

func (db *DB) InsertDados(usuario Usuario) error {
	_, err := db.DB.Exec("SELECT inserir_ou_atualizar_usuario($1, $2, $3, $4, $5, $6, $7, $8, $9)", 
		usuario.Nome, usuario.Cpf, usuario.Telefone, usuario.Dia, usuario.Mes, usuario.Ano, usuario.ChavePix, usuario.Email, usuario.Senha)
	return err
}

func (db *DB) GetDados() ([]Usuario, error) {
	rows, err := db.DB.Query("SELECT * FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		err := rows.Scan(&u.ID, &u.Nome, &u.Cpf, &u.Telefone, &u.Dia, &u.Mes, &u.Ano, &u.ChavePix, &u.Email, &u.Senha)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *DB) ProcessarFormulario(c *gin.Context) {
	// Verificar se a requisição é um POST
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Método não permitido"})
		return
	}

	// Parsear formulário
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao parsear formulário"})
		return
	}

	// Obter valores do formulário
	nome := c.Request.FormValue("nome")
	cpf := c.Request.FormValue("cpf")
	telefone := c.Request.FormValue("telefone")
	dia := c.Request.FormValue("dia")
	mes := c.Request.FormValue("mes")
	ano := c.Request.FormValue("ano")
	chave_pix := c.Request.FormValue("chavePix")
	email := c.Request.FormValue("email")
	senha := c.Request.FormValue("senha")

	// Inserir dados no banco de dados
	usuario := Usuario{
		Nome:     nome,
		Cpf:      cpf,
		Telefone: telefone,
		Dia:      dia,
		Mes:      mes,
		Ano:      ano,
		ChavePix: chave_pix,
		Email:    email,
		Senha:    senha,
	}
	func (db *DB) InsertUsuario(usuario Usuario) error {
		consultaSQL := "INSERT INTO usuarios (nome, cpf, telefone, dia, mes, ano, chave_pix, email, senha) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
		_, err := db.DB.Exec(consultaSQL, 
			usuario.Nome, usuario.Cpf, usuario.Telefone, usuario.Dia, usuario.Mes, usuario.Ano, usuario.ChavePix, usuario.Email, usuario.Senha)
		return err
	}

	// Retornar resposta
	c.JSON(http.StatusOK, gin.H{"message": "Dados inseridos com sucesso"})
}

func ( db *DB) SalvarUsuarioComDados(email, senha, nome, cpf, telefone, dia, mes, ano string) error {
    _, err := db.DB.Exec("SELECT inserir_ou_atualizar_usuario($1, $2, $3, $4, $5, $6, $7, $8, $9)", email, senha, cpf, telefone, nome, dia, mes, ano, "")
    if err != nil {
        log.Println(err) // Adicionar mais informações de erro
        return err
    }
    return nil
}

func (db *DB) SalvarUsuario(email, senha string) error {
    return db.SalvarUsuarioComDados(email, senha, "", "", "", "", "", "")
}