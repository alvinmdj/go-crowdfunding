# Go Crowdfunding

## Links

- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/docs/)
- [GORM - Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)
- [Gin - Model Binding & Validation](https://gin-gonic.com/docs/examples/binding-and-validation/)

## Requirements

- Go v1.18+
- MySQL

## Installation

- Clone this repository:

```sh
git clone https://github.com/alvinmdj/go-crowdfunding.git
```

- Import ```go_crowdfunding.sql``` and setup db config:

```sh
# in main.go line 14
dsn := "<db_username>:<db_password>@tcp(127.0.0.1:3306)/<db_name>?charset=utf8mb4&parseTime=True&loc=Local"
```

- Go inside the directory:

```sh
cd go-crowdfunding
```

- Run:

```go
go run main.go
```

## Installed Package

- [Go Gin](https://github.com/gin-gonic/gin)

```go
go get -u github.com/gin-gonic/gin
```

- [GORM with MySQL Driver](https://gorm.io/docs/)

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```
