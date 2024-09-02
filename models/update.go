package models

import "github.com/Gabriel-Rossi-dev/API--POSTGRESQL/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title =$2, descriptin =$3
	done =$4 WHERE id =$1`, id, todo.Title, todo.Description, todo.Done)

	if err != nil {
		return 0, err
	}

	//A Função retorna a quantidade de linhas e o erro
	return res.RowsAffected()

}
