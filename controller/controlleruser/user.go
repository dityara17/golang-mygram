package controlleruser

import (
	"net/http"

	"github.com/arfan21/golang-mygram/helper"
	"github.com/arfan21/golang-mygram/service/serviceuser"
	"github.com/gin-gonic/gin"
)

type ControllerUser interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type controller struct {
	srv serviceuser.ServiceUser
}

func New(srv serviceuser.ServiceUser) ControllerUser {
	return &controller{srv}
}

func (c *controller) Create(ctx *gin.Context) {
	data := new(model.UserRequest)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, "BAD REQUEST", nil, err))
		return
	}

	response, err := c.srv.Create(data)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), "BAD REQUEST", nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusCreated, "CREATED", response, nil))
}

func (c *controller) GetByID(ctx *gin.Context) {

}
