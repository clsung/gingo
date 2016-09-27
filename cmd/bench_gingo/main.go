package main

import (
	"flag"
	"log"

	"github.com/clsung/gingo"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

var (
	redisAddress = flag.String("redis-address", ":6379", "Address to the Redis server")
	redisAuth    = flag.String("redis-auth", "", "Password to the Redis server")
)

func main() {
	flag.Parse()
	log.Printf("redis address is %s\n", *redisAddress)
	log.Printf("redis auth is %s\n", *redisAuth)
	routerNoPool := gin.Default()
	routerNoPool.GET("/ready", checkRedis)
	routerNoPool.GET("/ready/:auth", checkRedisWithAuth)
	go routerNoPool.Run()

	redis := gingo.NewRedisStore(*redisAddress, *redisAuth)
	router := gin.Default()
	router.GET("/ready", func(c *gin.Context) {
		ret, err := redis.Do("PING")
		if err != nil {
			c.JSON(500, gin.H{
				"message": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": ret,
		})
	})

	router.Run(":8181")
}

func checkRedisWithAuth(c *gin.Context) {
	auth := c.Param("auth")
	doRedis(c, "PING", auth)
}

func checkRedis(c *gin.Context) {
	doRedis(c, "PING", "")
}

func doRedis(c *gin.Context, cmd, auth string) {
	client, err := redis.Dial("tcp", *redisAddress)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	defer client.Close()
	if auth != "" {
		if _, err := client.Do("AUTH", auth); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid password",
			})
			return
		}
	}
	ret, err := client.Do(cmd)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": ret,
	})
}
