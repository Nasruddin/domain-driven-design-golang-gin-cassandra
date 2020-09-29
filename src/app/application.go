package app

import (
	"github.com/gin-gonic/gin"
	"golang-gin-cassandra/src/api/http"
	"golang-gin-cassandra/src/domain/users/service"
	"golang-gin-cassandra/src/repository/database"
)

var (
	router = gin.Default()
)

func StartApp()  {
	userHandler := http.NewHandler(service.NewService(database.NewDbRepository()))

	router.GET("/users/:user_id", userHandler.GetById)
	router.POST("/user", userHandler.Create)

	_ = router.Run(":8888")
}
