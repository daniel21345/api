package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	
)
var (
	host = "127.0.0.1:6379"
	conn,err = redis.Dial("tcp",host)
	url string
	)
func main(){
fmt.Println("input your url")
fmt.Scanf("%s",&url)
	addToKey(url)
}
func addToKey(str string){
	ex,_:=redis.Bool(conn.Do("exists",str))
	if ex == false{
		conn.Do("set",str,"")
		conn.Do("expire",str,"15")
		fmt.Println("login success")
	}else{
		ttl,_:=conn.Do("ttl",str)
		fmt.Println("please wait fot",ttl,"second")
	}
}