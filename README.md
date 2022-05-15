# Go Crowdfunding

## Links

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [Gin - Model Binding & Validation](https://gin-gonic.com/docs/examples/binding-and-validation/)
- [GORM](https://gorm.io/docs/)
- [GORM - Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)
- [GORM - Preloading](https://gorm.io/docs/preload.html)
- [GoDotEnv](https://github.com/joho/godotenv)
- [GoSimple - Slug](https://github.com/gosimple/slug)
- [Golang-JWT](https://github.com/golang-jwt/jwt)
- [JWT](https://jwt.io/)

## Requirements

- [Go v1.18+](https://go.dev/)
- [MySQL](https://www.mysql.com/)

## Installation

- Clone this repository:

```sh
git clone https://github.com/alvinmdj/go-crowdfunding.git
```

- Go inside the directory:

```sh
cd go-crowdfunding
```

- Get all required dependencies

```go
go get .
```

- Copy ```.env.example``` to ```.env``` and setup variables in ```.env```:

```sh
cp .env.example .env
```

- Run:

```go
go run main.go
// or 
go run .
```

## Installed Package

- [Gin](https://github.com/gin-gonic/gin)

```go
go get -u github.com/gin-gonic/gin
```

- [GORM with MySQL Driver](https://gorm.io/docs/)

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

- [GoDotEnv](https://github.com/joho/godotenv)

```go
go get github.com/joho/godotenv
```

- [Golang-JWT](https://github.com/golang-jwt/jwt)

```go
go get -u github.com/golang-jwt/jwt
```

- [GoSimple Slug](https://github.com/gosimple/slug)

```go
go get -u github.com/gosimple/slug
```
