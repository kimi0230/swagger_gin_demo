# Swagger Gin Demo

## Install Gin
    go get -u github.com/gin-gonic/gin

## Install package
    govendor sync

## Swagger

#### Install swagger
    go get -u github.com/swaggo/swag/cmd/swag  

#### Initial swagger document
    swag init
---

## Start server
 Defaut is "DebugMode".
 
#### Run ReleaseMode
    go run main.go app
#### Run DebugMode
    go run main.go dev
#### Run TestMode
    go run main.go qa

## API Annotations
[README](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)
 
    title	        簡單API專案的標題或主要的業務功能
    version	        目前這專案/API的版本
    description	簡單描述
    tersOfService	服務條款
    contact.name	作者名稱
    contact.url	作者網址
    contact.email	作者email
    license.name	許可證名稱
    license.url	許可證網址
    host	        服務名稱或者是ip
    BasePath	基本URL路徑, (/api/v1)
    schemes	        提供的協定, (http, https)

## API Operation
    summary	描述該API
    tags	歸屬同一類的API的tag
    accept	request的context-type
    produce	response的context-type
    param	參數按照: `參數名` `參數類型` `參數的資料類型` `是否必填` `註解`
    header	response header: `return code` `參數類型` `資料類型` `註解`
    router	path httpMethod
    
## Reference
* [swaggo/swag](https://github.com/swaggo/swag)
