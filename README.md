
# CRUD - Go

Um simples CRUD em Go com GORM.


## Stack utilizada

**Back-end:** Go
**Banco de dados:** MySQL
**ORM:** GORM
**Container:** Docker


## Rodando localmente

Clone o projeto

```bash
  git clone https://github.com/kameikay/sample-crud-go-gorm-relations.git
```

Entre no diretório do projeto

```bash
  cd sample-crud-go-gorm-relations
```

Inicie o docker com docker-compose

```bash
  docker-compose up -d
```

Execute o bash do MySQL dentro do container

```bash
  docker-compose exec mysql bash
```

Dentro do container, acesse o banco de dados pelo bash

```bash
  mysql -uroot -p gormdb
```

Execute o main.go comentando os comandos que não for utilizar
```bash
  go run main.go
```