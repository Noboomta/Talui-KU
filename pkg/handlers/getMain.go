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

	myHtml := "<h3>GET /on/getAllTalui</h3>" +
	"<br>" +
	"<h3>POST /on/insertTalui</h3>" +
	"<br>" +
	"<h3>GET /complete/GetAllTalui</h3>" +
	"<br>" +
	"<h3>POST /complete/insertTalui</h3>" +
	"<br>" +
	"<h3>POST /arriveAt</h3>" +
	"<br>" +
	"<h3>GET /station/using/entry</h3>" +
	"<br>" +
	"<h3>GET /station/using/dest</h3>" +
	"<br>" +
	"<h3>GET /station/find/entry</h3>" +
	"<br>" +
	"<h3>GET /station/find/dest</h3>" +
	"<br>"
	
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
    return c.SendString(myHtml)
}