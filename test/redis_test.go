package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestREDIS(t *testing.T) {
	c, err := redis.Dial("tcp", "8.134.120.93"+":"+"6379",
		redis.DialPassword("000415"))
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	defer c.Close()
}
