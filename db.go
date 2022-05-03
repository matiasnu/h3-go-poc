package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	data *MongoDB
)

type MongoDB struct {
	client *mongo.Client
}

func initDb() {
	host := "localhost"
	port := 27017

	credential := options.Credential{
		Username: "matias",
		Password: "matias",
	}
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")
	data = &MongoDB{
		client: client,
	}
}

func getShopsByRing(lookupIndexes []string) {
	var result []ReferencePoint
	collection := data.client.Database("shops").Collection("shops")
	filter := bson.M{"indexh3": bson.M{"$in": lookupIndexes}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if err = cur.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	for _, result := range result {
		fmt.Printf("Shop ID %s near in %dkm\n", result.ShopID, searchRadiusKm)
	}
}

func addShop(shopid string, index string) {
	insert := ReferencePoint{
		Indexh3: index,
		ShopID:  shopid,
	}
	collection := data.client.Database("shops").Collection("shops")
	insertResult, err := collection.InsertOne(context.TODO(), insert)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Shop had been inserted: ", insertResult.InsertedID)
}
