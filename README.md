API EM GOLANG 

Tecnologias Utilizadas:

* [Viper](http://github.com/spf13/viper)  ```(go get github.com/spf13/viper)```
* [Postegres](http://github.com/lib/pq)  ``` (go get github.com/lib/pq) ```
* [Docker](https://www.docker.com/)

<h2>Comandos Docker :<h2>
    
``` console
docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
```
``` console
docker exec -it api-todo psql -U postgres
```
``` console
create database api_todo;
```
``` console
create user user_todo; 
```
``` console
alter user user_todo with encrypted password '1122';
```
``` console
grant all privileges on database api_todo to user_todo;
```
<h2>mostra as tabelas <H2>
    
``` console
\dt
```
``` console
create table todos (id serial primary key, title varchar, description text, done bool);
```
``` console
grant all privileges on all tables in schema public to user_todo;
```
``` console
 grant all privileges on all sequences in schema public to user_todo;
```

