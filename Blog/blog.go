package blogapp

import(
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
type BlogPost struct{
	Title string
	TLDR string
	Body string 
	Notes string
	images []string
}

func ShowLatest() []BlogPost{
	b:=BlogPost{"First Post", "Test", "I like Golang alot","This is cool", []string{"selfie.jpg","cake.png"}}
	fmt.Println("Hi")
	return []BlogPost{b}
}
func Initialize() *mongo.Collection{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://phelps:1995@cluster0.ddb3v.mongodb.net/blog?retryWrites=true&w=majority"))
	
	if err != nil{
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("blog")
	postsCollection := quickstartDatabase.Collection("posts")

	return postsCollection
	// blogResult, err := postsCollection.InsertOne(ctx, bson.D{
	// 	{Key: "title", Value: "My journey in programming"},
	// 	{Key: "topic", Value: "Code"},
	// 	{Key: "Author", Value: "Phelps Merrell"},
	// })
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(blogResult.InsertedID)	
}
