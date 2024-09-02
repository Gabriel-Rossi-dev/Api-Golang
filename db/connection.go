package db

import (
	"database/sql"
	"fmt"

	"github.com/Gabriel-Rossi-dev/API--POSTGRESQL/configs"

	//driver para Postgres
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.DataBase)

	//A função sql.open Espera 2 parametros
	// - Driver Name
	// - DataSource Name
	//Essa função retorna dois valores um ponteiro de sql.DB
	// e um erro
	//Por esse motivo passa duas variáveis para receber os retornos
	conn, err := sql.Open("postgres", sc)
	//se o retorno do sql.open for diferente de vazio
	//Então nao foi possível conectar ao banco , então panic
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
