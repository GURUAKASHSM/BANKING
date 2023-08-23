package models

type Profile struct {
	//Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name          string `json:"name,omitempty" bson:"name"`
	Phoneno       int    `json:"no" bson:"no"`
	Profilestatus bool   `json:"status" bson:"ststus"`
}
