package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"poc-model/model"
	"poc-model/service"
	"strconv"
)

var DB *gorm.DB
var userService service.UserService

func main() {
	godotenv.Load(".env")
	envConfig := getConfig()

	r := gin.Default()
	Initialize(envConfig.DbConfig)

	userService = service.UserServiceInitialize(DB)

	r.GET("/ping", ping)
	r.GET("/user", addUser)
	r.GET("/user2", addUser2)
	r.GET("/all-user", getUser)
	r.Run(":8080") // listen and serve on localhost:8080
}

func ping(c *gin.Context) {
	c.JSON(200, "pong!!!!!!")
}

func addUser(c *gin.Context) {
	email := "abcd"
	allergie1 := model.Ingredient{IsActive: true,Name: "egg"}
	allergie2 := model.Ingredient{IsActive: true,Name: "meat"}
	var allergies []model.Ingredient
	allergies = append(allergies, allergie1)
	allergies = append(allergies, allergie2)
	user := model.User{Email: &email, Allergies:  &allergies}
	userService.AddUser(user)

	c.JSON(200, "pong!!!!!!")
}

func getUser(c *gin.Context) {
	userId := c.Query("id")
	u64, _ := strconv.ParseUint(userId, 10, 32)
	user := userService.GetUser(uint(u64))

	c.JSON(200, user)
}

func addUser2(c *gin.Context) {
	email := "xyz"
	allergie1, _ := userService.GetIngredientById(uint(1))

	var allergies []model.Ingredient
	allergies = append(allergies, *allergie1)
	user := model.User{Email: &email, Allergies:  &allergies}
	userService.AddUser(user)

	c.JSON(200, "pong!!!!!!")
}

func getConfig() model.EnvConfig {
	return model.EnvConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DbConfig: model.MySQLConfig{
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Host:     os.Getenv("MYSQL_HOSTNAME"),
			Port:     os.Getenv("MYSQL_PORT"),
			DbName:   os.Getenv("MYSQL_DBNAME"),
		},
	}
}

func Initialize(config model.MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Ingredient{})
	db.AutoMigrate(&model.User{})

	if err = db.AutoMigrate(&model.TimeAvailable{}); err == nil && db.Migrator().HasTable(&model.TimeAvailable{}) {
		if err := db.First(&model.TimeAvailable{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			TimeAvailableF(db)
		}
	}

	DB = db
}

func TimeAvailableF(db *gorm.DB) {
	var times = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	creates := []model.TimeAvailable{}
	for index, time := range times {
		ID := (index * 2) + 1
		t := model.TimeAvailable{
			ID:       uint(ID),
			Time:     strconv.Itoa(time) + ":00",
			Sequence: ID,
		}
		creates = append(creates, t)
		ID = (index * 2) + 2
		t = model.TimeAvailable{
			ID:       uint(ID),
			Time:     strconv.Itoa(time) + ":30",
			Sequence: ID,
		}
		creates = append(creates, t)
	}
	db.Create(&creates)
}