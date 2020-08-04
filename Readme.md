# Swagger Gin Demo
![image](https://github.com/kimi0230/swagger_gin_demo/blob/master/screenshot/demo.png)

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

---
## Folder structure
    ├── Readme.md
    ├── app
    │   ├── controllers
    │   │   └── api
    │   │       ├── test
    │   │       │   └── test.go
    │   │       └── v1
    │   │           └── account
    │   │               └── account.go
    │   ├── models
    │   │   ├── account.go
    │   │   └── error.go
    │   └── services
    │       └── httputil
    │           └── error.go
    ├── docs
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── main.go
    ├── routes
    │   └── routes.go
    └── tmp
        └── runner-build
---

## API annotations
[swaggo/swag README](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)
 
    title	        簡單API專案的標題或主要的業務功能
    version	        目前這專案API的版本
    description	簡單描述
    tersOfService	服務條款
    tag.name	標籤名稱
    tag.description	標籤描述
    tag.docs.url	標籤的外部文檔URL
    tag.docs.description	標籤的外部文檔說明
    contact.name	作者名稱
    contact.url	作者網址
    contact.email	作者email
    license.name	許可證名稱
    license.url	許可證網址
    host	        服務名稱或者是ip
    BasePath	基本URL路徑, (/api/v1)
    schemes	        提供的協定, (http, https)

## API Operation
    description	操作行為的詳細說明
    summary	描述該API
    id	用於標示API的唯一字符,在所有API操作中必須唯一。
    tags	歸屬同一類的API的tag, 以逗號分隔
    accept	request的context-type
    produce	response的context-type
    param	參數按照: `參數名` `參數類型` `參數的資料類型` `是否必填` `註解`
    header	response header: `return code` `參數類型` `資料類型` `註解`
    router	path httpMethod
    
![image](https://github.com/kimi0230/swagger_gin_demo/blob/master/screenshot/APIOperation.png)


### Param Type
https://github.com/swaggo/swag#param-type
   
    query
    path
    header
    body
    formData
    
Data Type
https://github.com/swaggo/swag#param-type

    string (string)
    integer (int, uint, uint32, uint64)
    number (float32)
    boolean (bool)
    user defined struct

Mime Types
https://github.com/swaggo/swag#mime-types

Attribute
https://github.com/swaggo/swag#attribute

## API Documentation link
    http://localhost:8080/swagger/index.html

## Reference
* [swaggo/swag](https://github.com/swaggo/swag)