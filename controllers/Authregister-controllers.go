package controllers

import (
	"strconv"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/pulsarcoder/learn/reactWithgo/database"
	"github.com/pulsarcoder/learn/reactWithgo/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

// Register details
func Register(c *fiber.Ctx) error {

	var data map[string]string

	// error shortcut method & cleaner
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// if err != nil{
	// 	 return err
	// }
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.Db.Create(&user)
	return c.JSON(user)
}

//login details

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	var user models.User
	// query checking the details is there or not
	database.Db.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	//JWT create tokens

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		//the Issuer is our -> user
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewTime(3600000), //1 daya
	})
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "cloud not login",
		})
	}
	return c.JSON(token)
}
