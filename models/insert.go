package models

import "github.com/Gabriel-Rossi-dev/API--POSTGRESQL/db"

func Insert(todo Todo) (id int64, err error) {
	//Abre conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	//fecha conexão com o banco de dados
	defer conn.Close()

	//comando sql para inserção esperando os parametros
	sql := `INSERT INTO todos (title, description, done) 
	VALUES ($1, $2, $3)	RETURNING id`

	//passando os parametros para a qry e alimentando as variáveis err, id
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
