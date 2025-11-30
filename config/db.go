package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// تست اتصال
	if err = Client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	log.Println("Connected to MongoDB!")

	// ایجاد Unique Index روی email (فقط یک بار اجرا می‌شه یا اگر وجود داشت رد می‌شه)
	createUniqueEmailIndex()
}

func GetCollection(name string) *mongo.Collection {
	return Client.Database("cinema").Collection(name)
}

func createUniqueEmailIndex() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetCollection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("unique_email"),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		if !isIndexConflict(err) {
			log.Printf("Warning: Could not create unique index on email: %v", err)
		}
	} else {
		log.Println("Unique index on email created successfully")
	}
}

// اگر ایندکس قبلاً وجود داشته باشه، مونگو ارور می‌ده. این تابع اون رو تشخیص می‌ده
func isIndexConflict(err error) bool {
	return err.Error() == "index already exists" || 
		   err.Error() == "an equivalent index already exists" ||
		   err.Error() == "Index already exists with different options"
}