package dbmodels

import (
	"fmt"
	"talui/pkg/database"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

type TaluiOn struct {
	Id          int     `json:"id"`
	Entry       string  `json:"entry"`
	Entry_ts    string  `json:"entry_ts"`
	Dest        string  `json:"dest"`
	Line        string  `json:"line"`
}

func GetAllTaluiOn() ([]TaluiOn, error) {
	fmt.Println("get all TaluiOn")
	db := database.DB
	result, err := db.Query("SELECT * FROM `talui_on`")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer result.Close()
	
	dataArray := []TaluiOn{}
	for result.Next() {
		var data TaluiOn
		result.Scan(&data.Id, &data.Entry, &data.Entry_ts, &data.Dest, &data.Line)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}

func InsertTaluiOn(entry string, dest string, line string) (error) {
	fmt.Println("Inserting TaluiOn")
	db := database.DB
	
	_ = "INSERT INTO talui_on (entry, dest, line) VALUES ( '" + entry + "', '" + dest + "', '" + line + "' )"
	query := fmt.Sprintf("INSERT INTO talui_on (entry, dest, line) VALUES ( '%s', '%s', '%s' )", entry, dest, line)
	insert, err := db.Query(query)
	
	if err != nil {
		fmt.Println("Error inserting")
		fmt.Println(err)
		return err
	}
	
	defer insert.Close()
	
	return nil
}

func FindManyTaluiOn(dest string, line string) ([]TaluiOn, error) {
	fmt.Println("Find many TaluiOn")
	db := database.DB
	
	_ = "SELECT * FROM `talui_on` WHERE `dest` LIKE 'B' AND `line` LIKE 'green'"
	query := fmt.Sprintf("SELECT * FROM `talui_on` WHERE `dest` LIKE '%s' AND `line` LIKE '%s'", dest, line)
	
	result, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer result.Close()
	
	dataArray := []TaluiOn{}
	for result.Next() {
		var data TaluiOn
		result.Scan(&data.Id, &data.Entry, &data.Entry_ts, &data.Dest, &data.Line)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}

func RemoveManyTaluiOn(dest string, line string) ([]TaluiOn, error) {
	fmt.Println("Remove many TaluiOn")
	db := database.DB
	
	_ = "SELECT * FROM `talui_on` WHERE `dest` LIKE 'B' AND `line` LIKE 'green'"
	query := fmt.Sprintf("DELETE FROM `talui_on` WHERE `dest` LIKE '%s' AND `line` LIKE '%s'", dest, line)
	
	result, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer result.Close()
	
	dataArray := []TaluiOn{}
	for result.Next() {
		var data TaluiOn
		result.Scan(&data.Id, &data.Entry, &data.Entry_ts, &data.Dest, &data.Line)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}