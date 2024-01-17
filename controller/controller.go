package controller

import (
	"fmt"
	"log"
	"os"
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/joho/godotenv"
)

var collection *mongo.Collection

func init(){
	
	if err:= godotenv.Load(); err!= nil{
		log.Fatal("Error loading env file: ",err.Error())
	}
	db_password := os.Getenv("MONGO_PASSWORD")
	db_user := os.Getenv("MONGO_USER")
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@tournament.hiorwpf.mongodb.net/?retryWrites=true&w=majority",db_user,db_password)

    const dbName = "codots"
	const colName = "matchups"

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal("Error connecting to database: ",err.Error())
	}

	fmt.Println("MongoDB connection success")

	collection := client.Database(dbName).Collection(colName)
	fmt.Printf("Collection instance '%s' is ready", collection.Name())
}

