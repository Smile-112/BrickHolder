package userActions

import (
	"autentification-service/models"

	//models "authentication-service/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"Логин"`
	Password string `json:"Пароль"`
}
type TokenResponse struct {
	Token string `json:"token"`
}

// @Summary      Login User
// @Description  Login User with name and password
// @Accept      json
// @Produce      json
// @Param      body body Credentials true "User"
// @Success    200 {object} TokenResponse "Успешный вход пользователя "
// @Router       /login [post]
// @Tags Autentification
func Login(c *gin.Context) {
	var creds Credentials
	//var user models.User
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Недопустимый запрос"})
		return
	}

	/*if err := dataBase.Db.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Пользователь не найден"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Неверный логин или пароль"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Неверный логин или пароль"})
		return
	}

	token, err := Secrets.GenerateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})*/
}

// @Summary      Register User
// @Description  Register User with name and password
// @Accept      json
// @Produce      json
// @Param      body body Credentials true "User"
// @Success    200 "Успешная регистрация нового пользователя"
// @Router       /register [post]
// @Tags Registration
func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Недопустимый запрос"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: creds.Username,
		Password: string(hash),
	}

	if err := chechUserLogin(user.Username); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	/*if err := dataBase.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	*/
	c.Status(http.StatusCreated)
}

func chechUserLogin(userLogin string) error {
	//var user models.User

	/*if err := dataBase.Db.Where("username = ?", userLogin).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	*/
	log.Printf("ERROR: пользователь стаким логином уже существует %s", userLogin)
	return fmt.Errorf("Пользователь с таким логином уже существует")

}
