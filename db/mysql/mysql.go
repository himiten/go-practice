package mysql

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"os"
)

func GetInstance() *sql.DB{
	// Connect to database
	instance, err := sql.Open("mysql", "root:txj-123456@tcp(120.76.193.12:3306)/feiche51")
	if err != nil {
		panic(err.Error())
	}

	return instance
}
type City struct {
	Id int
	Name string
}

func GetCity(where string) (cities []City){
	instance:=GetInstance()
	rows,err := instance.Query("select id,name from city where 1=1 and "+where)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panicing %s\n", e)
		}
	}()
	defer instance.Close()
	//cities:=make([]City,10)
	for rows.Next() {
		var city City
		if err := rows.Scan(&city.Id,&city.Name); err != nil {
			panic(err.Error())
		}else {
			cities=append(cities,city)
			log.Printf("%v\n", cities)
		}

	}
	if err := rows.Err(); err != nil {
		panic(err.Error())

	}
	return cities
}
