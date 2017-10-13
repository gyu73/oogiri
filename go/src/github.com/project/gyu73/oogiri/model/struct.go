package model

import (
	"time"
)

// User returns model object for user.
type User struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Salt    string     `json:"salt"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
	Oogiris []Oogiri
	Answers []Answer
}

// Article returns model object for article.
type Oogiri struct {
	ID      int      `json:"id"`
	Title   string     `json:"title"`
	Body    string     `json:"body"`
	Image   string     `json:"image"`
	UserId  int64      `json:"user_id"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
	Deleted *time.Time `json:"deleted"`
	Answers []Answer
}

type Answer struct {
	Id       int64
	UserId   int64      `json:"comment_id"`
	OogiriId int64      `json:"article_id"`
	Body     string     `json:"body"`
	Created  *time.Time `json:"created"`
	Updated  *time.Time `json:"updated"`
	Oogiri	 Oogiri
	User	 User
}
