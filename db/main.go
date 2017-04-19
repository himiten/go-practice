package main

import (
	"log"
	my "./mysql"
)
type City struct {
	Id int
	Name string
}
func main()  {
	instance:=my.GetInstance()
	rows,err := instance.Query("select id,name from city where id=1")
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(rows)
	defer instance.Close()
	for rows.Next() {
		var city City
		if err := rows.Scan(&city.Id,&city.Name); err != nil {
			panic(err.Error())
		}
		log.Printf("%v\n", city)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	ct:=my.GetCity("id in (2,3)")
	log.Printf("%v\n", ct)
}
