package forms

import "gopkg.in/mgo.v2/bson"

// SignupUserCommand defines user form struct
type SignupUserCommand struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	About string `json:"about" binding:"required"`
	Image string `json:"image" binding: "required"`
}

// LoginUserCommand defines user login form struct
type LoginUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PasswordResetCommand defines user password reset form struct
type PasswordResetCommand struct {
	Password string `json:"password" binding:"required"`
	Confirm  string `json:"confirm" binding:"required"`
}

// ResendCommand defines resend email payload
type ResendCommand struct {
	Email string `json:"email" binding:"required"`
}

type PostCommand struct {
	Id bson.ObjectId`json:"_id,omitempty" binding:"required"`
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}