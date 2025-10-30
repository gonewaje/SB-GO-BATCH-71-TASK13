package controllers

import (
	"bioskop/db"
	"bioskop/repository"
	"bioskop/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBioskop(ctx *gin.Context) {
	var (
		result gin.H
	)

	bioskop, err := repository.GetAllBioskop(db.DB)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": bioskop,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func InsertBioskop(ctx *gin.Context) {
	// var person structs.Bioskop
	var newBioskop structs.Bioskop

	// err := ctx.BindJSON(&newBioskop)

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newBioskop.Nama == "" || newBioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Name and Location cant be empty",
		})
		return
	}

	err := repository.InsertBioskop(db.DB, newBioskop)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newBioskop)

	// err := c.BindJSON(&person)
	// if err != nil {
	// 	panic(err)
	// }

	// err = repository.InsertBioskop(database.DB, person)
	// if err != nil {
	// 	panic(err)
	// }

	// c.JSON(http.StatusOK, person)
}

func UpdateBioskop(c *gin.Context) {
	var newBioskop structs.Bioskop
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&newBioskop)
	if err != nil {
		panic(err)
	}

	newBioskop.ID = id

	err = repository.UpdateBioskop(db.DB, newBioskop)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, newBioskop)
}

func DeleteBioskop(c *gin.Context) {
	var newBioskop structs.Bioskop
	id, _ := strconv.Atoi(c.Param("id"))

	newBioskop.ID = id
	err := repository.DeleteBioskop(db.DB, newBioskop)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, newBioskop)
}
