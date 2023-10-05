package routes

import (
	"fmt"
	"net/http"
	"prgc/model"
	"prgc/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

// routes functions
func getUserCart(ctx *gin.Context) {
    user_id := ctx.Param("user")
    fmt.Println(user_id)

    

    ctx.IndentedJSON(http.StatusOK, nil)
}


func addPebbleToCart(ctx *gin.Context) {
    user_id, err := strconv.Atoi(ctx.Param("user"))
    testErr(err)
    pebble_id, err := strconv.Atoi(ctx.Param("pebble"))
    testErr(err)
    quantity, err := strconv.Atoi(ctx.Param("quantity"))
    testErr(err)
    
    cart_repo := repo.NewCartRepo()
    if (!cart_repo.UserHasCart(user_id)) {
        createCartWithItem(user_id, pebble_id, quantity)
    } else {
        cart_repo.AddPebbleToCart(user_id, pebble_id, quantity)
    }

    ctx.IndentedJSON(http.StatusOK, nil)
}

func removePebbleFromCart(ctx *gin.Context) {

}



// sub func
func createCartWithItem(user_id, pebble_id, quantity int) model.JsonCart {
    pebble_repo := repo.NewPebbleRepo()
    cart_repo := repo.NewCartRepo()
    pebble, err := pebble_repo.GetPebbleById(pebble_id)
    testErr(err)
    pebble.ID = pebble_id   // AAAAAAAAAAAAAAAAAAAAAAAAAAAH

    content := make(map[*model.Pebble]int)
    content[&pebble] = quantity

    cart := model.Cart{UserID: user_id, Content: content}
    cart_repo.InsertNewCart(cart)

    return cart.JsonCompatible()
}
