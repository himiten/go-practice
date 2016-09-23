package db

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"os"
)

func getDB() *sql.DB{
	// Connect to database
	db, err := sql.Open("mysql", "root:my+tiger@/gyennoweb_record")
	if err != nil {
		panic(err.Error())
	}

	return db
}
type MyTit struct {
	Id int
	Title string
}
func GetReportTitles(where string){
	db:=getDB()
	rows,err := db.Query("select id,title from reports where 1=1 and "+where)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panicing %s\n", e)
		}
	}()
	defer db.Close()
	for rows.Next() {
		var title MyTit
		if err := rows.Scan(&title.Id); err != nil {
			panic(err.Error())
		}
		log.Printf("%v\n", title)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

}
