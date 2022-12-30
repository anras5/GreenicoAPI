package handlers

import (
	"database/sql"
	"fmt"
	"github.com/anras5/GreenicoAPI/internal/json_models"
	"github.com/anras5/GreenicoAPI/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetReading(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newReading json_models.SensorData

		query := `SELECT temperature, humidity, pressure, voc, uv, light
				  FROM "READINGS" r 
				  WHERE r.pico_id = 1 AND r.created_at = (SELECT max(created_at) FROM "READINGS")`
		row := db.QueryRow(query)

		err := row.Scan(
			&newReading.Temperature,
			&newReading.Humidity,
			&newReading.Pressure,
			&newReading.VOC,
			&newReading.UV,
			&newReading.Light,
		)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, newReading)

	}
}

func GetReadings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var startDate, endDate time.Time
		start_date, _ := c.GetQuery("startdate")
		end_date, _ := c.GetQuery("enddate")

		layout := "2006-01-02"
		startDate, err := time.Parse(layout, start_date)
		if err != nil {

			c.IndentedJSON(http.StatusBadRequest, gin.H{"info": "wrong format on startdate"})
			return
		}
		endDate, err = time.Parse(layout, end_date)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"info": "wrong format on enddate"})
			return
		}

		query := `SELECT id, temperature, humidity, pressure, voc, uv, light, created_at, updated_at, pico_id
				  FROM "READINGS" r
				  WHERE created_at > $1 AND created_at < $2`
		rows, err := db.Query(query, startDate, endDate)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error getting data from database"})
			return
		}
		defer rows.Close()

		var readings []models.SensorDataModel

		for rows.Next() {
			var newReading models.SensorDataModel
			err = rows.Scan(
				&newReading.ID,
				&newReading.Temperature,
				&newReading.Humidity,
				&newReading.Pressure,
				&newReading.VOC,
				&newReading.UV,
				&newReading.Light,
				&newReading.CreatedAt,
				&newReading.UpdatedAt,
				&newReading.PicoId,
			)
			readings = append(readings, newReading)
		}
		c.IndentedJSON(http.StatusOK, readings)
	}
}

func GetAllReadings(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := `SELECT id, temperature, humidity, pressure, voc, uv, light, created_at, updated_at, pico_id
				  FROM "READINGS" r`
		rows, err := db.Query(query)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error getting data from database"})
			return
		}
		defer rows.Close()

		var readings []models.SensorDataModel

		for rows.Next() {
			var newReading models.SensorDataModel
			err = rows.Scan(
				&newReading.ID,
				&newReading.Temperature,
				&newReading.Humidity,
				&newReading.Pressure,
				&newReading.VOC,
				&newReading.UV,
				&newReading.Light,
				&newReading.CreatedAt,
				&newReading.UpdatedAt,
				&newReading.PicoId,
			)
			readings = append(readings, newReading)
		}

		log.Println(readings)
		c.IndentedJSON(http.StatusOK, readings)
	}
}

func InsertReading(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newReading json_models.SensorData

		if err := c.BindJSON(&newReading); err != nil {
			return
		}

		sqlStatement := `INSERT INTO "READINGS" (temperature, humidity, pressure, voc, uv, light, created_at, updated_at, pico_id)
					 	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
		_, err := db.Exec(sqlStatement,
			newReading.Temperature,
			newReading.Humidity,
			newReading.Pressure,
			newReading.VOC,
			newReading.UV,
			newReading.Light,
			time.Now(),
			time.Now(),
			1,
		)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusCreated, newReading)
	}
}

func DeleteReading(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := `DELETE FROM "READINGS" 
				  WHERE DATE_PART('day', current_date - created_at) > 7
				  RETURNING temperature, humidity, pressure, voc, uv, light`

		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}

func PageRender(db *sql.DB, pageTopic string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query string
		switch pageTopic {
		case "voc":
			query = `SELECT id, voc, created_at
				  FROM "READINGS"`
		case "temperature":
			query = `SELECT id, temperature, created_at
				  FROM "READINGS"`
		case "humidity":
			query = `SELECT id, humidity, created_at
				  FROM "READINGS"`
		case "light":
			query = `SELECT id, light, created_at
				  FROM "READINGS"`
		case "pressure":
			query = `SELECT id, pressure, created_at
				  FROM "READINGS"`
		case "uv":
			query = `SELECT id, pressure, created_at
				  FROM "READINGS"`
		}
		rows, err := db.Query(query)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error getting data from database"})
			return
		}
		defer rows.Close()

		var readings []models.SensorDataModel

		for rows.Next() {
			var newReading models.SensorDataModel
			switch pageTopic {
			case "voc":
				err = rows.Scan(
					&newReading.ID,
					&newReading.VOC,
					&newReading.CreatedAt,
				)
			case "temperature":
				err = rows.Scan(
					&newReading.ID,
					&newReading.Temperature,
					&newReading.CreatedAt,
				)
			case "humidity":
				err = rows.Scan(
					&newReading.ID,
					&newReading.Humidity,
					&newReading.CreatedAt,
				)
			case "light":
				err = rows.Scan(
					&newReading.ID,
					&newReading.Light,
					&newReading.CreatedAt,
				)
			case "pressure":
				err = rows.Scan(
					&newReading.ID,
					&newReading.Pressure,
					&newReading.CreatedAt,
				)
			case "uv":
				err = rows.Scan(
					&newReading.ID,
					&newReading.UV,
					&newReading.CreatedAt,
				)
			}
			readings = append(readings, newReading)
		}
		c.HTML(http.StatusOK, fmt.Sprintf("%s.page.gohtml", pageTopic), gin.H{"data": readings})
	}
}

func MainRender(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newReading json_models.SensorData

		query := `SELECT temperature, humidity, pressure, voc, uv, light
				  FROM "READINGS" r 
				  WHERE r.pico_id = 1 AND r.created_at = (SELECT max(created_at) FROM "READINGS")`
		row := db.QueryRow(query)

		err := row.Scan(
			&newReading.Temperature,
			&newReading.Humidity,
			&newReading.Pressure,
			&newReading.VOC,
			&newReading.UV,
			&newReading.Light,
		)
		if err != nil {
			panic(err)
		}
		log.Println(newReading)

		c.HTML(http.StatusOK, "index.page.gohtml", newReading)
	}
}
