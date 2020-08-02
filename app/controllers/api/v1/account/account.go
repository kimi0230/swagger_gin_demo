package account

import (
	"fmt"
	"net/http"
	"strconv"

	"swagger_gin_demo/app/models"
	"swagger_gin_demo/app/services/httputil"

	"github.com/gin-gonic/gin"
)

// 註解	描述
// summary	描述該API
// tags	歸屬同一類的API的tag
// accept	request的context-type
// produce	response的context-type
// param	參數按照: `參數名` `參數類型` `參數的資料類型` `是否必填` `註解`
// header	response header: `return code` `參數類型` `資料類型` `註解`
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
	account, err := models.AccountOne(aid)
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
func ListAccounts(c *gin.Context) {
	q := c.Request.URL.Query().Get("name")
	accounts, err := models.AccountsAll(q)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// AddAccount godoc
// @Summary Add an account
// @Description add by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body models.AddAccount true "Add account"
// @Success 200 {object} models.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [post]
func AddAccount(c *gin.Context) {
	var addAccount models.AddAccount
	if err := c.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := addAccount.Validation(); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	account := models.Account{
		Name: addAccount.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	account.ID = lastID
	c.JSON(http.StatusOK, account)
}

// UpdateAccount godoc
// @Summary Update an account
// @Description Update by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Param  account body models.UpdateAccount true "Update account"
// @Success 200 {object} models.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [patch]
func UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	var updateAccount models.UpdateAccount
	if err := c.ShouldBindJSON(&updateAccount); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	account := models.Account{
		ID:   aid,
		Name: updateAccount.Name,
	}
	err = account.Update()
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, account)
}

// DeleteAccount godoc
// @Summary Delete an account
// @Description Delete by account ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID" Format(int64)
// @Success 204 {object} models.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [delete]
func DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	err = models.Delete(aid)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// UploadAccountImage godoc
// @Summary Upload account image
// @Description Upload file
// @Tags accounts
// @Accept  multipart/form-data
// @Produce  json
// @Param  id path int true "Account ID"
// @Param file formData file true "account image"
// @Success 200 {object} Message
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id}/images [post]
func UploadAccountImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	type Message struct {
		Message string `json:"message" example:"message"`
	}
	c.JSON(http.StatusOK, Message{Message: fmt.Sprintf("upload complete userID=%d filename=%s", id, file.Filename)})
}
