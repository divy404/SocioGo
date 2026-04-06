package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)
var (
	ErrNotFound = errors.New("record not found")
	ErrConflict = errors.New("record already exists")
	QueryTimeoutDuration = time.Second * 5
)
type Storage struct {
	Posts interface {
		GetbyID(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
		Delete(context.Context,int64) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, int64) ([]PostWithMetadata, error)
	}
	Users interface {
		GetbyID(context.Context, int64) (*User, error)
		Create(context.Context, *User) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error) 
	}
	Followers interface {
		Follow(context.Context, int64, int64) error
		Unfollow(context.Context, int64, int64) error
	}
}

	func NewStorage(db *sql.DB) Storage {
		return Storage {
			Posts: &PostStore{db},
			Users: &UserStore{db},
			Comments: &CommentStore{db},
			Followers: &FollowerStore{db},
		}
	}
// just adding comment