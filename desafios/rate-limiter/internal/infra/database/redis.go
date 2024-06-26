package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct{
	dbHost string
	dbPort int
	dbPassword string
	dbName int

}

func NewRedisClient(host string, port int, password string, dbName int) *Redis {
	return &Redis{
		host,
		port,
		password,
		dbName,
	}
}

func(r *Redis) Client() *redis.Client{
	client := r.connectDB()
	_, err := client.Ping().Result()

	if err != nil {
        panic(err)
    }

	fmt.Println("Redis successfully connected")
	return client
}
 
func(r *Redis) connectDB() *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%v", r.dbHost, r.dbPort), //"localhost:6379"
		Password: "", // r.dbPassword
		DB: r.dbName, // 0
    })
}