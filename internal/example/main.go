package main

import (
	"context"
	"time"

	"github.com/go-jarvis/confredis"
	"github.com/sirupsen/logrus"
)

type person struct {
	Name string
	Age  int
}

func (p *person) String() {
	logrus.Infof("name = %s, age = %d", p.Name, p.Age)
}
func main() {
	r := redisServer()

	p1 := &person{
		Name: "zhangsan",
		Age:  30,
	}

	ctx := context.Background()

	key := "user1"
	err := r.Set(ctx, key, p1)
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second * 1)
		pp := &person{}
		err := r.Get(ctx, key, pp)
		if err != nil {
			panic(err)
		}
		pp.String()
	}

}

func redisServer() *confredis.Server {
	r := &confredis.Server{
		DB:       12,
		Password: "Redis123123",
	}
	r.Init()

	return r

}
