package models

type User struct {
	Id       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
}

type Account struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Bucket struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type LineItem struct {
	Id          int     `json:"id" bson:"id"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Amount      float64 `json:"amount" bson:"amount"`
}
