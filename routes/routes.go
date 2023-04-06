package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"test-case/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/transactions", controllers.GetAllTransaction)
	r.POST("/save-transactions", controllers.SaveTransaction)
	r.GET("/dummy-transactions/:dummy", controllers.GenerateDummy)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
