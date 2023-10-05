package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"prgc/repo"
)

const DEFAULT = 0b0000
const CATEGORIES = 0b0001
const ORDER = 0b0010
const KEYWORD = 0b0100

func getPebble(ctx *gin.Context) {
	id_str := ctx.Param("id")
    id, err := strconv.Atoi(id_str)
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, errors.New("Invalid pebble id"))
        return
    }

    pebble_repo := repo.NewPebbleRepo()
    pebble, err := pebble_repo.GetPebbleById(id)

    if (err != nil) {
        ctx.IndentedJSON(http.StatusNotFound, errors.New("Unknown pebble id"))
        return
    }

    ctx.Header("Access-Control-Allow-Origin", "*")
    ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
    ctx.IndentedJSON(http.StatusOK, pebble)
}


func searchPebble(ctx *gin.Context) {
    // get params
    categories := ctx.Param("categories")
    order := ctx.Param("order")
    keywords := ctx.Param("keywords")

    // format params
    var categories_list []string
    dec := json.NewDecoder(strings.NewReader(categories))
    err := dec.Decode(&categories_list)
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, err)
    }

    // set flags
    var flags int64
    if (len(categories_list) != 0) {
        flags |= CATEGORIES
    }
    if (order != "nil") {
        flags |= ORDER
    }
    if (keywords != "nil") {
        flags |= KEYWORD
    }

    pebble_repo := repo.NewPebbleRepo()
    pebbles, err := pebble_repo.SearchPebbles(repo.Search{flags, categories_list, order, keywords})
    if (err != nil) {
        ctx.IndentedJSON(http.StatusNotFound, err)
    }


    // send response
    ctx.IndentedJSON(http.StatusOK, pebbles)
}
