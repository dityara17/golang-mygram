package server

import "github.com/gin-gonic/gin"

func Start() error {
	r := gin.Default()
	NewRouter(r, nil)
	return nil
}
