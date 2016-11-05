package dao

import (
	"github.com/jinzhu/gorm"
)

// RedditPost - Struct for storing Reddit Posts
// Used by `gorm`
type RedditPost struct {
	gorm.Model
	PermaLink string
}
