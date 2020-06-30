package models

import (
	"Go-Blog-Project-master/forms"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}


type PostModel struct{}

func (p *PostModel) NewPost(data forms.PostCommand) error {

	postCollection := dbConnect.Use("BlogDB", "posts")

	err := postCollection.Insert(bson.M{
		"title": data.Title,
		"content": data.Content,
	})


	return err
}




func (p *PostModel) GetPostByID(id string) (post Post, err error) {
	postCollection := dbConnect.Use(databaseName, "posts")

	err = postCollection.Find(bson.M{"_id": id}).One(&post)

	return post, err
}

func GetAllPosts() (data []forms.PostCommand) {
	postCollection := dbConnect.Use("BlogDB", "posts")

	err := postCollection.Find(nil).All(&data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

