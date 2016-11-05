package tasks

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
	"github.com/masnun/masnun-bot/dao"
	"github.com/masnun/masnun-bot/utils"
)

func PushReddit() {
	session, err := geddit.NewLoginSession(
		os.Getenv("REDDIT_USERNAME"),
		os.Getenv("REDDIT_PASSWORD"),
		"gedditAgent v1",
	)

	if err != nil {
		fmt.Println("Reddit Login Error: ", err)
		return
	}

	subOpts := geddit.ListingOptions{
		Limit: 15,
	}

	submissions, _ := session.Frontpage(geddit.DefaultPopularity, subOpts)

	for _, s := range submissions {
		if exists := dao.Exists(s.Permalink); !exists {
			fmt.Printf("Title: %s\nAuthor: %s\n\n", s.Title, s.Permalink)
			dao.Create(s.Permalink)
			utils.SendTelegramMessage(fmt.Sprintf("New reddit post: %s", s.Permalink))
		} else {
			fmt.Println("Exists: ", s.Permalink)
		}

	}

}
