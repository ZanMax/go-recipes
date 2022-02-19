# Swagger for GIN framework
Converts Go annotations to Swagger Documentation
[Swag on GitHub](https://pandao.github.io/editor.md/en.html)

#### Install

```sh
go get -u github.com/swaggo/swag/cmd/swag

# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest
```

#### Configuration

##### Import
```sh
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files" // swagger embed files
```
##### Add
```sh
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

#### Add anotations

#### Generation
```sh
swag init -g cmd/main.go
```
