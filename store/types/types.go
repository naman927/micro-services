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
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	URL       string             `json:"url,omitempty" bson:"url,omitempty"`
	Product   Product            `json:"product,omitempty" bson:"product,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_At,omitempty" bson:"updated_at,omitempty"`
}

