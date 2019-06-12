package controllers

import (
	"fmt"

	"../model"

	"../config"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var (
		user     model.User
		response config.Response
	)
	if err := idb.DB.Where("id = ?", id).First(&user).Delete(&user).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		response.Message = "Data user telah dihapus"
		response.Status = 200
		response.Data = user
		c.JSON(response.Status, response)
	}
}

func (idb *InDB) UpdateUser(c *gin.Context) {

	var (
		user     model.User
		response config.Response
	)
	id := c.Params.ByName("id")

	if err := idb.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(500)
	} else {
		c.BindJSON(&user)
		idb.DB.Save(&user)
		response.Message = "Data user telah dibuat"
		response.Status = 200
		response.Data = user
		c.JSON(response.Status, response)
	}

}

func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user     model.User
		response config.Response
	)
	c.BindJSON(&user)

	if err := idb.DB.Create(&user).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		response.Message = "Data user telah dibuat"
		response.Status = 200
		response.Data = user
		c.JSON(response.Status, response)
	}

}

func (idb *InDB) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var (
		user     model.User
		response config.Response
	)
	if err := idb.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		response.Message = "Data user berhasil ditampilkan"
		response.Status = 200
		response.Data = user
		c.JSON(response.Status, response)
	}

}

func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		users    []model.User
		response config.Response
	)

	if err := idb.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		response.Message = "Data seluruh user ditampilkan"
		response.Status = 200
		response.Data = users
		c.JSON(response.Status, response)
	}

}
