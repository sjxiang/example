package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/example/pkg/apperr"
)



func (app *App) Register(ctx *gin.Context) {
	type Credentials struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	var req Credentials
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return 
	}

	// validate

	if err := app.users.InsertOne(req.Name, req.Email, req.Password); err != nil {
		// err 友好提示
		ctx.JSON(apperr.HTTPStatus(err), err)
		return
	}


}