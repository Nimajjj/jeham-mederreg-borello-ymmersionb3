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

    ctx.IndentedJSON(http.StatusOK, categories)
}

