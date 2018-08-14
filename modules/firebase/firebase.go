package firebase

import (
	"context"
	"log"

	"firebase.google.com/go/db"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// CredentialsFile hoge
const (
	DatabaseURL = "https://nikuman-46d52.firebaseio.com"
)

// Create Firebase インスタンス
func Create(ctx context.Context) *firebase.App {
	conf := &firebase.Config{
		DatabaseURL: DatabaseURL,
	}
	opt := option.WithCredentialsFile("./config/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

// ConnectDb Firebase
func ConnectDb(ctx context.Context, app *firebase.App) *db.Client {
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return client
}

// GetRef Firebase
func GetRef(ctx context.Context, client *db.Client, query string) map[string]interface{} {
	ref := client.NewRef(query)
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	return data
}
