package server

import (
	"github.com/arfan21/golang-mygram/controller/controlleruser"
	_ "github.com/arfan21/golang-mygram/docs"
	"github.com/arfan21/golang-mygram/middleware"
	"github.com/arfan21/golang-mygram/repository/repositoryuser"
	"github.com/arfan21/golang-mygram/service/serviceuser"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("", middleware.Authorization, ctrlUser.Update)

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
