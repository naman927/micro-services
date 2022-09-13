package model

import (
	"context"
	"fmt"
	"time"

	"github.com/store/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpsertProduct(data types.StoreRequest) (int, error) {
	fmt.Println("upsert with data", data)
	db, err := GetDB()
	if err != nil {
		fmt.Println("err while getting db:", err)
		return -1, err
	}

	o := types.ProductModel{}
	product := db.Collection("product")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	p := types.ProductModel{}
	p.Product = data.Product
	p.URL = data.URL
	p.UpdatedAt = time.Now()

	product.FindOne(ctx, bson.M{"url": data.URL}).Decode(&o)
	if o.URL != "" {
		fmt.Println("updating.....")
		_, err := product.UpdateOne(ctx, bson.M{"url": o.URL}, bson.D{
			primitive.E{
				Key: "$set",
				Value: bson.D{
					primitive.E{
						Key:   "product",
						Value: p.Product,
					},
				},
			},
		})
		if err != nil {
			fmt.Println("err while updating the data:", err)
			return -1, nil
		}
		return 1, err
	} else {
		fmt.Println("inserting.....")
		p.CreatedAt = time.Now()
		_, err := product.InsertOne(ctx, p)
		if err != nil {
			fmt.Println("err while inserting the data:", err)
			return -1, nil
		}
		return 0, err
	}
}
