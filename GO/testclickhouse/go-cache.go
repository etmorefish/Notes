package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	c := cache.New(10*time.Second, 30*time.Second) // 默认过期时间10s；清理间隔30s，即每30s钟会自动清理过期的键值对

	// 设置一个键值对，过期时间是 3s
	c.Set("a", "testa", 3*time.Second)

	// 设置一个键值对，采用 New() 时的默认过期时间，即 10s
	c.Set("foo", "bar", cache.DefaultExpiration)

	// 设置一个键值对，没有过期时间，不会自动过期，需要手动调用 Delete() 才能删除
	c.Set("baz", 42, cache.NoExpiration)

	v, found := c.Get("a")
	fmt.Println(v, found) // testa,true

	<-time.After(5 * time.Second) // 延时5s

	v, found = c.Get("a") // nil,false
	fmt.Println(v, found)

	<-time.After(6 * time.Second)
	v, found = c.Get("foo") // nil,false
	fmt.Println(v, found)

	v, found = c.Get("baz") // 42,true
	fmt.Println(v, found)

	//TestCache()
	//TestCacheTimes()
	//TestNewFrom()
	//TestOnEvicted()
	//TestFileSerialization()
}
