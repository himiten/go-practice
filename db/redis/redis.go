package db

import (
	"log"
	"gopkg.in/redis.v4"
)
func getClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redistiger", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func UseClient(){
	cli:=getClient()
	err:=cli.Set("foo","hello",0).Err()
	if err!=nil{
		log.Fatalln(err)
	}
	val, err := cli.Get("foo").Result()
	if err==nil{
		log.Println(val)
	}
}
