package helpers

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)
import "math/rand"

var validate *validator.Validate

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ValidateForm(StructValidator interface{}) error {
	validate = validator.New()
	err := validate.Struct(StructValidator)
	if err != nil {
		var errorText string

		for _, err := range err.(validator.ValidationErrors) {
			errorText += err.Field() + ","
		}

		if len(errorText) > 0 {
			return errors.New("Please complete these fields: " + strings.Trim(errorText,","))
		}
	}

	return nil
}

func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func GetAuthorizationToken(c *gin.Context) string {
	auth := c.GetHeader("Authorization")
	auth = strings.Replace(auth, "Bearer ", "", 1)
	auth = strings.Replace(auth, "Basic ","", 1)
	return auth
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}