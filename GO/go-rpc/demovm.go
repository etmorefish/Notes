package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi"
)

const (
	ip       = ""
	user     = ""
	password = ""
)

func main() {
	u := &url.URL{
		Scheme: "https",
		Host:   ip,
		Path:   "/sdk",
	}
	ctx := context.Background()
	u.User = url.UserPassword(user, password)
	client, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}
