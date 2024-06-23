package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string             `json:"title"`
	Cover        string             `json:"cover"` // URL gambar cover buku
	Year         int                `json:"year"`
	Synopsis     string             `json:"synopsis"`
	Author       string             `json:"author"`
	Genre        string             `json:"genre"`
	Location     string             `json:"location"`
	Holder       string             `json:"holder"`
	BorrowedDate time.Time          `json:"borrowed_date"`
}
