package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	//godotenv.Load(".env")
	//envConfig := getConfig()
	//
	r := gin.Default()
	//Initialize(envConfig.DbConfig)

	r.GET("/ping", ping)
	r.Run(":8080") // listen and serve on localhost:8080
}

func ping(c *gin.Context) {
	c.JSON(200, "pong!!!!!!")
}

func getConfig() EnvConfig {
	return EnvConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DbConfig: MySQLConfig{
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Host:     os.Getenv("MYSQL_HOSTNAME"),
			Port:     os.Getenv("MYSQL_PORT"),
			DbName:   os.Getenv("MYSQL_DBNAME"),
		},
	}
}

func Initialize(config MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	if err = db.AutoMigrate(&TimeAvailable{}); err == nil && db.Migrator().HasTable(&TimeAvailable{}) {
		if err := db.First(&TimeAvailable{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			TimeAvailableF(db)
		}
	}

	DB = db
}

func TimeAvailableF(db *gorm.DB) {
	var times = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	creates := []TimeAvailable{}
	for index, time := range times {
		ID := (index * 2) + 1
		t := TimeAvailable{
			ID:       uint(ID),
			Time:     strconv.Itoa(time) + ":00",
			Sequence: ID,
		}
		creates = append(creates, t)
		ID = (index * 2) + 2
		t = TimeAvailable{
			ID:       uint(ID),
			Time:     strconv.Itoa(time) + ":30",
			Sequence: ID,
		}
		creates = append(creates, t)
	}
	db.Create(&creates)
}