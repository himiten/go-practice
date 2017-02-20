package db

import(
	"log"
	"gopkg.in/mgo"
	"gopkg.in/mgo/bson"
	"time"
)

func getMongoCli() *mgo.Session{
	session, err := mgo.Dial("192.168.10.21")
	if err != nil {
		log.Fatalln(err)
	}else {
		log.Printf("%T\n",session)
	}
	return session
}
type Person struct {
	Name string
	Sex byte
	Age int
	Birth time.Time
}

func operatMongoDB(){
	person:=Person{}
	sess:=getMongoCli()
	defer sess.Close()
	sess.SetMode(mgo.Monotonic, true)
	c:=sess.DB("test").C("people")
	err:=c.Insert(&Person{"hello",'f',27,time.Now()},
		&Person{"world",'m',22,time.Now()})
	if err !=nil{
		log.Fatal("插入",err)
	}
	c.RemoveAll(bson.M{"age": 22})
	err=c.Find(bson.M{"name":"hello"}).One(&person)
	if err !=nil{
		log.Fatal("查找",err)
	}
	log.Printf("%T\t%v\n",person,person)

}