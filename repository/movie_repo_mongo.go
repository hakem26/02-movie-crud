package repository

import (
	"context"
	"example/moviecrud/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieMongoRepo struct {
	col *mongo.Collection
}

func NewMovieMongoRepo(col *mongo.Collection) *MovieMongoRepo {
	return &MovieMongoRepo{col}
}

func (r *MovieMongoRepo) Create(m *models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.col.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (r *MovieMongoRepo) FindAll() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var movies []*models.Movie
	for cur.Next(ctx) {
		var m models.Movie
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		movies = append(movies, &m)
	}
	return movies, cur.Err()
}

func (r *MovieMongoRepo) FindByID(id string) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var movie models.Movie
	err = r.col.FindOne(ctx, bson.M{"_id": objId}).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return  nil, err
	}
	return &movie, nil
}

func (r *MovieMongoRepo) FindByDirector(director *models.Director) (*models.Director, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dir models.Director
	err := r.col.FindOne(ctx, bson.M{"director": director}).Decode(&dir)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &dir, nil
}

func (r *MovieMongoRepo) Update(id string, m *models.Movie) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	
	update := bson.M{
		"$set": bson.M{
			"isbn": m.Isbn,
			"title": m.Title,
			"director": m.Director,
		},
	}
	res, err := r.col.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return nil, mongo.ErrNilDocument
	}

	return r.FindByID(id)
}

func (r *MovieMongoRepo) Delete(id string) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := r.col.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNilDocument
	}
	return nil
}