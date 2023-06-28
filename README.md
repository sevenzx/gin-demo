# gin-demo

Includes quick learning of gin, gorm, gorm-gen, and using them to build a demo of a todo list

## Getting Started

```bash
# run project
go run server/main.go
```
```bash
# build project
go build -o bubble server/main.go
```

## Structure

```tree
├── server
│   ├── config
│   │   ├── config.go
│   │   └── config.yml
│   ├── core
│   │   ├── controller
│   │   ├── model
│   │   ├── query
│   │   └── router
│   ├── db
│   │   ├── gen
│   │   └── mysql
│   └── main.go
├── study
│   ├── gen
│   ├── gin
│   └── gorm
└── web
    ├── static
    └── templates

```
