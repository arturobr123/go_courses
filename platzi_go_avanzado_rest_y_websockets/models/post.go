package models

import "time"

// Post model
type Post struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	PostContent string    `json:"post_content"`
	CreatedAt   time.Time `json:"created_at"`
}
