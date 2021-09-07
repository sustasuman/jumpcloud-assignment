package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type HashRequest struct {
	Password string `json:"password"`
}

type PostHashReresponse struct {
	Id int64 `json:"id"`
}

func readPasswordValue(ctx *gin.Context) string {

	//read from json request
	password := ctx.PostForm("password")
	if password == "" {
		var req HashRequest
		if err := ctx.BindJSON(&req); err != nil {
			if req.Password == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{"Invalid Request": err.Error()})
				return ""
			}
		}
		return req.Password
	}
	return password
}

func PostHash(ctx *gin.Context) {

	password := readPasswordValue(ctx)
	if password == "" {
		return
	}
	encodedString, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		HandleDbError(err, ctx)
		return
	}
	hashId, dberr := SaveHash(string(encodedString))
	if dberr != nil {
		HandleDbError(dberr, ctx)
		return
	} else {
		response := PostHashReresponse{Id: hashId}
		ctx.JSON(http.StatusOK, response)
		return
	}
}
