package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String para conexão com o Postgres
	DatabaseConnectionString = ""
	//Porta onde a API vai estar rodando
	Port = 0
	//Chave que valida o token
	SecretKey []byte
)

//Carregar vai inicializar as variáveis de ambiente.
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	DatabaseConnectionString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("POSTGRES_PORT"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
