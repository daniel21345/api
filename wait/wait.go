package wait
import (
	"fmt"
	"github.com/daniel21345/wait/redis"
	
)


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
