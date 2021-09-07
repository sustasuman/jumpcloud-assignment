package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HashResponse struct {
	Hash string `json:"hash"`
}

func GetHash(ctx *gin.Context) {
	id := ctx.Param("id")
	idNum, error := strconv.ParseInt(id, 10, 64)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, "invalid request")
	}
	encodedString, err := FetchHash(idNum)
	if err != nil {
		HandleDbError(err, ctx)
	} else if encodedString == "" {
		ctx.HTML(404, "Resource not found", gin.H{})
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		response := HashResponse{Hash: encodedString}
		ctx.JSON(http.StatusOK, response)
	}
}
