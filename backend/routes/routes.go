package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// public

func Init() {
	router = gin.Default()
}

func SetupRoutes() {
	router.GET("/get/:item/:id", getItem)
}

func Run(host string) {
	router.Run(host)
}

// private

func getItem(ctx *gin.Context) {
	item := ctx.Param("item")
	id := ctx.Param("id")

	switch item {
	case "user":
		getUserById(ctx, id)
	default:
		ctx.IndentedJSON(http.StatusNotFound, errors.New("Unknown end point"))
	}
}
