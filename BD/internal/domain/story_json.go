package domain

import (
	"math/rand"
	"time"

	"github.com/goombaio/namegenerator"
)

type (
	StoryJSON struct {
		User      string   `json:"user"`
		StoryType string   `json:"story_type"`
		Duration  int      `json:"duration"`
		Media     Media    `json:"media"`
		Captions  Captions `json:"captions"`
		Tags      []string `json:"tags"`
	}

	Media struct {
		URL  string `json:"url"`
		Size string `json:"size"`
	}

	Captions struct {
		Text            string `json:"text"`
		Style           string `json:"style"`
		TextColor       string `json:"text_color"`
		BackgroundColor string `json:"background_color"`
	}
)

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func NewStoryJSON() StoryJSON {
	const (
		basicFieldLen = 20
		urlFieldLen   = 10
		textFieldLen  = 120
	)

	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return StoryJSON{
		User:      nameGenerator.Generate(),
		StoryType: randomString(basicFieldLen),
		Duration:  rand.Intn(30) + 1,
		Captions: Captions{
			Text:            randomString(textFieldLen),
			Style:           randomString(basicFieldLen),
			TextColor:       randomString(basicFieldLen),
			BackgroundColor: randomString(basicFieldLen),
		},
		Media: Media{
			URL:  "https://example.com/" + randomString(urlFieldLen),
			Size: "small",
		},
		Tags: []string{"nature", "sunset", "photography"},
	}
}
