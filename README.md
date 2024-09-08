
Tecnologias Utilizadas:

* [Viper](http://github.com/spf13/viper)  ```(go get github.com/spf13/viper)```
* [Postegres](http://github.com/lib/pq)  ``` (go get github.com/lib/pq) ```
* [Chi]()
* [Docker](https://www.docker.com/)

<h2>Comandos Docker :<h2>

Baixando imagem do Postgres via Docker, passando a porta padrão como 5432 e dando o nome api-todo para a imagem docker.
``` console
docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
```
Conectando ao BD
``` console
docker exec -it api-todo psql -U postgres
```
Criando Base de Dados
``` console
create database api_todo;
```
Conectando a Base de Dados
``` console
\c api_todo; 
```
Criando Usuário do BD
``` console
create user user_todo; 
```
Definindo senha para o Usuário
``` console
alter user user_todo with encrypted password '1122';
```
Dando as Permissões da base api_todo para o usuário
``` console
grant all privileges on database api_todo to user_todo;
```
Criando tabela de Todos.
``` console
create table todos (id serial primary key, title varchar, description text, done bool);
```
Dando as permissões de todoas as tabelas da base para o usuário
``` console
grant all privileges on all tables in schema public to user_todo;
```
Dando as permissões de sequences para o Usuário
``` console
 grant all privileges on all sequences in schema public to user_todo;
```
