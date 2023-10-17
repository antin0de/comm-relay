package main

import (
	"log"
	"os"

	"antin0.de/comm-relay/handlers"
	"antin0.de/comm-relay/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlDsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	models.Migrate(db)

	h := handlers.HandlerParams{Db: db}

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	cookieSecret := os.Getenv("COOKIE_SECRET")
	store := cookie.NewStore([]byte(cookieSecret))
	r.Use(sessions.Sessions("session", store))

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"root": os.Getenv("PASSWORD"),
	}))

	authorized.GET("/ping", h.Ping())
	authorized.POST("/createChannel", h.CreateChannel())
	authorized.POST("/updateChannel", h.UpdateChannel())
	authorized.POST("/deleteChannel", h.DeleteChannel())
	authorized.POST("/listChannels", h.ListChannels())
	authorized.POST("/createEmailConnector", h.CreateEmailConnector())
	authorized.POST("/updateEmailConnector", h.UpdateEmailConnector())
	authorized.POST("/deleteEmailConnector", h.DeleteEmailConnector())
	authorized.POST("/sendMessage", h.SendMessage())

	listenAddress := os.Getenv("LISTEN_ADDRESS")
	r.Run(listenAddress)
}
