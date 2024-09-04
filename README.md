API EM GOLANG 

Tecnologias Utilizadas:

* [Viper](http://github.com/spf13/viper)  (go get github.com/spf13/viper)
* [Postegres](http://github.com/lib/pq)   (go get github.com/lib/pq) 
* [Docker](https://www.docker.com/)

<h2>Comandos Docker :<h2>
    *  docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
    *  docker exec -it api-todo psql -U postgres
    *  create database api_todo;
    *  create user user_todo;
    *  alter user user_todo with encrypted password '1122';
    *  grant all privileges on database api_todo to user_todo;
    *  \l
    *  \c api_todo
    *  create table todos (id serial primary key, name varchar, description text, done bool);
    *  grant all privileges on all tables in schema public to user_todo;
    

