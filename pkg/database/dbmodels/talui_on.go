package dbmodels

import (
	"fmt"
	"talui/pkg/database"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

type TaluiOn struct {
	Id       int    `json:"id"`
	Entry    string `json:"entry"`
	Entry_ts string `json:"entry_ts"`
	Dest     string `json:"dest"`
	Line     string `json:"line"`
}

type TaluiTracker struct {
	Line    string `json:"line"`
	Current string `json:"current"`
	Next    string `json:"next"`
}

func GetAllTaluiOn() ([]TaluiOn, error) {
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

func InsertTaluiOn(entry string, dest string, line string) error {
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

func UpdateTracker(curr string, next string, line string) error {
	db := database.DB

	_ = "INSERT INTO talui_on (entry, dest, line) VALUES ( '" + curr + "', '" + next + "', '" + line + "' )"
	query := fmt.Sprintf("UPDATE `talui_tracker` SET current='%s', next='%s' WHERE line='%s'", curr, next, line)
	update, err := db.Query(query)

	if err != nil {
		fmt.Println("Error inserting")
		fmt.Println(err)
		return err
	}

	defer update.Close()
	return nil
}

func GetTaluiTracker() ([]TaluiTracker, error) {
	db := database.DB
	result, err := db.Query("SELECT * FROM `talui_tracker`")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer result.Close()

	dataArray := []TaluiTracker{}
	for result.Next() {
		var data TaluiTracker
		result.Scan(&data.Line, &data.Current, &data.Next)
		dataArray = append(dataArray, data)
	}

	return dataArray, nil
}
