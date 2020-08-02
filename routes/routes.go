package routes

import (
	"swagger_gin_demo/app/controllers/api/v1/account"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter : Create gin router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	// 測試用 controller
	v1 := r.Group("/api/v1")
	{
		v1.GET("/kimi", func(c *gin.Context) {
			result := map[string]interface{}{
				"Data": map[string]interface{}{
					// "AppName": os.Getenv("APP_NAME"),
					// "AppENV":  os.Getenv("APP_ENV"),
					"Msg":     "Hello Kimi!",
					"Time":    time.Now().Format("2006-01-02 15:04:05"),
					"TimeUTC": time.Now().UTC().Format("2006-01-02 15:04:05"),
				},
				"Result": map[string]interface{}{
					"Status":  true,
					"Code":    "2000",
					"Message": "正確狀況",
				},
			}
			c.JSON(http.StatusOK, result)
		})

		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", account.ShowAccount)
			accounts.GET("", account.ListAccounts)
			accounts.POST("", account.AddAccount)
			accounts.PATCH(":id", account.UpdateAccount)
			accounts.DELETE(":id", account.DeleteAccount)
			accounts.POST(":id/images", account.UploadAccountImage)
		}

	}

	return r
}
