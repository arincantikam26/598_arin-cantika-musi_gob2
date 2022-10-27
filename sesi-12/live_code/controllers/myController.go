package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Myself struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	CodeParticipant string `json:"code_participant"`
}

var MyselfDatas = []Myself{}

func CreateMy(ctx *gin.Context) {
	var newMy Myself

	if err := ctx.ShouldBindJSON(&newMy); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newMy.ID = fmt.Sprintf("my%d", len(MyselfDatas)+1)

	MyselfDatas = append(MyselfDatas, newMy)

	ctx.JSON(http.StatusCreated, gin.H{
		"myself": newMy,
	})

}

func GetAllMy(ctx *gin.Context) {
	var myselfs []Myself

	for k, _ := range MyselfDatas {
		myselfs = append(myselfs, MyselfDatas[k])
	}

	if myselfs == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("Myself not found"),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"myself": myselfs,
	})

}
