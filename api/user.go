package api

import (
	"ecommerce/database"
	"ecommerce/helpers"
	"ecommerce/repositories/tokenRepository"
	"github.com/gin-gonic/gin"
	"os"
)

type UpdateProfileRule struct {
	Name  string `form:"name" validate:"required"`
	Email string `form:"name" validate:"required"`
}

func UserUpdateProfile(c *gin.Context) {
	db := database.Connect(c)

	ruleData := UpdateProfileRule{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email")}

	err := helpers.ValidateForm(ruleData)
	if err != nil {
		c.JSON(400, gin.H{"status": 0, "message": err.Error()})
		return
	}

	tokenData, found := tokenRepository.FindByAccessToken(db, helpers.GetAuthorizationToken(c))
	if found == 0 {
		c.JSON(400, gin.H{"status": 0, "message": "User is not found!"})
		return
	}

	user := tokenData.Users
	user.Name = ruleData.Name
	user.Email = ruleData.Email

	file, _ := c.FormFile("file")
	dest, destErr := os.Getwd()
	if destErr != nil {
		c.JSON(500, gin.H{"status": 0, "message": "File can't upload!"})
		return
	}
	fileErr := c.SaveUploadedFile(file, dest+"/static")
	if fileErr == nil {
		user.Photo = file.Filename
	}

	db.Save(user)

	c.JSON(200, gin.H{"status": 1, "message": "ok"})
}
