package controllers

import (
	"fmt"
	"net/http"
	"time"
	"ums/connection"
	"ums/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type JwtCustomClaim struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(credential *models.Credential) string {
	claims := &JwtCustomClaim{
		credential.Email,
		credential.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("myjwtsecret"))
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return tokenString
}

func Login(c echo.Context) error {
	claimCred := new(models.Credential)
	if err := c.Bind(claimCred); err != nil {
		return err
	}
	var credential models.Credential
	sqlStatement := "SELECT * FROM credentials WHERE email=$1"
	err := connection.DB.QueryRow(sqlStatement, claimCred.Email).Scan(&credential.Id, &credential.Password, &credential.Role, &credential.Status, &credential.Email)
	if err != nil {
		return err
	}
	match := CheckPasswordHash(claimCred.Password, credential.Password)
	if match {
		token := generateToken(&credential)
		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	} else {
		return echo.ErrUnauthorized
	}
}
