package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type url struct {
	gorm.Model
	url       string
	shortened string
}

func initMigration() {

	db, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Connected to DB brrrrrr")
	}

	db.AutoMigrate(&url{})
}
func generateKey() string {

	key := make([]byte, 6)
	for i := range key {
		key[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(key)
}

func addURL(RecURL string) {

	var randomKey = generateKey()
	var shortenedURL string = "sigh.gq/" + randomKey

	println(shortenedURL)

	db.Create(&url{url: RecURL, shortened: shortenedURL})

}

func createURL(c *gin.Context) {

}

func main() {

	initMigration()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var randomKey = generateKey()
		c.JSON(200, gin.H{
			"status": randomKey,
		})
	})

	r.POST("/add", createURL)

	r.Run()
}
