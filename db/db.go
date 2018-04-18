package db

import (
	"database/sql"
	"fmt"
	"parallel/assets"
	"parallel/config"
	"strings"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

var dbpostgre *sql.DB
var err error
var dbredis *redis.Client

func DBConnectRedis() {
	dbredis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//fmt.Println(">>>>>>>>>>>>>>>>> Successfully connected to Database <<<<<<<<<<<<<<<<<")

}
func DBInsertRedis(id string, info string) {
	err := dbredis.Set(id, info, 0).Err()
	if err != nil {
		panic(err)
	}
}
func DBDeleteRedis(id string) {
	err := dbredis.Del(id).Err()
	if err != nil {
		panic(err)
	}
}
func DBGetAllKeysRedis() []string {
	var ReturnData []string
	allkeys, _ := dbredis.Keys("*").Result()
	for _, currentkey := range allkeys {
		//fmt.Println("KEY>> " + currentkey)
		currentvalue, _ := dbredis.Get(currentkey).Result()
		fmt.Println("KEY>> " + currentvalue)
		ReturnData = append(ReturnData, currentvalue)
	}
	return ReturnData
}
func DBConnectPostgres(configpath string) {
	allconfig := config.GetConfig(configpath)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", strings.Join(allconfig.DBHost, ""), strings.Join(allconfig.DBPort, ""), strings.Join(allconfig.DBUser, ""), strings.Join(allconfig.DBPass, ""), strings.Join(allconfig.DBName, ""))
	dbpostgre, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//fmt.Println(">>>>>>>>>>>>>>>>> Successfully connected to Database <<<<<<<<<<<<<<<<<")
}
func DBInsertPostgres(a *assets.Asset) {

	point := fmt.Sprintf(`'POINT( %.6f %.6f )'`, a.Lat, a.Lon)

	sqlStatement := `
		INSERT INTO parallel.webscrapingresults
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, postgis.ST_GeomFromText( ` + point + ` )  );`

	_, err := dbpostgre.Exec(sqlStatement, a.Business, a.Code, a.Type, a.Agency, a.Location, a.City, a.Area, a.Price, a.Numrooms, a.Numbaths, a.Parking, a.Status, a.Link)
	//fmt.Println(a.Business, a.Code, a.Type, a.Agency, a.Location, a.City, a.Area, a.Price, a.Numrooms, a.Numbaths, a.Status, a.Link)
	if err != nil {
		fmt.Println(err)
	}
}
