package models

import (
	"Go-Blog-Project-master/forms"
	"Go-Blog-Project-master/helpers"
	"fmt"
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	About string `json:"about" bson:"about"`
	Image string `json:"image" bson:"image"`
	IsVerified bool          `json:"is_verified" bson:"is_verified"`
}

type UserModel struct{}

func (u *UserModel) Signup(data forms.SignupUserCommand) error {


	userCollection := dbConnect.Use("BlogDB", "user")

	err := userCollection.Insert(bson.M{
		"name":     data.Name,
		"email":    data.Email,
		"password": helpers.GeneratePasswordHash([]byte(data.Password)),
		"about": data.About,
		"image": data.Image,
		// This will come later when adding verification
		"is_verified": false,
	})

	return err
}

func (u *UserModel) GetUserByEmail(email string) (user User, err error) {

	userCollection := dbConnect.Use("BlogDB", "user")

	err = userCollection.Find(bson.M{"email": email}).One(&user)
	return user, err
}

func GetMyPage() (data forms.SignupUserCommand) {
var email string
	userCollection := dbConnect.Use("BlogDB", "user")

	err := userCollection.Find(bson.M{"email": email}).One(&data)
	if err != nil{
		fmt.Println(err)
	}

return

}


func (u *UserModel) GetUserByID(id string) (user User, err error) {
	userCollection := dbConnect.Use(databaseName, "user")

	err = userCollection.Find(bson.M{"_id": id}).All(&user)

	return user, err
}



func (u *UserModel) UpdateUserPass(email string, password string) (err error) {
	collection := dbConnect.Use(databaseName, "user")

	err = collection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"password": password}})

	return err
}

func (u *UserModel) VerifyAccount(email string) (err error) {
	collection := dbConnect.Use(databaseName, "user")

	err = collection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"is_verified": true}})

	return err
}
