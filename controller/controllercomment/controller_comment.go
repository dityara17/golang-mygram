package controllercomment

import (
	"github.com/arfan21/golang-mygram/helper"
	"github.com/arfan21/golang-mygram/model/modelcomment"
	"github.com/arfan21/golang-mygram/service/servicecomment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ControllerComment interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	srv servicecomment.ServiceComment
}

func (c *controller) Create(ctx *gin.Context) {
	request := new(modelcomment.Request)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	userID := ctx.MustGet("user_id")
	request.UserID = userID.(uint)

	// service
	create, err := c.srv.Create(*request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, create, nil))
	return
}

func (c *controller) Get(ctx *gin.Context) {
	responses, err := c.srv.Get()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, responses, nil))
}

func (c *controller) Update(ctx *gin.Context) {
	paramKeyID := ctx.Param("commentID")
	commentID, err := strconv.Atoi(paramKeyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	request := new(modelcomment.RequestUpdate)

	userID := ctx.MustGet("user_id")
	request.UserID = userID.(uint)

	err = ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	update, err := c.srv.Update(*request, uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	ctx.JSON(http.StatusAccepted, helper.NewResponse(http.StatusAccepted, update, nil))
	return
}

func (c *controller) Delete(ctx *gin.Context) {
	paramKeyID := ctx.Param("commentID")
	commentID, err := strconv.Atoi(paramKeyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	err = c.srv.Delete(uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusAccepted, "Your comment successfully deleted", nil))
	return
}

func New(srv servicecomment.ServiceComment) ControllerComment {
	return &controller{srv: srv}
}
