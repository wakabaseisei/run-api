package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wakabaseisei/runapp/interfaces/controllers"
	"github.com/wakabaseisei/runapp/interfaces/database"
	"github.com/wakabaseisei/runapp/usecase"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	userRepository := database.NewUserRepository(r.DB.Connection)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	r.Gin.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Home": "200 OK",
		})
	})

	r.Gin.GET("/users", func(c *gin.Context) { userController.GetAll(c) })
	r.Gin.GET("/users/:id", func(c *gin.Context) { userController.Get(c) })
	r.Gin.POST("/users", func(c *gin.Context) { userController.Post(c) })
	r.Gin.DELETE("/users/:id", func(c *gin.Context) { userController.Delete(c) })
	r.Gin.PUT("/users/:id", func(c *gin.Context) { userController.Update(c) })

	r.Gin.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health-check": "200 OK",
		})
	})
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
