package controller

import (
	"buildapi2/model"
	"buildapi2/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Yash:iJB!wnxSYg3r*4e@cluster0.8k2k5tn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbname = "netflix"
const colname = "watchlist"

var collection *mongo.Collection	//using mongodb collection

func init()  {		//mongodb connection
	clientOptions := options.Client().ApplyURI(connectionString)	//assigning mongodb URI to work on
	var client *mongo.Client
	client,err := mongo.Connect(context.TODO(),clientOptions)		//connecting mongodb

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection Success!")

	collection = client.Database(dbname).Collection(colname)	
}

func GetMyAllMovies(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	allmovies := service.GetAllmovies(collection)
	json.NewEncoder(w).Encode(allmovies)
	
}

func CreateMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	service.InsertOneMovie(collection,movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	service.UpdatetOneMovie(collection,params["id"])
	json.NewEncoder(w).Encode(params)	
}

func DeleteaMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	service.DeleteOneMovie(collection,params["id"])
	json.NewEncoder(w).Encode(params["id"])	
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	
	count := service.DeleteAllMovies(collection)
	json.NewEncoder(w).Encode(count)
}