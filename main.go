package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func chooseRepo() *infra.MysqlDriver {
	var c *infra.MysqlDriver
	dbConfig := c.Configure("./configs", "mysql")
	dbConfig.NewDriver()
	return c
}

func main() {
	repo := chooseRepo()
	service := backupPolicy.NewService(repo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//Routers
	e.Static("/*", "public/dist/")

	e.Static("/public/*", "public/reportes/")

	e.File("/", "public/dist/index.html")
	//e.File("/*", "public/dist/index.html")
	//	e.POST("/login", controllers.Login)
	e.Logger.Fatal(e.Start(":1234"))
}
