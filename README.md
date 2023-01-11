# Golang Clean Architecture

In this Golang REST API project, I am trying to follow clean architecture principles. It is designed to be modular, scalable, and easy to maintain. The code is organized into logical layers, with each layer having a specific responsibility. This separation of concerns makes it easier to understand and modify the code, as well as perform testing and debugging.
# Features
## User : 
- Register a new user
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

- Go 1.19.x
- Viper 
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
│   └── user
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
└── README.md
```

## How to Install

- Clone it

```
$ git clone https://github.com/hebobibun/golang-clean-architecture
```

- Go to directory

```
$ cd golang-clean-architecture
```

- Create a new database

- Rename `local.env.example` to `local.env`
- Adjust `local.env` as your environment settings

- Run the project

```
$ go run .
```

## Enjoy

Keep learning! ^^
