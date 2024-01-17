package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Codot struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type     string             `json:"type,omitempty"`
	Defeated bool               `json:"defeated,omitempty"`
}

