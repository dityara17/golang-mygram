package controllerphoto

import (
	"github.com/arfan21/golang-mygram/helper"
	"github.com/arfan21/golang-mygram/model/modelphoto"
	"github.com/arfan21/golang-mygram/service/servicephoto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ControllerPhoto interface {
	Create(ctx *gin.Context)
	GetByUserID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	srv servicephoto.ServicePhoto
}

func New(srv servicephoto.ServicePhoto) ControllerPhoto {
	return &controller{srv: srv}
}

func (c *controller) Delete(ctx *gin.Context) {
	paramKeyID := ctx.Param("photoID")
	photoID, _ := strconv.Atoi(paramKeyID)
	err := c.srv.Delete(photoID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, "Your Photo has been successfully deleted", nil))
	return
}

func (c *controller) Update(ctx *gin.Context) {
	data := new(modelphoto.Request)

	err := ctx.ShouldBind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	photoParamID := ctx.Param("photoID")
	photoID, err := strconv.Atoi(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	// must have user id from bearer
	userID := ctx.MustGet("user_id")
	data.UserID = userID.(uint)

	update, err := c.srv.Update(*data, photoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, update, nil))
	return
}

func (c *controller) Create(ctx *gin.Context) {
	data := new(modelphoto.Request)

	err := ctx.ShouldBind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	// must have user id from bearer
	userID := ctx.MustGet("user_id")
	data.UserID = userID.(uint)

	response, err := c.srv.Create(*data)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

func (c *controller) GetByUserID(ctx *gin.Context) {
	response, err := c.srv.GetPhotos()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}
