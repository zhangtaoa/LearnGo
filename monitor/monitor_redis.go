package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func redis_connect() (c redis.Conn, err error){
	redis_conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis err",err)
		return
	}
	return redis_conn, err
}

func redis_get(c redis.Conn, key string) (rsp string, err error) {
	return redis.String(c.Do("get", key))
}


func main() {
	redis_con, err := redis_connect()
	if err != nil {
		fmt.Println("redis_connect err:", err)
	}

	redis_con.Do("set", "name", "zhangtao")

	res, err := redis.String(redis_conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get ", err)
	}

	fmt.Println((res))

}
