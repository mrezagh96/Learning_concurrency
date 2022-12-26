package main

import (
	"fmt"
	"github.com/jzelinskie/geddit"
	"os"
)

var redditSession *geddit.LoginSession

type Story struct {
	title string
	url   string
}

func init() {
	var err error
	redditSession, err = geddit.NewLoginSession("Nearby-Gene-3146", "njfZMX8keh5wm4z", "gdAgent v0")
	if err != nil {
		fmt.Println(err)
		os.Exit(11)
	}

}

func main() {
	S := loadReddit()
	fmt.Println(S)
}

func loadReddit() []Story {
	var stories []Story
	var listoption geddit.ListingOptions
	submission, err := redditSession.SubredditSubmissions("programming", "", listoption)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, S := range submission {
		story := Story{
			title: S.Title,
			url:   S.URL,
		}
		stories = append(stories, story)
	}
	return stories
}
