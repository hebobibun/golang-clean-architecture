# CleanArch

Clean Architecture with Golang
# Features
## User : 
- Register anew user
- Login user
- Show profile
- Update profile
- Deactivate account
## Book :
- Insert a new book
- Show all books in the system
- Show all my books
- Show one of my book
- Update book information
- Delete a book
# API Documentation

[Click here](https://documenter.getpostman.com/view/23707537/2s8Z75SUzT) to see the documentation.

# Tools & Requirements

- Go 1.19.3
- Gorm & MySQL
- Echo v4
- JWT

## Folder Structure Pattern
```
├── config
│   └── config.go
│   └── db.go
└── features
│   └── book
│   │   └── data
│   │   │   └── model.go
│   │   │   └── query.go
│   │   └── handler
│   │   │   └── handler.go
│   │   │   └── request.go
│   │   │   └── response.go
│   │   └── services
│   │   │   └── service_test.go
│   │   │   └── service.go
│   │   └── entity.go
└── helper
│   └── jwt.go
│   └── response.go
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── local.env.example
├── main.go
├── README.md
```

## How to Install

- Clone it

```
$ git clone https://github.dev/ALTA-BE14-Habib/CleanArch
```

- Go to directory

```
$ cd CleanArch
```

- Create a new database

- Rename `local.env.example` to `local.env`
- Adjust `local.env` as your environment settings

- Run the project

```
$ go run .
```
