package server

import (
	"github.com/arfan21/golang-mygram/config/configdb"
	"github.com/gin-gonic/gin"
)

func Start() error {
	db, err := configdb.New()
	if err != nil {
		return err
	}

	r := gin.Default()
	NewRouter(r, db)

	r.Use(gin.Recovery())

	r.Run()
	return nil
}
