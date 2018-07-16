package controllers

import (
	"go-crud-demo/model"
	"github.com/gin-gonic/gin"
	"go-crud-demo/forms"
)

var userModel = new(model.UserModel)

type UserController struct{}

//创建一个用户
func (user *UserController) Create(c *gin.Context){
	var data forms.CreateUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406,gin.H{
			"message":"Invalid form",
			"form":data,
		})
		c.Abort()
		return
	}

	err := userModel.Create(data)
	if err != nil {
		c.JSON(406,gin.H{
			"message":"user could not be created",
			"error":err,
		})
		c.Abort()
		return
	}

	c.JSON(200,gin.H{
		"message":"success",
	})
}
//查找所有用户
func (user *UserController) GetAll(c *gin.Context){
  list, err := userModel.GetAll()
  if err != nil {
  	c.JSON(404,gin.H{
  		"message":"Find Error",
  		"error":err.Error(),
		})
  	c.Abort()
	}else{
		c.JSON(200,gin.H{
			"data":list,
		})
	}
}
//查找一个用户
func (user *UserController) GetOne(c *gin.Context){
	id := c.Param("id")
	result, err := userModel.GetOne(id)
	if err != nil {
		c.JSON(404,gin.H{
			"message":"Find User Error",
			"error":err.Error(),
		})
		c.Abort()
	}else{
		c.JSON(200,gin.H{
			"data":result,
		})
	}
}
//更新一个用户
func (user *UserController) UpdateOne(c *gin.Context){
	id := c.Param("id")
	data := forms.UpdateUserCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406,gin.H{
			"message":"Invalid Parameters",
		})
		c.Abort()
		return
	}

	err := userModel.UpdateOne(id,data)
	if err != nil {
		c.JSON(406,gin.H{
			"message" : "Update Failed",
		})
		c.Abort()
		return
	}

	c.JSON(200,gin.H{
		"message" : "Update Success",
	})

}
//删除一个用户
func (user *UserController) DeleteOne(c *gin.Context){
	id := c.Param("id")
	err := userModel.DeleteOne(id)
	if err != nil {
		c.JSON(406,gin.H{
			"message" : "Delete Failed",
		})
		c.Abort()
		return
	}
	c.JSON(200,gin.H{
		"message" : "Delete Success",
	})

}