# Swagger Gin Demo

## Install package
govendor sync

## Install swagger
    go get -u github.com/swaggo/swag/cmd/swag  

## Initial swagger
    swag init

## Start server (Defaut is "DebugMode")
#### Run ReleaseMode
    go run main.go app
#### Run DebugMode
    go run main.go dev
#### Run TestMode
    go run main.go qa

## Reference
* [swaggo/swag](https://github.com/swaggo/swag)