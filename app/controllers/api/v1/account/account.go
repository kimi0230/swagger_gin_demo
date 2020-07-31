package account

import (
	"net/http"
	"strconv"

	"swagger_gin_demo/app/services/httputil"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/model"
)

// 註解	描述
// summary	描述該API
// tags	歸屬同一類的API的tag
// accept	request的context-type
// produce	response的context-type
// param	參數按照: 參數名 參數類型 參數的資料類型 是否必須 註解 (中間都要空一格)
// header	response header return code 參數類型 資料類型 註解
// router	path httpMethod

// ShowAccount godoc
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "Account ID"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} models.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [get]
func ShowAccount(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	account, err := model.AccountOne(aid)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, account)
}

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param name query string false "name search" Format(email)
// @Success 200 {array} models.Account
// @Router /accounts [get]
func ListAccounts(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("name")
	accounts, err := model.AccountsAll(q)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
