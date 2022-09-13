package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoreRequest struct {
	URL     string  `json:"url"`
	Product Product `json:"product"`
}

type Product struct {
	Title       string  `json:"title"`
	ImageUrl    string  `json:"imageurl"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	NoOfReviews string  `json:"totalreview"`
}

type ProductModel struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	URL       string             `json:"url" bson:"url"`
	Product   Product            `json:"product" bson:"product"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_At" bson:"updated_at"`
}
