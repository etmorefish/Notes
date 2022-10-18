package main

import (
	"context"
	"fmt"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
	"time"
)

func main() {

	/*
		ch := make(chan string)
		go func() {
			for i := 1; i < 10; i++ {
				s := strconv.Itoa(i)
				ch <- s
			}
		}()
		for i := 1; i < 10; i++ {
			num := <-ch //从c中接收数据，并赋值给num
			fmt.Println("num = ", num)
		}
		time.Sleep(time.Second)
	*/

	db := ch.Connect(
		// pw 1uYxRua5
		// clickhouse://<user>:<password>@<host>:<port>/<database>?sslmode=disable
		//ch.WithDSN("clickhouse://:1uYxRua5@192.168.5.121:8123/default?sslmode=disable"),
		ch.WithDSN("clickhouse://localhost:9000/default?sslmode=disable"),
	)
	//db := ch.Connect(
	//	ch.WithAddr("192.168.5.121:9008"),
	//	//ch.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	//	//ch.WithUser("test"),
	//	ch.WithPassword("1uYxRua5"),
	//	ch.WithDatabase("default"),
	//	ch.WithTimeout(5*time.Second),
	//	ch.WithDialTimeout(5*time.Second),
	//	ch.WithReadTimeout(5*time.Second),
	//	ch.WithWriteTimeout(5*time.Second),
	//	ch.WithQuerySettings(map[string]interface{}{
	//		"prefer_column_name_to_alias": 1,
	//	}),
	//)

	db.AddQueryHook(chdebug.NewQueryHook(
		chdebug.WithVerbose(true),
		chdebug.FromEnv("CHDEBUG"),
	))
	ctx := context.Background()

	//res, err := db.NewCreateTable().Model((*Span)(nil)).Exec(ctx)
	//if err != nil {
	//	return
	//}
	//fmt.Println(res)
	span := &Span{}
	res, err := db.NewInsert().Model(span).Exec(ctx)
	if err != nil {
		return
	}
	fmt.Println(res)
}

type Span struct {
	ch.CHModel `ch:"partition:toYYYYMM(time)"`

	ID   uint64
	Text string    `ch:",lc"` // low cardinality column
	Time time.Time `ch:",pk"` // ClickHouse primary key to be used in order by
}
