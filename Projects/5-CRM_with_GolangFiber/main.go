package main

import (
	"crm/lead"
	"crm/database"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/leads",lead.GetLeads)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase(){
var err error
database.DBConn,err = gorm.Open("sqlite3","leads.db")
if err !=nil {
	panic("failed to connect to the database")	
} 
fmt.Println("Connection Successful!")
database.DBConn.AutoMigrate(&lead.Lead{})
fmt.Println("Database migrated")

}
