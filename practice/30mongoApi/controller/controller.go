package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/shivamsalame/mongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

// MOST Important
var collection *mongo.Collection

// connect with mongoDB

// init will work at first time code runs
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ")
}

// MongoDB helpers - file

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id:", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string)  {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result,err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count:", deleteCount)
}

// delete all records
func deleteAllMovies() int64 {
	// empty parantheses means all records	
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return
}

// get all movies
func getAllMovies()  {
	collection.
}