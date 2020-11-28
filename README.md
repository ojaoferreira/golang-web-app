# golang-web-app

#### 1. Pré-requisitos

1.1 Instância do banco MySQL
```
mysql> create database **{name-banco}**
mysql> use **{name-banco}**
mysql> create table posts (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(50) not null, body text);
```

1.2 Variáveis de ambiente

| NAME        | Requerida | Padrão |
|-------------|-----------|--------|
| DB_HOST     | Sim       |        |
| DB_PORT     | Sim       |        |
| DB_NAME     | Sim       |        |
| DB_USER     | Sim       |        |
| DB_PASS     | Sim       |        |
| DB_CHARSET  | Não       | utf8   |
| APP_PORT    | Não       | 8080   |

#### 2. Buildando e executando à aplicação

$ git clone https://github.com/ojaoferreira/golang-web-app.git

$ cd golang-web-app

$ go get

$ go build -o golang-web-bin

$ ./golang-web-bin


