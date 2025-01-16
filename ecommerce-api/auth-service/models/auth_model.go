package models

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
}