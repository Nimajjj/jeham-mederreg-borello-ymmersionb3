package routes

import (
//    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


var router *gin.Engine

func testErr(err error) {
    if err != nil {
        panic(err)
    }
}

func Init() {
	router = gin.Default()
    //router.Use(cors.Default())
}

func SetupRoutes() {
    // pebble details
	router.GET("/pebble/:id", getPebble)

    // search pebbles
    router.GET("/pebbles/:categories/:order/:keywords", searchPebble)

    // categories
    router.GET("/categories/", getAllCat)

    // cart
    router.GET("/cart/:user", getUserCart)
    router.GET("/cart/add/:user/:pebble/:quantity", addPebbleToCart)
    router.GET("/cart/remove/:user/:pebble/:quantity", removePebbleFromCart)
}

func Run(host string) {
	router.Run(host)
}

