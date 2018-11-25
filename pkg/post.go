package root

import "time"

type Post struct {
	ID        string
	OwnerID   string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []string
	Likes     []string
}

type PostService interface {
	Post(p *Post) (Post, error)
}
