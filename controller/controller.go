package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AliAkberAakash/mongo-with-go/config"
	"github.com/AliAkberAakash/mongo-with-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// helpers

func insertMovie(movie model.Netflix) {
	collection := config.GetCoursesCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	inserted, err := collection.InsertOne(ctx, movie)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Inserted movie with id ", inserted.InsertedID)
}

func markMovieAsWatched(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal("Invalid id format")
		return
	}

	collection := config.GetCoursesCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	updateOp := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(ctx, filter, updateOp)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("modified count", result.ModifiedCount)
}

func deleteSingleMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal("Invalid id format")
		return
	}

	collection := config.GetCoursesCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Deleted items", deleteCount)
}

func deleteAllMovies() int64 {
	collection := config.GetCoursesCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteMany(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
		return -1
	}

	return result.DeletedCount
}

func getAllMovies() []primitive.M {
	collection := config.GetCoursesCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cursor.Next(ctx) {
		var movie bson.M

		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, primitive.M(movie))
	}

	return movies
}
