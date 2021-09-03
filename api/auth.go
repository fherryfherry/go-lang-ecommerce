package api

import (
	"ecommerce/database"
	"ecommerce/helpers"
	"ecommerce/models"
	"ecommerce/repositories/tokenRepository"
	"ecommerce/repositories/userRepository"
	"github.com/gin-gonic/gin"
	"time"
)

func AuthRequestToken(c *gin.Context) {
	username := "ferry"
	password := "123456"

	if username == c.PostForm("username") && password == c.PostForm("password") {
		db := database.Connect(c)
		var accessToken,refreshToken,ipAddress,userAgent string
		var expiredAt time.Time

		accessToken = helpers.Base64Encode(helpers.RandomString(128))
		refreshToken = helpers.Base64Encode(helpers.RandomString(128))
		ipAddress = c.Request.RemoteAddr
		userAgent = c.Request.UserAgent()
		expiredAt = time.Now().AddDate(0,0,3)

		// Insert into token table
		token := models.Tokens{
			ExpiredAt: expiredAt,
			IpAddress: ipAddress,
			UserAgent: userAgent,
			AccessToken: accessToken,
			RefreshToken: refreshToken}
		db.Create(&token)

		c.JSON(200, gin.H{"status": 1, "message": "ok","access_token":accessToken,"refresh_token":refreshToken,"expired_at":expiredAt})
		return
	} else {
		c.JSON(400, gin.H{"status": 0,"message":"Invalid credential"})
		return
	}
}

func AuthRefreshToken(c *gin.Context) {
	db := database.Connect(c)

	data, total := tokenRepository.FindByRefreshToken(db, helpers.GetAuthorizationToken(c))
	if total == 0 {
		c.JSON(400, gin.H{"status":0,"message":"Invalid token!"})
		return
	}

	data.AccessToken = helpers.Base64Encode(helpers.RandomString(128))
	data.RefreshToken = helpers.Base64Encode(helpers.RandomString(128))
	data.ExpiredAt = time.Now().AddDate(0,0,3)
	db.Save(&data)

	c.JSON(200, gin.H{"status": 1, "message": "ok","access_token":data.AccessToken,"refresh_token":data.RefreshToken,"expired_at":data.ExpiredAt})
}

func AuthLogin(c *gin.Context) {
	if c.Param("email") != "" && c.Param("password") != "" {
		db := database.Connect(c)

		data, total := userRepository.FindByEmail(db, c.Param("email"))
		if total == 0 {
			c.JSON(400,gin.H{"status":0,"message":"The email is not registered at our system!"})
			return
		}

		if helpers.CheckPasswordHash(c.Param("password"), data.Password) == true {

			// Update token with the user logged in
			token, _ := tokenRepository.FindByAccessToken(db, helpers.GetAuthorizationToken(c))
			token.Users = data
			db.Save(&token)

			c.JSON(200, gin.H{"status":1,"message":"ok"})
		}

	} else {
		c.JSON(400,gin.H{"status":0,"message":"Credential is required"})
	}
}

func AuthLogout(c *gin.Context) {
	db := database.Connect(c)

	token, _ := tokenRepository.FindByAccessToken(db, helpers.GetAuthorizationToken(c))
	db.Delete(&token)

	c.JSON(200,gin.H{"status":1,"message":"You have been log out!"})
}
