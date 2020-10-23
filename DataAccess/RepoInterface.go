package data
import(
    _ "github.com/go-sql-driver/mysql"
    "go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/bson/primitive"
    "log"
    //"fmt"
    "go.mongodb.org/mongo-driver/mongo/options")


type BlogStore interface { 
    Create(b BlogPost) interface{}
    UpdateByID() int64
    ReadAll() //BlogPost
	DeleteByID() bool
}



type MongoDb struct{
	DB *mongo.Client
}
func NewMongoDB(dataSourceName string) (*MongoDb, error){

    client, err := mongo.NewClient(options.Client().ApplyURI(dataSourceName))
	
	if err != nil{
		log.Fatal(err)
	}
	
    return &MongoDb{client}, nil
	
}



