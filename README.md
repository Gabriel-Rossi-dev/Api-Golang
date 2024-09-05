API EM GOLANG 

Tecnologias Utilizadas:

* [Viper](http://github.com/spf13/viper)  (go get github.com/spf13/viper)
* [Postegres](http://github.com/lib/pq)   (go get github.com/lib/pq) 
* [Docker](https://www.docker.com/)

<h2>Comandos Docker :<h2>
    <h3>*  docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5 <h3>
    <h3>*  docker exec -it api-todo psql -U postgres <h3>
    <h3>*  create database api_todo; <h3>
    <h3>*  create user user_todo; <h3>
    <h3>*  alter user user_todo with encrypted password '1122'; <h3>
    <h3>*  grant all privileges on database api_todo to user_todo; <h3>
    <h3>*  \l <h3>
    <h3>*  \c api_todo <h3>
    <h3>*  create table todos (id serial primary key, name varchar, description text, done bool); <h3>
    <h3>*  grant all privileges on all tables in schema public to user_todo; <h3>
    

