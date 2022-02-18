# Swagger for GIN framework
Converts Go annotations to Swagger Documentation

#### Install

```sh
go get -u github.com/swaggo/swag/cmd/swag

# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest
```

#### Configuration

```sh
swag init -g cmd/main.go
```

#### Add anotations