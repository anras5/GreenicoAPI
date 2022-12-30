package main

import (
	"database/sql"
	"github.com/anras5/GreenicoAPI/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func main() {

	err := godotenv.Load()

	psqlInfo := os.Getenv("PSQL_INFO")

	// connect to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*/**")

	// API
	router.GET("/api/reading", handlers.GetReading(db))
	router.GET("/api/readings", handlers.GetReadings(db))
	router.GET("/api/all-readings", handlers.GetAllReadings(db))
	router.POST("/api/reading", handlers.InsertReading(db))
	router.DELETE("/api/reading", handlers.DeleteReading(db))

	// WEB APP
	router.GET("/", handlers.MainRender(db))
	router.GET("/voc", handlers.PageRender(db, "voc"))
	router.GET("/temperature", handlers.PageRender(db, "temperature"))
	router.GET("/humidity", handlers.PageRender(db, "humidity"))
	router.GET("/light", handlers.PageRender(db, "light"))
	router.GET("/pressure", handlers.PageRender(db, "pressure"))
	router.GET("/uv", handlers.PageRender(db, "uv"))
	router.Static("/static", "./static/")
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
