package main

import (
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
	r.GET("/get-user", getUser)
	r.GET("/user/:userID/chef", addChef)
	r.GET("/chef", getChef)
	r.GET("/:chefID/course", addCourse)
	r.GET("/get-course", getCourse)
	r.Run(":8080") // listen and serve on localhost:8080
}

func ping(c *gin.Context) {
	c.JSON(200, "pong!!!!!!")
}

func addUser(c *gin.Context) {
	email := "abcd"
	allergie1 := model.Ingredient{Name: "egg"}
	allergie2 := model.Ingredient{Name: "meat"}
	image := model.Image{ImageUrl: "image url test"}
	var allergies []model.Ingredient
	allergies = append(allergies, allergie1)
	allergies = append(allergies, allergie2)
	user := model.User{Email: &email, Allergies:  &allergies, ProfileImage: &image}
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

func addChef(c *gin.Context) {
	userID := c.Param("userID")
	u64, _ := strconv.ParseUint(userID, 10, 32)
	chef := model.Chef{Name: "Chef Test",UserID: uint(u64)}
	userService.AddChef(chef)
	c.JSON(200, "pong!!!!!!")
}

func getChef(c *gin.Context) {
	chefId := c.Query("id")
	u64, _ := strconv.ParseUint(chefId, 10, 32)
	chef := userService.GetChef(uint(u64))

	c.JSON(200, chef)
}

func addCourse(c *gin.Context) {
	chefID := c.Param("chefID")
	u64, _ := strconv.ParseUint(chefID, 10, 32)
	image := model.Image{ImageUrl: "image url test course"}
	menu1 := model.CourseMenu{Name: "menu1",Sequence: 1}
	menu2 := model.CourseMenu{Name: "menu2",Sequence: 2}
	menu3 := model.CourseMenu{Name: "menu3",Sequence: 3}
	var menues []model.CourseMenu
	menues = append(menues, menu1)
	menues = append(menues, menu2)
	menues = append(menues, menu3)
	course := model.Course{Status: model.CourseStatus("PUBLISHED"),ChefID: uint(u64), CoverImage: &image, CourseMenus: &menues}
	cour, _ := userService.AddCourse(course)
	c.JSON(200, cour)
}

func getCourse(c *gin.Context) {
	courseID := c.Query("id")
	u64, _ := strconv.ParseUint(courseID, 10, 32)
	course := userService.GetCourse(uint(u64))

	c.JSON(200, course)
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
	db.AutoMigrate(&model.Ingredient{})
	db.AutoMigrate(&model.TypeOfCuisine{})
	db.AutoMigrate(&model.Chef{}, &model.ChefLocation{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Image{})
	db.AutoMigrate(&model.Course{}, &model.CourseMenu{})


	DB = db
}