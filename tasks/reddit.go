package tasks

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
	"github.com/masnun/telegram-bot/dao"
	"github.com/masnun/telegram-bot/utils"
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
			utils.SendTelegramMessage(fmt.Sprintf("%s : https://www.reddit.com/%s", s.Title, s.Permalink))
		} else {
			fmt.Println("Exists: ", s.Permalink)
		}

	}

}
