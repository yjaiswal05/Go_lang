package service

import (
	"buildapi2/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
var collection *mongo.Collection	//using mongodb collection

func GetAllmovies(collection *mongo.Collection) []model.Netflix {
	cur, err := collection.Find(context.Background(), bson.D{{}}) //using MongoDb cursor
	if err != nil {
		log.Fatal(err)
	}
	var movies []model.Netflix
	for cur.Next(context.Background()) {
		var movie model.Netflix
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}
func InsertOneMovie(collection *mongo.Collection,movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie with id: ", inserted.InsertedID)
}
func UpdatetOneMovie(collection *mongo.Collection,movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}                       //filtering out our required data that needs to be updated
	update := bson.M{"$set": bson.M{"watched": true}} //bson.M or bson.D both can be used

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated count : ", updated.MatchedCount)
}
func DeleteOneMovie(collection *mongo.Collection,movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count : ", deleted.DeletedCount)
}
func DeleteAllMovies(collection *mongo.Collection) int64{
	filter := bson.D{{}} //selecting everything {}
	deleted, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count : ", deleted.DeletedCount)
	return deleted.DeletedCount
}