package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "talui/pkg/database/dbmodels"
)

func GetMainPage(c *fiber.Ctx) error {
	fmt.Println("/")
	// router := []string{
	// 	"/on/getAllTalui",
	// 	"/on/insertTalui",
	// 	"/complete/GetAllTalui",
	// 	"/complete/insertTalui",
	// 	"/arriveAt",
	// }
	api_docs := "https://talui-ku-server.herokuapp.com/swagger/index.html"
	myHtml := fmt.Sprintf("<p>Check out our Swagger api here!</p><br><a href='%s'>api-docs</a>", api_docs)
	
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    return c.SendString(myHtml)
}