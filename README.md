# Greenico API

Weather station with Raspberry Pi Pico - "Greenico"

[GreenicoAPI](https://github.com/anras5/GreenicoAPI) is only a part of `Greenico` project.

The first part is [Greenico on Raspberry Pi Pico](https://github.com/anras5/Greenico). Be sure to check it out also.

## Table of contents

1. [The purpose of the project](#Purpose)
2. [Description](#Description)
3. [Technical information](#Technical)
4. [Contact](#Contact)

## Purpose

The aim of the project is to allow the user to view the current state of the weather by designing and building a weather
station. The preview will be possible through a browser, i.e. a web application in the form of clear graphs and on the
display attached to the weather station.

## Description

Go was chosen as the programming language for this part of the project with Gin framework due to the ease of
implementing the RESTful API and the simplicity of connecting the application with
Postgres database. The web application was built using Go templates and language JavaScript.
The Postgres database was created using the Soda migration tool.

## Available endpoints

API endpoints:

```go
// API
router.GET("/api/reading", handlers.GetReading(db))
router.GET("/api/readings", handlers.GetReadings(db))
router.GET("/api/all-readings", handlers.GetAllReadings(db))
router.POST("/api/reading", handlers.InsertReading(db))
router.DELETE("/api/reading", handlers.DeleteReading(db))
```

Web application endpoints:

```go
// WEB APP
router.GET("/", handlers.MainRender(db))
router.GET("/voc", handlers.PageRender(db, "voc"))
router.GET("/temperature", handlers.PageRender(db, "temperature"))
router.GET("/humidity", handlers.PageRender(db, "humidity"))
router.GET("/light", handlers.PageRender(db, "light"))
router.GET("/pressure", handlers.PageRender(db, "pressure"))
router.GET("/uv", handlers.PageRender(db, "uv"))
```

## Technical

API was written in **Go** programming language with the use of **Gin framework**. All requirements are listed
in `go.mod` file.

### How to run the app

1. Create (in the root directory of the project) and fill values in `.env` file:
```
HOSTNAME=
PORT=
DBNAME=
DBUSER=
DB_PASSWD=
```
2. Run `run.bat` file

## Contact
Contact me at:
anras1filip@gmail.com