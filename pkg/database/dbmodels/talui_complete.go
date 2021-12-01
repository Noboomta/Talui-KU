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

type StationCount struct {
	Station     string  `json:"station"`
	Value       int     `json:"value"`
	Time        string  `json:"time"`
}

func GetAllTaluiComplete() ([]TaluiComplete, error) {
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
	db := database.DB
	
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

func GetStationUsingEntry() ([]StationCount, error) {
	db := database.DB
	
	query := "SELECT entry as station, count(entry) as value, TIMESTAMP(DATE(entry_ts)" +
	" , CONCAT(floor(hour(entry_ts)/2)*2, ':00:00')) as time" + 
	" FROM talui_complete GROUP BY station, time"
	result, err := db.Query(query);
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	defer result.Close()
	
	dataArray := []StationCount{}
	for result.Next() {
		var data StationCount
		result.Scan(&data.Station, &data.Value, &data.Time)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}

func GetStationUsingDest() ([]StationCount, error) {
	db := database.DB
	
	query := "SELECT dest as station, count(dest) as value, TIMESTAMP(DATE(dest_ts)" +
	" , CONCAT(floor(hour(dest_ts)/2)*2, ':00:00')) as time" + 
	" FROM talui_complete GROUP BY station, time"
	result, err := db.Query(query);
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	defer result.Close()
	
	dataArray := []StationCount{}
	for result.Next() {
		var data StationCount
		result.Scan(&data.Station, &data.Value, &data.Time)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}

func FindStationUsingEntryByStationAndTime(station string, time string) ([]StationCount, error) {
	db := database.DB
	
	// query := "SELECT total.station as station, total.value, total.time" +
	// "FROM (" +
	// "SELECT entry as station, count(entry) as value, TIMESTAMP(DATE(entry_ts), " +
	// "CONCAT(floor(hour(entry_ts)/2)*2, ':00:00')) as time" +
	// "FROM talui_complete GROUP BY station, time" +
	// ") as total" +
	// "WHERE" +
	// fmt.Sprintf("total.station = '%s'", station) +
	// fmt.Sprintf("AND total.time = '%s'", time)
	
	q2 := fmt.Sprintf("SELECT total.station as station, total.value, total.time FROM ( SELECT entry as station, count(entry) as value, TIMESTAMP(DATE(entry_ts), CONCAT(floor(hour(entry_ts)/2)*2, ':00:00')) as time FROM talui_complete GROUP BY station, time ) as total WHERE total.station = '%s' AND total.time = '%s'", station, time)
	
	result, err := db.Query(q2);
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	defer result.Close()
	
	dataArray := []StationCount{}
	for result.Next() {
		var data StationCount
		result.Scan(&data.Station, &data.Value, &data.Time)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}

func FindStationUsingDestByStationAndTime(station string, time string) ([]StationCount, error) {
	db := database.DB
	
	// query := "SELECT total.station as station, total.value, total.time" +
	// "FROM (" +
	// "SELECT dest as station, count(dest) as value, TIMESTAMP(DATE(dest_ts), " +
	// "CONCAT(floor(hour(dest_ts)/2)*2, ':00:00')) as time" +
	// "FROM talui_complete GROUP BY station, time" +
	// ") as total" +
	// "WHERE" +
	// fmt.Sprintf("total.station = '%s'", station) +
	// fmt.Sprintf("AND total.time = '%s'", time)
	
	q2 := fmt.Sprintf("SELECT total.station as station, total.value, total.time FROM ( SELECT dest as station, count(dest) as value, TIMESTAMP(DATE(dest_ts), CONCAT(floor(hour(dest_ts)/2)*2, ':00:00')) as time FROM talui_complete GROUP BY station, time ) as total WHERE total.station = '%s' AND total.time = '%s'", station, time)
	
	fmt.Println(q2)
	
	result, err := db.Query(q2);
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	defer result.Close()
	
	dataArray := []StationCount{}
	for result.Next() {
		var data StationCount
		result.Scan(&data.Station, &data.Value, &data.Time)
		dataArray = append(dataArray, data)
	}
	
	return dataArray, nil
}