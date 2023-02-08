package emoji

import (
	_ "embed"
	"encoding/json"
)

// emoji list from https://github.com/muan/unicode-emoji-json

type Emoji struct {
	Emoji string `form:"emoji" json:"emoji" xml:"emoji"`
	//Group string `form:"group" json:"group" xml:"group"`
	Slug string `form:"slug" json:"slug" xml:"slug"`
}

//go:embed emojis.json
var emojis []byte

var Emojis []Emoji
var EmojisMap map[string]bool

func init() {
	var parsedEmojis map[string]map[string]string
	json.Unmarshal(emojis, &parsedEmojis)

	Emojis = make([]Emoji, len(parsedEmojis))
	EmojisMap = make(map[string]bool, len(parsedEmojis))

	index := 0
	for key, value := range parsedEmojis {
		Emojis[index] = Emoji{
			Emoji: key,
			//Group: value["group"],
			Slug: value["slug"],
		}
		EmojisMap[key] = true
		index++
	}
}

func IsEmoji(emojiTest string) bool {
	if EmojisMap[emojiTest] == true {
		return true
	}
	return false
}
