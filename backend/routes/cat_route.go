package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"prgc/repo"
)

func getAllCat(ctx *gin.Context) {
    cat_repo := repo.NewCategorieRepo()

    categories, err := cat_repo.GetAllCategories()
    if err != nil {
        panic(err)
    }

    ctx.Header("Access-Control-Allow-Origin", "*")
    ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
    ctx.IndentedJSON(http.StatusOK, categories)
}

