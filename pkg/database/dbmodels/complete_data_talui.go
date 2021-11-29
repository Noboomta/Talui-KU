package dbmodels

import (
	"fmt"
	"talui/pkg/database"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

type CompleteTalui struct {
	Id         int    `json:"service_name"`
	Enter      string `json:"enter"`
	Down       string `json:"down"`
	Enter_time string `json:"enter_time"`
	Down_time  string `json:"down_time"`
}

func GetAllTalui() ([]CompleteTalui, error) {
	fmt.Println("complete talui getall")
	db := database.DB
	results, err := db.Query("SELECT * FROM talui_using")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer results.Close()

	taluis := []CompleteTalui{}
	for results.Next() {
		var talui CompleteTalui
		results.Scan(&talui.Id, &talui.Enter, &talui.Enter_time, &talui.Down, &talui.Down_time)
		// fmt.Println(talui.Id, talui.Enter, talui.Enter_time, talui.Down, talui.Down_time)
		taluis = append(taluis, talui)
	}

	return taluis, nil
}
