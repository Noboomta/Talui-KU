package dbmodels

import (
	"fmt"
	"talui/pkg/database"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

type TaluiComplete struct { 
	Id          int     `json:"id"`
	Entry       string  `json:"entry"`
	Entry_ts    string  `json:"entry_ts"`
	Dest        string  `json:"dest"`
	Dest_ts     string  `json:"dest_ts"`
	Line        string  `json:"line"`
}

func GetAllTaluiComplete() ([]TaluiComplete, error) {
	fmt.Println("get all TaluiComplete")
	db := database.DB
	result, err := db.Query("SELECT * FROM `talui_complete`")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	defer result.Close()
	
	dataArray := []TaluiComplete{}
	for result.Next() {
		var data TaluiComplete
		result.Scan(&data.Id, &data.Entry, &data.Entry_ts, &data.Dest, &data.Dest_ts, &data.Line)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}

func InsertTaluiComplete(entry string, entry_ts string, dest string, line string) (error){
	fmt.Println("Inserting TaluiComplete")
	db := database.DB
	
	_ = "INSERT INTO talui_complete (entry, entry_ts, dest, line) VALUES ( '" + entry + "', '" + entry_ts + "', '" + dest + "', '" + line + "' )"
	query := fmt.Sprintf("INSERT INTO talui_complete (entry, entry_ts, dest, line) VALUES ( '%s', '%s', '%s', '%s' )", entry, entry_ts, dest, line)
	insert, err := db.Query(query)
	
	if err != nil {
		fmt.Println("Error inserting")
		fmt.Println(err)
		return err
	}
	
	defer insert.Close()
	
	return nil
}

