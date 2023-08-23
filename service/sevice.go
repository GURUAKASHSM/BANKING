package service

import (
	"context"
	"fmt"
	"log"
	"mongoapi/config"
	"mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(profile models.Profile) {

	inserted, err := config.Collection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted", inserted.InsertedID)

}
func Updateone(objectid string) {
	id, err := primitive.ObjectIDFromHex(objectid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"profilestatus": true}}
	result, err := config.Collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result of updated", result.MatchedCount)
}
func DeleteOne(objectid string) {
	id, err := primitive.ObjectIDFromHex(objectid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	delete, err := config.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", delete.DeletedCount)
}
func DeleteAll() int64 {
	filter := bson.D{{}}
	delete, err := config.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Many:", delete.DeletedCount)
	return delete.DeletedCount
}
func Getalldata() []primitive.M {
	coursor, err := config.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var Profiles []primitive.M

	for coursor.Next(context.Background()) {
		var profile bson.M
		err := coursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)

	}
	defer coursor.Close(context.Background())
	return Profiles
}
func Getdatabyid(objectid string) []primitive.M {
	id, err := primitive.ObjectIDFromHex(objectid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	coursor, err := config.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var Profiles []primitive.M

	for coursor.Next(context.Background()) {
		var profile bson.M
		err := coursor.Decode(&profile)
		if err != nil {
			log.Fatal(err)
		}
		Profiles = append(Profiles, profile)

	}
	defer coursor.Close(context.Background())
	return Profiles
}
