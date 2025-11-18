# 🐹 API de Produtos em Go — GORM + MySQL + Docker

---

## ✨ Visão Geral

Este projeto é uma **API REST** desenvolvida em **Go (Golang)** utilizando o **GORM** como ORM e **MySQL** como banco de dados.  
A aplicação possui endpoints para **criação, listagem e gerenciamento de produtos**, com arquitetura limpa e preparada para evolução.

---

## 🚀 Funcionalidades

- ✔️ CRUD de produtos (`POST`, `GET`, `PUT`, `DELETE`)
- ✔️ Persistência com **GORM**
- ✔️ Banco MySQL rodando via **Docker Compose**
- ✔️ Estrutura organizada (controllers, services, entities)
- ✔️ Tratamento de erros e respostas padronizadas
- ✔️ Migrations automáticas via GORM

---
- **Controller**: recebe requisições HTTP  
- **UseCase / Service**: regra de negócio  
- **Repository**: acesso ao banco via GORM  
- **Database**: conexão com MySQL via Docker  

---

## 🧰 Requisitos

- **Go 1.22+**
- **Docker + Docker Compose**
- **MySQL 8+** (se rodar localmente)
- Porta **8000** liberada

---
## 📦 Estrutura do Projeto

```bash
/DoceSim
├── cmd
│   └── main.go
├── db
│   └── conn.go
├── product
│   ├── controller.go
│   ├── usecase.go
│   └── entity.go
├── docker-compose.yml
├── go.mod
└── README.md
```

## 🐬 Docker Compose (MySQL + API)

```yaml
version: '3.8'

services:
  go_db:
    image: mysql:8.0
    container_name: docesim
    environment:
      MYSQL_ROOT_PASSWORD: {MYSQL_ROOT_PASSWORD}
      MYSQL_DATABASE: {MYSQL_DATABASE}
      MYSQL_USER: {MYSQL_USER}
      MYSQL_PASSWORD: {MYSQL_PASSWORD}
    ports:
      - "3306:3306"
    volumes:
      - data_db:/var/lib/mysql

volumes:
  data_db:{}
```
## ⚙️ Configuração do Banco

### `db/conn.go`

```go
dsn := "user:user123@tcp(mysql:3306)/docesim?charset=utf8mb4&parseTime=True&loc=Local"

db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
    panic("Erro ao conectar no banco: " + err.Error())
}
```
## ▶️ Como rodar o projeto

---

### 1️⃣ Clonar o repositório

```bash
git clone https://github.com/luger-mano/DoceSim.git
cd seu-projeto
```
## 2️⃣ Subir MySQL + API

```bash
docker compose up -d
```
## 🌐 Inicialização da API

A API estará acessível assim que os serviços subirem:

👉 **http://localhost:8000**

------------------------------------------------------------------------

## 🐬 3️⃣ Verificar logs do MySQL

``` bash
docker logs docesim
```

------------------------------------------------------------------------

## 🧪 Testando a API

### 📌 Criar um Produto --- **POST**

**Endpoint:**

    POST http://localhost:8000/products

**Body Exemplo:**

``` json
{
  "name": "Produto Teste",
  "description": "Produto Teste",
  "value": 49.90
}
```

------------------------------------------------------------------------

### 📌 Listar Produtos --- **GET**

    GET http://localhost:8000/products

**Exemplo de Resposta:**

``` json
[
  {
    "id": 1,
    "name": "Produto Teste",
    "price": 49.9,
    "createdAt": "2025-11-18T02:51:00Z"
  }
]
```

------------------------------------------------------------------------

## 🌐 Endpoints Disponíveis

### ▶️ **POST /products**

Cria um novo produto.

### ▶️ **GET /products**

Lista todos os produtos.

### ▶️ **PUT /products/:id**

Atualiza um produto existente.

### ▶️ **DELETE /products/:id**

Exclui um produto.

------------------------------------------------------------------------

## 🧩 Possíveis Melhorias Futuras

-   🔐 Autenticação **JWT**
-   ✔️ Validação com **go-playground/validator**
-   📄 Paginação nas listagens
-   ⚡ Implementação de **cache** com Redis
-   🛠️ Versionamento da API (**v1**, **v2**...)

------------------------------------------------------------------------

## 🙌 Obrigado por ver até aqui!

Sinta-se livre para sugerir melhorias, abrir issues ou pedir novas
funcionalidades! 🚀



