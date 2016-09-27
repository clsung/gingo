package main

import (
	"flag"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

var (
	redisAddress = flag.String("redis-address", ":6379", "Address to the Redis server")
)

func main() {
	router := gin.Default()
	router.GET("/ready", checkRedis)
	router.GET("/ready/:auth", checkRedisWithAuth)

	router.Run()
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
