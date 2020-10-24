package data

import (
	
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// type BlogPost struct {
// 	Title  string
// 	Topic  string
// 	Body   string
// 	Author string
// 	Images []string
// }

type BlogPost struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Title string `bson:"title,omitempty"`
	Topic string	`bson:"topic,omitempty"`
	Body string `bson:"body,omitempty"`
	Author string `bson:"author,omitempty"`
	Images []string `bson:"images,omitempty"`
}


func (db *MongoDb) Create(b BlogPost) interface{} {
	client := db.DB
	quickstartDatabase := client.Database("blog")
	postsCollection := quickstartDatabase.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Println(b)
	blogResult, err := postsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: string(b.Title)},
		{Key: "topic", Value: string(b.Topic)},
		{Key: "author", Value: string(b.Author)},
		{Key: "body", Value: string(b.Body)},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blogResult.InsertedID)
	return blogResult.InsertedID
}
func (db *MongoDb) UpdateByID() int64 {
	return 40
}
type ReadStruct struct{
	posts []bson.M
}
func (db *MongoDb) ReadAll() /*BlogPost*/{
	client := db.DB
	quickstartDatabase := client.Database("blog")
	postsCollection := quickstartDatabase.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	cursor, err := postsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var posts []bson.M
	if err = cursor.All(ctx, &posts); err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
	

}
func (db *MongoDb) DeleteByID() bool {
	return true
}
