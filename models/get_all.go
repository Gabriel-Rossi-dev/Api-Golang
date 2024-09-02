package models

import "github.com/Gabriel-Rossi-dev/API--POSTGRESQL/db"

func GetAll() (todos []Todo, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	//Busca todas as informações do banco de dados
	rows, err := conn.Query(`SELECT * FROM todos `)
	if err != nil {
		return
	}
	for rows.Next() {
		var todo Todo
		//Scanea os ponteiros de cada coluna
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}
