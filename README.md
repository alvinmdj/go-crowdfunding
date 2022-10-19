# Go Crowdfunding

Back-end with Go Gin web framework, consists of admin pages and API for front-end needs.

Admin can performs:

- Login & logout (session based).
- View users, campaigns, and transactions.
- Create & edit users data & campaign data.
- Upload user & campaign image.

API: [API Documentation (Postman)](https://documenter.getpostman.com/view/16534190/Uz5DrdeP)

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
- [Midtrans](https://midtrans.com/)
- [Go Midtrans](https://github.com/veritrans/go-midtrans)
- [Midtrans Docs Snap API](https://snap-docs.midtrans.com/)
- [Midtrans Docs Handle After Payment HTTP(S) Notification](https://docs.midtrans.com/en/after-payment/http-notification)
- [ngrok](https://ngrok.com/)
- [Gin - CORS](https://github.com/gin-contrib/cors)
- [Gin - Multitemplate](https://github.com/gin-contrib/multitemplate)
- [Bootadmin Template (Archived)](https://web.archive.org/web/20201129084141/https://github.com/kjdion84/bootadmin)
- [Bootadmin Docs (Archived)](https://web.archive.org/web/20210301183117/https://bootadmin.net/demo/docs)
- [accounting - money and currency formatting for golang](https://github.com/leekchan/accounting)
- [Gin - Sessions](https://github.com/gin-contrib/sessions)

## Requirements

- [Go v1.18+](https://go.dev/)
- [MySQL](https://www.mysql.com/)
- [Midtrans](https://midtrans.com/)

## Installation

- Clone this repository:

```sh
git clone https://github.com/alvinmdj/go-crowdfunding.git
```

- Go inside the directory:

```sh
cd go-crowdfunding
```

- Get all dependencies:

```go
go get .
```

- Create MySQL database by importing ```go_crowdfunding.sql```:

```go
// Import go_crowdfunding.sql from /config/db/go_crowdfunding.sql
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

## Installed Packages

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

- [Go Midtrans](https://github.com/veritrans/go-midtrans)

```go
go get -u github.com/veritrans/go-midtrans
```

- [Gin CORS](https://github.com/gin-contrib/cors)

```go
go get github.com/gin-contrib/cors
```

- [Gin - Multitemplate](https://github.com/gin-contrib/multitemplate)

```go
go get github.com/gin-contrib/multitemplate
```

- [accounting - money and currency formatting for golang](https://github.com/leekchan/accounting)

```go
go get github.com/leekchan/accounting
```

- [Gin - Sessions](https://github.com/gin-contrib/sessions)

```go
go get github.com/gin-contrib/sessions
```

## ngrok Commands

- Help

```sh
ngrok help
```

- Run

```sh
ngrok http <port>

# example
ngrok http 8080

# https://docs.midtrans.com/en/after-payment/http-notification
# Setup in Midtrans Configuration - Payment Notification URL:
# <ngrok-url>/api/v1/transactions/notification
```

- Required Environment Variables on Heroku:

```sh
APP_ENV=production
GOVERSION=go1.18
DB_HOST=
DB_NAME=
DB_USERNAME=
DB_PASSWORD=
JWT_SECRET=
MIDTRANS_CLIENT_KEY=
MIDTRANS_SERVER_KEY=

# Also use 'heroku/go' buildpack (from heroku project dashboard > settings > buildpacks > add buildpack > choose go)

# Push changes to Heroku with: git push heroku main
```
