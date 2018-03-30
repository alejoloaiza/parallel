package db

import (
  "database/sql"
  "fmt"
  "parallel/config"
  "strings"
  _"github.com/lib/pq"
)

var dbcon *sql.DB
var err error
func DBConnectPostgres(configpath string) {
	allconfig := config.GetConfig(configpath)
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",strings.Join(allconfig.DBHost, ""), strings.Join(allconfig.DBPort, ""), strings.Join(allconfig.DBUser, ""),strings.Join(allconfig.DBPass, ""), strings.Join(allconfig.DBName, ""))
	dbcon, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  fmt.Println(">>>>>>>>>>>>>>>>> Successfully connected to Database <<<<<<<<<<<<<<<<<")
}
func DBInsertPostgres(id string,agency string, sector string, price string, area string, rooms string, baths string, link string, status string){
	sqlStatement := `
	INSERT INTO public.webscrapingresults (id,agency,sector,price,area,rooms,baths,link,status)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := dbcon.Exec(sqlStatement, id, agency, sector, price,area,rooms,baths,link, true )
	if err != nil {
	  fmt.Println(err)
	}
}