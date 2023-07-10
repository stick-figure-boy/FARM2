package route

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	_ "github.com/hiroki-Fukumoto/farm2/docs"
	"github.com/hiroki-Fukumoto/farm2/domain/user"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Health Check
// @Tags healthCheck
// @Accept json
// @Produce json
// @Success 200 {struct { Message string }{ Message: "Health Check OK" }}
// @Router /v1/health-check [get]
func SetupRouter(db *sqlx.DB) *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:8081",
		},
		AllowMethods: []string{
			"*",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	appApiV1 := route.Group("/v1")

	gHealthCheck := appApiV1.Group("health-check")
	{
		gHealthCheck.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Health Check OK",
			})
		})
	}

	gUser := appApiV1.Group("users")
	{
		r := user.NewUserRepository(db)
		s := user.NewUserService(r)
		c := user.NewUserController(s)

		gUser.POST("/", c.RegisterUser)
		gUser.GET("/:accountID", c.FindUser)
	}

	return route
}
