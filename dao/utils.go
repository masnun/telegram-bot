package dao

import (
	"fmt"
)

// Exists - check if an item exists in db
func Exists(PermaLink string) bool {
	db := Init()
	defer db.Close()
	var post RedditPost
	result := db.First(&post, "perma_link = ?", PermaLink)
	if result.Error != nil {
		errorMsg := result.Error.Error()

		if errorMsg == "record not found" {
			return false
		}

		fmt.Println(errorMsg)

	}

	return true

}

// Create post
func Create(PermaLink string) {
	db := Init()
	defer db.Close()
	var post = RedditPost{PermaLink: PermaLink}
	result := db.Create(&post)
	if result.Error != nil {
		errorMsg := result.Error.Error()
		fmt.Println(errorMsg)

	}
}
