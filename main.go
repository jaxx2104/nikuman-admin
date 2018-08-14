package main

import (
	"context"
	"fmt"

	"./modules/firebase"
)

const (
	isTest = true
)

func main() {
	ctx := context.Background()
	app := firebase.Create(ctx)
	client := firebase.ConnectDb(ctx, app)
	data := firebase.GetRef(ctx, client, "posts")

	for _, item := range data {
		itemMap := item.(map[string]interface{})
		fmt.Println(itemMap["body"])
	}
	// fmt.Println(data)
}
