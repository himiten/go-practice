package db

import (
	"log"
	"github.com/garyburd/redigo/redis"
	"time"
	"flag"
)

func getRedigoCli() (redis.Conn ,error){
	conn, err := redis.Dial("tcp", "localhost:6379",redis.DialPassword("redistiger"))
	/*if err != nil {
		// handle error
		log.Fatalln(err)
	}
	if _, err := conn.Do("AUTH", "redistiger"); err != nil {
		conn.Close()
		return nil, err
	}*/
	return  conn,err
}

func UseRedigoCli(){
	cli,err:=getRedigoCli()
	if err!=nil{
		panic(err.Error())
	}
	defer cli.Close()
	n,err:=cli.Do("SET","foo","hello")

	if err!=nil{
		log.Fatalln(err)
	}else {
		log.Println(n)
	}
	reply, err := redis.String(cli.Do("GET","foo"))
	if err!=nil{
		log.Fatalln(err)
	}else{
		log.Println("reply:",reply)
	}

}
var (
	//pool *redis.Pool
	redisServer = flag.String("redisServer","localhost:6379","")
	redisPassword = flag.String("redisPassword", "redistiger", "")
)
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 4,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
func getConnPool() *redis.Pool{
	flag.Parse()
	log.Printf("%v\n",*redisServer)
	log.Printf("%v\n",*redisPassword)
	return newPool(*redisServer, *redisPassword)
}

func UseConnPool(){
	conn := getConnPool().Get()
	defer conn.Close()

	ok,err:=conn.Do("SET","foo","world")

	if err!=nil{
		log.Fatalln(err)
	}else {
		log.Println(ok)
	}
	reply, err := redis.String(conn.Do("GET","foo"))
	if err!=nil{
		log.Fatalln(err)
	}else{
		log.Println("reply:",reply)
	}
}


