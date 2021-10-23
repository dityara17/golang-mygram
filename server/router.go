package server

import (
	"github.com/arfan21/golang-mygram/controller/controlleruser"
	"github.com/arfan21/golang-mygram/repository/repositoryuser"
	"github.com/arfan21/golang-mygram/service/serviceuser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	r.POST("/users", ctrlUser.Create)
}
