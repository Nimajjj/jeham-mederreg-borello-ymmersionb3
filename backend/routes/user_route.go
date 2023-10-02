package routes

import (
	_ "fmt"
	_ "net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

func getUserById(ctx *gin.Context, idStr string) {
	/*
		    id, err := strconv.Atoi(idStr)
			if err != nil {
				return
			}

			userRepo := repo.NewUserRepo()

			user, err := userRepo.GetUserByID(id)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			fmt.Println("User: ", user)
			ctx.IndentedJSON(http.StatusOK, user)
	*/
}
