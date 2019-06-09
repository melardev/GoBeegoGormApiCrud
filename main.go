package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/melardev/GoBeegoGormApiCrud/infrastructure"
	"github.com/melardev/GoBeegoGormApiCrud/models"
	_ "github.com/melardev/GoBeegoGormApiCrud/routers"
	"github.com/melardev/GoBeegoGormApiCrud/seeds"
	"os"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}

func main() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}

	// if in develop mode
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	database := infrastructure.OpenDbConnection()
	defer database.Close()
	migrate(database)
	seeds.Seed(database)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Print SQL statements executed
	orm.Debug = true
	beego.Run()
}
